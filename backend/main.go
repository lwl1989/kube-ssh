package main

import (
	"flag"
	"fmt"
	"github.com/lwl1989/kube-ssh/backend/internal/config"
	"github.com/lwl1989/kube-ssh/backend/internal/handler"
	commonHandler "github.com/lwl1989/kube-ssh/backend/internal/library/handler"
	"github.com/lwl1989/kube-ssh/backend/internal/library/sys"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strings"
	"time"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	origins := strings.Split(c.Cors, ",")
	userServer := rest.MustNewServer(c.RestConf,
		rest.WithCustomCors(func(header http.Header) {
			middleware.ReSetAllowHeaders(header)
		}, nil, origins...),
	)
	defer userServer.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(userServer, ctx)
	commonHandler.SetErrorHandler()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//go rpcStart()
	SmoothStart()
	userServer.Start()

}

// SmoothStart 平滑启动
func SmoothStart() {
	// pid=$(ps -ef |grep ${MODULE} | grep ${ENV} | grep -v "grep" |awk '{print $2}')
	//if [ -z "$pid"  ]; then
	//  echo "no start"
	//else
	//  echo "kill old process ${pid}"
	//  kill -9 $pid
	//fi
	// 关闭之前监听的端口
	sys.KillPort(int64(config.GlobalConfig.Port))
	ticker := time.NewTicker(time.Millisecond * 50)
	killed := false
	nums := 0
	for {
		if killed || nums > 20 {
			break
		}
		select {
		case <-ticker.C:
			res := sys.Netstat(int64(config.GlobalConfig.Port))
			if len(res) == 0 {
				ticker.Stop()
				killed = true
			}
			nums++
		}
	}
}
