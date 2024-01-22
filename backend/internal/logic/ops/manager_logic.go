package ops

import (
	"context"
	"github.com/go-libraries/kube-manager/backend/internal/logic/ops/common"
	OaTypes "github.com/go-libraries/kube-manager/backend/internal/middleware/types"
	"github.com/go-libraries/kube-manager/backend/internal/model"
	"github.com/go-libraries/kube-manager/backend/internal/svc"
	"github.com/go-libraries/kube-manager/backend/internal/types"
	"github.com/go-libraries/kube-manager/backend/internal/types/timex"
	"github.com/pkg/errors"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManagerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManagerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManagerLogic {
	return &ManagerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type managerListItem struct {
	model.UserManager
	Name string `json:"name"`
}

func (l *ManagerLogic) List(req types.PageCommonReq, r *http.Request) (resp types.ListResponse, err error) {
	var mas []model.UserManager
	q := l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Limit(req.GetLimit()).Offset(req.GetOffset()).Table(model.UserManagerTableName())
	q.Count(&resp.Count)
	q.Find(&mas)
	if len(mas) == 0 {
		return
	}
	users := common.GetItUserList(r)
	for _, v := range mas {
		var u OaTypes.OaUserInfo
		for _, us := range users {
			if v.UserId == us.Id {
				u = us
				break
			}
		}
		resp.List = append(resp.List, managerListItem{
			UserManager: v,
			Name:        u.Nickname,
		})
	}
	return
}

func (l *ManagerLogic) Upsert(userId int, role int, operatorId int) (err error) {
	var manager model.UserManager
	l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(model.UserManagerTableName()).Where("user_id=?", userId).First(&manager)
	if manager.Id == 0 {
		err = l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Table(model.UserManagerTableName()).Create(&model.UserManager{
			UserId:     userId,
			Role:       role,
			OperatorId: operatorId,
			Status:     1,
			Ctime:      timex.UnixTimeNow(),
			Mtime:      timex.UnixTimeNow(),
		}).Error
	} else {
		err = l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Table(model.UserManagerTableName()).Where("id=?", manager.Id).Updates(map[string]any{
			"role":        role,
			"operator_id": operatorId,
		}).Error
	}

	return
}

func (l *ManagerLogic) Status(req types.RequestStatusChange, operatorId int) (err error) {
	var manager model.UserManager
	l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(model.UserManagerTableName()).Where("id=?", req.Id).First(&manager)
	if manager.Id == 0 {
		return errors.New("此管理员不存在")
	}
	err = l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Table(model.UserManagerTableName()).Where("id=?", manager.Id).Updates(map[string]any{
		"status":      req.Status,
		"operator_id": operatorId,
	}).Error
	return
}
