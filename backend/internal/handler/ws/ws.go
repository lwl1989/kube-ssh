package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-libraries/kube-manager/backend/internal/cache/token"
	"github.com/go-libraries/kube-manager/backend/internal/config"
	"github.com/go-libraries/kube-manager/backend/internal/handler/ws/localcommand"
	"github.com/go-libraries/kube-manager/backend/internal/handler/ws/slave"
	"github.com/go-libraries/kube-manager/backend/internal/logic/ops/common"
	"github.com/go-libraries/kube-manager/backend/internal/model"
	"github.com/go-libraries/kube-manager/backend/internal/svc"
	"github.com/go-libraries/kube-manager/backend/internal/webtty"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

// WsServer provides a webTty HTTP endpoint.
type WsServer struct {
	factory  slave.Factory
	upgrader *websocket.Upgrader
	svcCtx   *svc.ServiceContext
	counter  *Counter
}

func NewWsServer(svcCtx *svc.ServiceContext, counter *Counter) *WsServer {
	factory, _ := localcommand.NewFactory("kubectl", []string{}, &localcommand.Options{
		CloseSignal:  1,
		CloseTimeout: -1,
	})
	var originChecker = func(r *http.Request) bool {
		return true
	}
	if svcCtx.Config.WSOrigin != "" {
		matcher, _ := regexp.Compile(svcCtx.Config.WSOrigin)
		if matcher != nil {
			originChecker = func(r *http.Request) bool {
				return matcher.MatchString(r.Header.Get("Origin"))
			}
		}
	}
	return &WsServer{
		factory: factory,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  webtty.MaxBufferSize,
			WriteBufferSize: webtty.MaxBufferSize,
			Subprotocols:    webtty.Protocols,
			CheckOrigin:     originChecker,
		},
		svcCtx:  svcCtx,
		counter: counter,
	}
}

func (server *WsServer) GenerateHandleWS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		}

		//todo:limit
		//if server.options.Once {
		//	success := atomic.CompareAndSwapInt64(once, 0, 1)
		//	if !success {
		//		http.Error(w, "Server is shutting down", http.StatusServiceUnavailable)
		//		return
		//	}
		//}

		num := server.counter.add(1)
		closeReason := "unknown reason"
		defer func() {
			x := recover()
			logx.Infof("%+v", x)
			num := server.counter.done()
			logx.Infof(
				"Connection closed: %s, reason: %s, connections: %d/%d",
				r.RemoteAddr, closeReason, num, server.svcCtx.Config.MaxWsConnection,
			)

			//if server.options.Once {
			//	cancel()
			//}
		}()

		if num > server.svcCtx.Config.MaxWsConnection {
			closeReason = "exceeding max number of connections"
			return
		}

		logx.Infof("New client connected: %s, connections: %d/%d", r.RemoteAddr, num, server.svcCtx.Config.MaxWsConnection)

		server.upgrader.ReadBufferSize = webtty.MaxBufferSize
		server.upgrader.WriteBufferSize = webtty.MaxBufferSize
		server.upgrader.EnableCompression = true
		conn, err := server.upgrader.Upgrade(w, r, nil)
		if err != nil {
			closeReason = err.Error()
			return
		}
		var cfg string
		defer func() {
			_ = conn.Close()
			if cfg != "" {
				_ = os.Remove(cfg)
			}
		}()
		_ = conn.SetCompressionLevel(9)
		err, cfg = server.processWSConn(ctx, conn)

		switch err {
		case ctx.Err():
			closeReason = "cancelation"
		case webtty.ErrSlaveClosed:
			closeReason = server.factory.Name()
		case webtty.ErrMasterClosed:
			closeReason = "client close"
		case webtty.ErrConnectionLostPing:
			closeReason = webtty.ErrConnectionLostPing.Error()
		default:
			closeReason = fmt.Sprintf("an error: %s", err)
		}
	}
}

