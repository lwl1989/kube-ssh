package ops

import (
	"github.com/lwl1989/kube-ssh/backend/internal/handler/api_utils"
	"github.com/lwl1989/kube-ssh/backend/internal/logic/ops"
	"github.com/lwl1989/kube-ssh/backend/internal/logic/ops/common"
	middleware "github.com/lwl1989/kube-ssh/backend/internal/middleware/types"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/lwl1989/kube-ssh/backend/internal/types"
	"net/http"
)

func ManagerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageCommonReq
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		l := ops.NewManagerLogic(r.Context(), svcCtx)
		resp, err := l.List(req, r)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, resp)
		}
	}
}

func ManagerUpsertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ManagerUpsertReq
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		l := ops.NewManagerLogic(r.Context(), svcCtx)
		us := r.Context().Value("user").(middleware.OaUserInfo)
		err := l.Upsert(req.UserId, req.Role, us.Id)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, nil)
		}
	}
}

func ManagerStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestStatusChange
		if err := api_utils.Parse(r, &req); err != nil {
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		l := ops.NewManagerLogic(r.Context(), svcCtx)
		us := r.Context().Value("user").(middleware.OaUserInfo)
		err := l.Status(req, us.Id)
		if err != nil {
			api_utils.ResponseStandWithError(w, r, err)
		} else {
			api_utils.ResponseStandSuccess(w, r, nil)
		}
	}
}

func ItUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := common.GetItUserList(r)
		api_utils.ResponseStandSuccess(w, r, users)
	}
}
