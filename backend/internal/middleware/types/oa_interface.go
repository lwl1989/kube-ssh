package types

import "net/http"

type IOaUser interface {
	GetUser(r *http.Request) (OaUserInfo, error)
	GetUserBySign(sign, userAgent string) (OaUserInfo, error)
	GetUserInfoByDepId(r *http.Request, depId int) (users []OaUserInfo, err error)
}
