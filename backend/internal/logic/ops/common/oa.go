package common

import (
	"github.com/lwl1989/kube-ssh/backend/internal/cache/token"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware/types"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetUser(tty *token.TtyParameter) *types.OaUserInfo {
	tUs, err := middleware.AuthMiddlewareObj.UserCall.GetUserBySign(tty.Sign, tty.UserAgent)
	if err != nil {
		logx.Errorf("tty getOaUser fail:%s", err.Error())
		return nil
	}
	return &tUs
}
