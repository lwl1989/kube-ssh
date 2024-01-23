package ops

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	cache "github.com/lwl1989/kube-ssh/backend/internal/cache/token"
	"github.com/lwl1989/kube-ssh/backend/internal/handler/api_utils"
	"github.com/lwl1989/kube-ssh/backend/internal/handler/ws"
	"github.com/lwl1989/kube-ssh/backend/internal/library/utils"
	"github.com/lwl1989/kube-ssh/backend/internal/logic/ops/common"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
	"time"
)

type ApiResponse struct {
	Token string `json:"token"`
}

type KubeConfigRequest struct {
	Id        int    `json:"id"`
	Pod       string `json:"pod"`
	Container string `json:"container"`
	Namespace string `json:"namespace"`
}

type KubeTokenRequest struct {
	Name      string `json:"name"`
	ApiServer string `json:"apiServer"`
	Token     string `json:"token"`
}

func K8sSignHandler(serverCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := ApiResponse{}
		var request KubeConfigRequest
		if err := api_utils.Parse(r, &request); err != nil {
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		err := common.CheckUserHasPermission(r.Context(), request.Id)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		}
		token := utils.GetRandomString(20)
		sign := r.Header.Get("Signature")
		ttyParameter := cache.TtyParameter{
			Id:        request.Id,
			Arg:       fmt.Sprintf("pod=%s&container=%s&id=%d&sign=%s&namespace=%s", request.Pod, request.Container, request.Id, sign, request.Namespace),
			Sign:      sign,
			UserAgent: r.Header.Get("User-Agent"),
		}
		if err = serverCtx.Cache.Add(token, &ttyParameter, time.Duration(serverCtx.Config.TokenExpiresDuration)*time.Second); err != nil {
			logx.Infof("save token and ttyParam err:%s", err.Error())
			api_utils.ResponseStandWithError(w, r, err)
			return
		} else {
			result.Token = token
		}
		api_utils.ResponseStandSuccess(w, r, result)
	}
}

func TerminalHtmlHandler(serverCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		token := request.URL.Query().Get("token")
		if token == "" {
			api_utils.ResponseStandWithError(w, request, errors.New("鉴权失败"))
			return
		}

		ttyArgs := serverCtx.Cache.Get(token)
		defer func() {
			x := recover()
			if x != nil {
				logx.Errorf("%+v", x)
				api_utils.ResponseStandWithError(w, request, errors.New("鉴权失败"))
			}
		}()

		if ttyArgs == nil {
			api_utils.ResponseStandWithError(w, request, errors.New("鉴权失败"))
			return
		}
		query, err := url.ParseQuery(ttyArgs.Arg)
		if err != nil {
			api_utils.ResponseStandWithError(w, request, errors.New("鉴权失败"))
			return
		}
		us := common.GetUser(ttyArgs)
		if us == nil {
			api_utils.ResponseStandWithError(w, request, errors.New("鉴权失败"))
			return
		}
		ctx := context.WithValue(request.Context(), "user", *us)
		err = common.CheckUserHasPermission(ctx, ttyArgs.Id)
		if err != nil {
			api_utils.ResponseStandWithError(w, request, err)
			return
		}

		//request.WithContext(context.WithValue(request.Context(), "ttyArgs", ttyArgs))
		//request.Header.Set("Signature", ttyArgs.Sign)
		hd := func(writer http.ResponseWriter, request *http.Request) {
			indexData, err := Asset("static/terminal.html")
			if err != nil {
				panic("terminal static not found")
			}

			indexData = bytes.Replace(indexData, []byte("{{ .title }}"), []byte(fmt.Sprintf("ssh@root:%s@%s", query.Get("pod"), query.Get("container"))), -1)
			_, _ = writer.Write(indexData)
		}
		hd(w, request)
	}
}

var counter *ws.Counter

func init() {
	counter = ws.NewCounter(time.Duration(0) * time.Second)
}

func TerminalWsHandler(serverCtx *svc.ServiceContext) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		f := ws.NewWsServer(serverCtx, counter).GenerateHandleWS()
		f(writer, request)
	}
}
