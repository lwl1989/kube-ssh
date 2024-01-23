package ops

import (
	"context"
	"github.com/lwl1989/kube-ssh/backend/internal/library/array"
	"github.com/lwl1989/kube-ssh/backend/internal/logic/ops/common"
	"github.com/lwl1989/kube-ssh/backend/internal/model"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/lwl1989/kube-ssh/backend/internal/types"
	"github.com/lwl1989/kube-ssh/backend/internal/types/timex"
	"github.com/pkg/errors"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type WhiteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWhiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WhiteLogic {
	return &WhiteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type WhiteListItem struct {
	model.UserWhite
	ClusterName string `json:"cluster"`
	Name        string `json:"name"`
	Namespaces  string `json:"namespaces"`
}

func (l *WhiteLogic) List(req types.PageCommonReq, r *http.Request) (resp types.ListResponse, err error) {
	var mas []model.UserWhite
	q := l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Limit(req.GetLimit()).Offset(req.GetOffset()).Table(model.UserWhiteTableName())
	q.Count(&resp.Count)
	q.Find(&mas)
	if len(mas) == 0 {
		return
	}
	clusterIds := array.GetItemValues[int, model.UserWhite](mas, "ClusterId")
	var clus []model.Cluster
	l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(model.ClusterTableName()).Where("id in (?)", clusterIds).Find(&clus)
	users := common.GetItUserList(r)
	for _, v := range mas {
		var item = WhiteListItem{
			UserWhite: v,
		}
		for _, us := range users {
			if v.UserId == us.Id {
				item.Name = us.Nickname
				break
			}
		}
		for _, c1 := range clus {
			if c1.Id == v.ClusterId {
				if ok, _ := c1.CheckCanUse(); !ok {
					item.ClusterName = "集群已禁用"
					break
				}
				item.ClusterName = c1.Name
				item.Namespaces = c1.NameSpaces
				break
			}
		}
		resp.List = append(resp.List, item)
	}
	return
}

func (l *WhiteLogic) Upsert(userId int, clusterId int, operatorId int) (err error) {
	var white model.UserWhite
	l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(white.TableName()).Where("user_id=?", userId).First(&white)
	if white.Id == 0 {
		err = l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Create(&model.UserWhite{
			UserId:     userId,
			ClusterId:  clusterId,
			OperatorId: operatorId,
			Status:     1,
			Ctime:      timex.UnixTimeNow(),
			Mtime:      timex.UnixTimeNow(),
		}).Error
	} else {
		err = l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Table(white.TableName()).Where("id=?", white.Id).Updates(map[string]any{
			"cluster_id":  clusterId,
			"operator_id": operatorId,
		}).Error
	}
	return
}

func (l *WhiteLogic) Status(req types.RequestStatusChange, operatorId int) (err error) {
	var white model.UserWhite
	l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(white.TableName()).Where("id=?", req.Id).First(&white)
	if white.Id == 0 {
		return errors.New("白名单配置id错误")
	}
	err = l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Table(model.UserManagerTableName()).Where("id=?", white.Id).Updates(map[string]any{
		"status":      req.Status,
		"operator_id": operatorId,
	}).Error
	return
}

func (l *WhiteLogic) Delete(req types.RequestWithId, operatorId int) (err error) {
	err = l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Table(model.UserWhiteTableName()).Where("id=?", req.Id).Delete(&model.UserWhite{}).Error
	return
}