func (server *WsServer) processWSConn(ctx context.Context, conn *websocket.Conn) (error, string) {
	typ, initLine, err := conn.ReadMessage()
	if err != nil {
		return errors.Wrapf(err, "failed to authenticate websocket connection"), ""
	}
	if typ != websocket.TextMessage {
		return errors.New("failed to authenticate websocket connection: invalid message type"), ""
	}

	var init InitMessage
	err = json.Unmarshal(initLine, &init)
	if err != nil {
		return errors.Wrapf(err, "failed to authenticate websocket connection"), ""
	}

	queryPath := init.Arguments
	query, err := url.Parse(queryPath)
	if err != nil {
		return errors.Wrapf(err, "failed to parse arguments"), ""
	}
	params := query.Query()
	tk := params.Get("token")
	defer func() {
		if tk != "" {
			_ = svc.GlobalService.Cache.Delete(tk)
		}
	}()
	ttyArgs, err := server.checkAuth(ctx, tk)
	params = ttyArgsToParams(ttyArgs, params)

	configPath := "/tmp/k8s-manager/" + tk
	var cluster model.Cluster
	svc.GetDb(config.DbDefault).Table(model.ClusterTableName()).Where("id=?", ttyArgs.Id).First(&cluster)
	ok, err := cluster.CheckCanUse()
	if !ok {
		return errors.Wrapf(err, "集群信息选择错误"), ""
	}
	err = os.WriteFile(configPath, bytes.NewBufferString(cluster.Config).Bytes(), 0600)
	params = server.buildCommand(params, configPath)
	var sla slave.Slave
	sla, err = server.factory.New(params)
	if err != nil {
		return errors.Wrapf(err, "failed to create backend"), configPath
	}
	defer sla.Close()

	opts := []webtty.Option{
		webtty.WithWindowTitle([]byte(fmt.Sprintf("ssh@root:%s@%s", params.Get("pod"), params.Get("container")))),
	}
	opts = append(opts, webtty.WithPermitWrite())

	tty, err := webtty.New(&wsWrapper{conn}, sla, opts...)
	if err != nil {
		return errors.Wrapf(err, "failed to create webtty"), configPath
	}

	err = tty.Run(ctx)

	return err, configPath
}

func (server *WsServer) buildCommand(params url.Values, configPath string) url.Values {
	params.Add("arg", "--kubeconfig="+configPath)
	params.Add("arg", "exec")
	params.Add("arg", "-it")
	params.Add("arg", params.Get("pod"))
	params.Add("arg", "-n")
	params.Add("arg", params.Get("namespace"))
	params.Add("arg", "-c")
	params.Add("arg", params.Get("container"))
	params.Add("arg", "--")
	params.Add("arg", "/bin/sh")
	return params
}

func (server *WsServer) checkAuth(ctx context.Context, tk string) (*token.TtyParameter, error) {
	var err error
	if tk == "" {
		return nil, errors.Wrapf(err, "鉴权失败")
	}
	ttyArgs := svc.GlobalService.Cache.Get(tk)
	if ttyArgs == nil {
		return ttyArgs, errors.Wrapf(err, "鉴权失败")
	}
	us := common.GetUser(ttyArgs)
	if us == nil {
		return ttyArgs, errors.Wrapf(err, "鉴权失败")
	}
	ctx = context.WithValue(ctx, "user", *us)
	err = common.CheckUserHasPermission(ctx, ttyArgs.Id)
	if err != nil {
		return ttyArgs, errors.Wrapf(err, err.Error())
	}

	return ttyArgs, nil
}

// titleVariables merges maps in a specified order.
// varUnits are name-keyed maps, whose names will be iterated using order.
func (server *WsServer) titleVariables(order []string, varUnits map[string]map[string]interface{}) map[string]interface{} {
	titleVars := map[string]interface{}{}

	for _, name := range order {
		vars, ok := varUnits[name]
		if !ok {
			panic("title variable name error")
		}
		for key, val := range vars {
			titleVars[key] = val
		}
	}

	// safe net for conflicted keys
	for _, name := range order {
		titleVars[name] = varUnits[name]
	}

	return titleVars
}

func ttyArgsToParams(ttyArgs *token.TtyParameter, params url.Values) url.Values {
	args, _ := url.ParseQuery(ttyArgs.Arg)
	for k, v := range args {
		if len(v) > 0 {
			params.Set(k, v[0])
		}
	}
	return params
}
