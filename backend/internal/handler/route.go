package handler

import (
	"github.com/lwl1989/kube-ssh/backend/internal/handler/api_utils"
	"github.com/lwl1989/kube-ssh/backend/internal/handler/ops"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

var healthRes = []byte{'o', 'k'}

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/health",
		Handler: func(writer http.ResponseWriter, request *http.Request) {
			writer.Write(healthRes)
		},
	})
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/api/user/info",
		Handler: middleware.AuthMiddlewareObj.MiddleWare(func(w http.ResponseWriter, r *http.Request) {
			api_utils.ResponseStandSuccess(w, r, r.Context().Value("user"))
		}),
	})
	RegisterK8sManagerHandlers(server, serverCtx)
	RegisterStaticHandler(server, serverCtx)

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: ops.DirHandler(serverCtx),
	})
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/terminal/ws",
		Handler: ops.TerminalWsHandler(serverCtx),
	})
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/terminal/",
		Handler: ops.TerminalHtmlHandler(serverCtx),
	})
}
