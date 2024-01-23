package middleware

import (
	"context"
	"github.com/lwl1989/kube-ssh/backend/internal/handler/api_utils"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware/mock"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware/types"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type AuthOaMiddleware struct {
	UserCall types.IOaUser
}

var AuthMiddlewareObj *AuthOaMiddleware

func SetOaCall(uc types.IOaUser) {
	AuthMiddlewareObj.UserCall = uc
}

func init() {
	AuthMiddlewareObj = new(AuthOaMiddleware)
	SetOaCall(mock.OaMock{})
}

func (auth *AuthOaMiddleware) MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//todo: 改造此处
		user, err := auth.UserCall.GetUser(r)
		if err != nil {
			logx.WithContext(r.Context()).Errorf("GetUser error: %s", err.Error())
			api_utils.ResponseStandWithError(w, r, err)
			return
		}
		next(w, r.WithContext(context.WithValue(r.Context(), "user", user)))
	}
}

func (auth *AuthOaMiddleware) GetUserInfoByDepId(r *http.Request, depId int) []types.OaUserInfo {
	us, err := auth.UserCall.GetUserInfoByDepId(r, depId)
	if err != nil {
		logx.WithContext(r.Context()).Errorf("GetUserInfoByDepId error: %s", err.Error())
		return nil
	}
	return us
}
