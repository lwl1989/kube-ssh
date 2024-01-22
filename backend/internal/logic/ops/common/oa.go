package common

import (
	"github.com/go-libraries/kube-manager/backend/internal/cache/token"
	"github.com/go-libraries/kube-manager/backend/internal/middleware"
	"github.com/go-libraries/kube-manager/backend/internal/middleware/types"
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
