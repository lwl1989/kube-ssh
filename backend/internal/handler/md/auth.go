package md

import (
	"github.com/lwl1989/kube-ssh/backend/internal/config"
	"github.com/lwl1989/kube-ssh/backend/internal/handler/api_utils"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware/types"
	"github.com/lwl1989/kube-ssh/backend/internal/model"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/pkg/errors"
	"net/http"
)

func AuthMiddleWare(fn http.HandlerFunc) http.HandlerFunc {
	next := func(w http.ResponseWriter, r *http.Request) {
		if user, ok := r.Context().Value("user").(types.OaUserInfo); ok {
			if user.OrgId != svc.GlobalService.Config.WhiteDepId {
				var manager model.UserManager
				svc.GetDb(config.DbDefault).ReadWithContext(r.Context()).Table(model.UserManagerTableName()).Where("user_id=?", user.Id).First(&manager)
				ok, _ := manager.Available()
				if !ok {
					api_utils.ResponseStandWithError(w, r, errors.New("未授权进入"))
					return
				}
			}
		} else {
			api_utils.ResponseStandWithError(w, r, errors.New("未授权进入"))
			return
		}
		fn(w, r)
	}
	return middleware.AuthMiddlewareObj.MiddleWare(next)
}

func AuthManagerMiddleWare(fn http.HandlerFunc) http.HandlerFunc {
	next := func(w http.ResponseWriter, r *http.Request) {
		if user, ok := r.Context().Value("user").(types.OaUserInfo); ok {
			var manager model.UserManager
			svc.GetDb(config.DbDefault).ReadWithContext(r.Context()).Table(model.UserManagerTableName()).Where("user_id=?", user.Id).First(&manager)
			ok, err := manager.Available()
			if !ok {
				api_utils.ResponseStandWithError(w, r, err)
				return
			}
			if manager.Role != 1 {
				api_utils.ResponseStandWithError(w, r, errors.New("授权等级不够"))
				return
			}
		} else {
			api_utils.ResponseStandWithError(w, r, errors.New("未授权进入"))
			return
		}
		fn(w, r)
	}
	return middleware.AuthMiddlewareObj.MiddleWare(next)
}

func GetNowLoginUserId(r *http.Request) int {
	us, ok := r.Context().Value("user").(types.OaUserInfo)
	if ok {
		return us.Id
	}
	return 0
}
