package mock

import (
	"github.com/lwl1989/kube-ssh/backend/internal/middleware/types"
	"net/http"
)

type OaMock struct {
}

// GetUser 当前请求进行sso鉴权 获取用户信息
func (oa OaMock) GetUser(r *http.Request) (types.OaUserInfo, error) {
	return types.OaUserInfo{
		Avatar:   "",
		DeptId:   3195,
		Email:    "",
		Id:       781,
		Mobile:   "",
		Nickname: "admin",
		OrgId:    3195,
	}, nil
}

func (oa OaMock) GetUserInfoByDepId(r *http.Request, depId int) (users []types.OaUserInfo, err error) {
	return []types.OaUserInfo{
		{
			Avatar:   "",
			DeptId:   depId,
			Email:    "",
			Id:       782,
			Mobile:   "",
			Nickname: "admin",
			OrgId:    3195,
		},
		{
			Avatar:   "",
			DeptId:   depId,
			Email:    "",
			Id:       781,
			Mobile:   "",
			Nickname: "admin",
			OrgId:    3195,
		},
	}, nil
}

func (oa OaMock) GetUserBySign(sign, userAgent string) (types.OaUserInfo, error) {
	return types.OaUserInfo{
		Avatar:   "",
		DeptId:   3195,
		Email:    "",
		Id:       781,
		Mobile:   "",
		Nickname: "admin",
		OrgId:    3195,
	}, nil
}
