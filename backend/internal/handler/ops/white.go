package ops

import (
	"github.com/lwl1989/kube-ssh/backend/internal/handler/api_utils"
	"github.com/lwl1989/kube-ssh/backend/internal/logic/ops"
	middleware "github.com/lwl1989/kube-ssh/backend/internal/middleware/types"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/lwl1989/kube-ssh/backend/internal/types"
	"net/http"
)

func WhiteUpsertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WhiteUpsertReq
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		l := ops.NewWhiteLogic(r.Context(), svcCtx)
		us := r.Context().Value("user").(middleware.OaUserInfo)
		err := l.Upsert(req.UserId, req.ClusterId, us.Id)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, nil)
		}
	}
}

func WhiteDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestWithId
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		l := ops.NewWhiteLogic(r.Context(), svcCtx)
		us := r.Context().Value("user").(middleware.OaUserInfo)
		err := l.Delete(req, us.Id)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, nil)
		}
	}
}

func WhiteStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestStatusChange
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		l := ops.NewWhiteLogic(r.Context(), svcCtx)
		us := r.Context().Value("user").(middleware.OaUserInfo)
		err := l.Status(req, us.Id)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, nil)
		}
	}
}

func WhiteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageCommonReq
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		l := ops.NewWhiteLogic(r.Context(), svcCtx)
		resp, err := l.List(req, r)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, resp)
		}
	}
}
