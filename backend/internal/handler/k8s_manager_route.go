package handler

import (
	"github.com/go-libraries/kube-manager/backend/internal/handler/md"
	"github.com/go-libraries/kube-manager/backend/internal/handler/ops"
	"github.com/go-libraries/kube-manager/backend/internal/middleware"
	"github.com/go-libraries/kube-manager/backend/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterK8sManagerHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/k8s/cluster",
				Handler: md.AuthMiddleWare(ops.ClusterUpsertHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/k8s/cluster",
				Handler: md.AuthMiddleWare(ops.ClusterDetailHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/k8s/clusters",
				Handler: middleware.AuthMiddlewareObj.MiddleWare(ops.ClusterListHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/k8s/workload/pods",
				Handler: middleware.AuthMiddlewareObj.MiddleWare(ops.WorkloadPodsHandler(serverCtx)),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/k8s/sign",
				Handler: middleware.AuthMiddlewareObj.MiddleWare(ops.K8sSignHandler(serverCtx)),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/manager/upsert",
				Handler: md.AuthManagerMiddleWare(ops.ManagerUpsertHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/managers",
				Handler: md.AuthManagerMiddleWare(ops.ManagerListHandler(serverCtx)),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/manager/status",
				Handler: md.AuthManagerMiddleWare(ops.ManagerStatusHandler(serverCtx)),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/white",
				Handler: md.AuthManagerMiddleWare(ops.WhiteUpsertHandler(serverCtx)),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/white",
				Handler: md.AuthManagerMiddleWare(ops.WhiteDeleteHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/whites",
				Handler: md.AuthManagerMiddleWare(ops.WhiteListHandler(serverCtx)),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/white/status",
				Handler: md.AuthManagerMiddleWare(ops.WhiteStatusHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/users",
				Handler: middleware.AuthMiddlewareObj.MiddleWare(ops.ItUsersHandler(serverCtx)),
			},
		},
	)
}
