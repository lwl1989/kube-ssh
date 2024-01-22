package ops

import (
	"github.com/go-libraries/kube-manager/backend/internal/handler/md"
	middleware "github.com/go-libraries/kube-manager/backend/internal/middleware/types"
	"github.com/go-libraries/kube-manager/backend/internal/types"
	"net/http"

	"github.com/go-libraries/kube-manager/backend/internal/handler/api_utils"
	"github.com/go-libraries/kube-manager/backend/internal/logic/ops"
	"github.com/go-libraries/kube-manager/backend/internal/svc"
)

func ClusterListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := ops.NewClusterLogic(r.Context(), svcCtx)
		us := r.Context().Value("user").(middleware.OaUserInfo)
		resp, err := l.ClusterList(us)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, resp)
		}
	}
}

func WorkloadPodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestWithId
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		l := ops.NewClusterLogic(r.Context(), svcCtx)
		resp, err := l.WorkloadPods(req)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, resp)
		}
	}
}

func ClusterUpsertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClusterItemDetail
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, nil, err)
			return
		}

		l := ops.NewClusterLogic(r.Context(), svcCtx)
		err := l.ClusterUpsert(req, md.GetNowLoginUserId(r))
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, nil)
		}
	}
}

func ClusterDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestWithId
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, nil, err)
			return
		}

		l := ops.NewClusterLogic(r.Context(), svcCtx)
		resp, err := l.ClusterDetail(&req)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, resp)
		}
	}
}
