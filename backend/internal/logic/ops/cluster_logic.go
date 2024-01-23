package ops

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lwl1989/kube-ssh/backend/internal/logic/ops/common"
	middleware "github.com/lwl1989/kube-ssh/backend/internal/middleware/types"
	"github.com/lwl1989/kube-ssh/backend/internal/model"
	"github.com/lwl1989/kube-ssh/backend/internal/types/timex"
	"github.com/spf13/cast"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strings"

	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/lwl1989/kube-ssh/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClusterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterLogic {
	return &ClusterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterLogic) ClusterList(us middleware.OaUserInfo) (resp types.ClustersResponse, err error) {
	var clusters []model.Cluster
	q := l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(model.ClusterTableName())
	if l.svcCtx.Config.WhiteDepId != us.OrgId {
		var manager model.UserManager
		l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(manager.TableName()).Where("user_id=?", us.Id).First(&manager)
		if ok, _ := manager.Available(); !ok {
			var cid []int
			l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(model.UserWhiteTableName()).Where("user_id=? and status=?", us.Id, 1).Pluck("cluster_id", &cid)
			if len(cid) == 0 {
				return
			}
			q.Where("id in (?)", cid)
		}
	}
	q.Find(&clusters)
	for _, v := range clusters {
		resp.List = append(resp.List, types.ClusterItem{
			Id:         v.Id,
			Name:       v.Name,
			NameSpaces: v.NameSpaces,
			ClusterId:  cast.ToString(v.Id),
			HasConfig:  v.Config != "",
			Server:     v.ServerApi,
			Status:     v.Status,
			Ctime:      v.Ctime,
			Mtime:      v.Mtime,
		})
	}
	return
}

func (l *ClusterLogic) WorkloadPods(req types.RequestWithId) (resp types.PodListRes, err error) {
	var cluster model.Cluster
	l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(model.ClusterTableName()).Where("id=?", req.Id).First(&cluster)
	var ok bool
	ok, err = cluster.CheckCanUse()
	if !ok {
		return
	}
	err = common.CheckUserHasPermission(l.ctx, cluster.Id)
	if err != nil {
		return
	}
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(cluster.Config))
	if err != nil {
		return resp, err
	}

	logx.Infof("use config with %+v", config)
	// Create an rest client not targeting specific API version
	clientRest, err := kubernetes.NewForConfig(config)
	if err != nil {
		return resp, err
	}

	namespaces := strings.Split(cluster.NameSpaces, ",")
	logx.Infof("load data from namespace %+v", namespaces)
	var pods []v1.Pod
	for _, namespace := range namespaces {
		ps, err := clientRest.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			return resp, err
		}
		//test
		//for i, v := range ps.Items {
		//	v.Status.ContainerStatuses = append(v.Status.ContainerStatuses, v.Status.ContainerStatuses...)
		//
		//	v.Spec.Containers = append(v.Spec.Containers, v.Spec.Containers...)
		//	ps.Items[i] = v
		//}
		pods = append(pods, ps.Items...)
	}

	if err != nil {
		return resp, err
	}
	bts, err := json.Marshal(map[string][]v1.Pod{
		"items": pods,
	})
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(bts, &resp)
	return
}

func (l *ClusterLogic) ClusterUpsert(req types.ClusterItemDetail, operatorId int) error {
	now := timex.UnixTimeNow()
	if req.Id > 0 {
		return l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Table(model.ClusterTableName()).
			Where("id=?", req.Id).Updates(map[string]any{
			"name":        req.Name,
			"namespaces":  req.NameSpace,
			"mtime":       now.Unix(),
			"operator_id": operatorId,
			"config":      req.Config,
			"server_api":  req.ServerApi,
		}).Error
	}
	return l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Create(&model.Cluster{
		ServerApi:  req.ServerApi,
		Name:       req.Name,
		NameSpaces: req.NameSpace,
		Config:     req.Config,
		OperatorId: operatorId,
		Status:     1,
		Ctime:      now,
		Mtime:      now,
	}).Error
}

func (l *ClusterLogic) ClusterStatus(req types.StatusChangeReq, operatorId int) (err error) {
	return l.svcCtx.DefaultDb.WriteWithContext(l.ctx).Table(model.ClusterTableName()).
		Where("id=?", req.Id).Updates(map[string]any{
		"status":      req.Status,
		"operator_id": operatorId,
	}).Error
}

func (l *ClusterLogic) ClusterDetail(req types.RequestWithId) (resp *types.ClusterItemDetail, err error) {
	var cluster model.Cluster
	l.svcCtx.DefaultDb.ReadWithContext(l.ctx).Table(cluster.TableName()).Where("id=?", req.Id).First(&cluster)
	if cluster.Id == 0 {
		return nil, errors.New("集群信息查找失败")
	}
	return &types.ClusterItemDetail{
		Id:         cluster.Id,
		Name:       cluster.Name,
		ServerApi:  cluster.ServerApi,
		NameSpace:  cluster.NameSpaces,
		Config:     cluster.Config,
		OperatorId: cluster.OperatorId,
		Status:     cluster.Status,
	}, nil
}
