package common

import (
	"context"
	"encoding/json"
	"github.com/lwl1989/kube-ssh/backend/internal/config"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware"
	"github.com/lwl1989/kube-ssh/backend/internal/middleware/types"
	"github.com/lwl1989/kube-ssh/backend/internal/model"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/pkg/errors"
	"net/http"
	"sync"
	"time"
)

func CheckUserHasPermission(ctx context.Context, clusterId int) (err error) {
	v := ctx.Value("user")
	us := v.(types.OaUserInfo)
	var manager model.UserManager
	svc.GetDb(config.DbDefault).Table(manager.TableName()).Where("user_id=?", us.Id).First(&manager)
	var ok bool
	ok, err = manager.Available()
	if ok {
		return nil
	}
	return CheckInWhiteList(ctx, clusterId)
}

func GetCLusterWhiteList(ctx context.Context) (cids []int, needCheck bool, err error) {
	us := ctx.Value("user").(types.OaUserInfo)

	if svc.GlobalService.Config.WhiteDepId != us.DeptId {
		svc.GetDb(config.DbDefault).Table(model.UserWhiteTableName()).Where("user_id=? and status=?", us.Id, 1).Pluck("cluster_id", &cids)
		if len(cids) == 0 {
			return nil, true, errors.New("error is nil")
		}
		return cids, true, nil
	}
	return nil, false, nil
}

func CheckInWhiteList(ctx context.Context, cid int) (err error) {
	cids, check, err := GetCLusterWhiteList(ctx)
	if err != nil {
		return err
	}
	if check {
		for _, v := range cids {
			if v == cid {
				return nil
			}
		}
		return errors.New("未授权进入")
	}
	return nil
}

var userMu *sync.RWMutex

func init() {
	userMu = &sync.RWMutex{}
}

func GetItUserList(r *http.Request) (users []types.OaUserInfo) {
	userMu.RLock()
	strUser := svc.GlobalService.Cache.GetValue("itUsers")
	userMu.RUnlock()
	if len(strUser) > 0 {
		_ = json.Unmarshal([]byte(strUser), &users)
		return users
	}
	userMu.Lock()
	defer userMu.Unlock()
	users = middleware.AuthMiddlewareObj.GetUserInfoByDepId(r, svc.GlobalService.Config.ItUserDepId)
	_ = svc.GlobalService.Cache.AddKeyValue("itUsers", users, time.Second*3600)
	return users
}
