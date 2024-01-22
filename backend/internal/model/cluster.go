package model

import (
	"errors"
	"github.com/go-libraries/kube-manager/backend/internal/types/timex"
)

type Cluster struct {
	Id         int            `gorm:"column:id;primaryKey;type:int(10)" json:"id"`
	ServerApi  string         `gorm:"column:server_api;size:64;not null;default:''" json:"server_api"`
	Name       string         `gorm:"column:name;unique;size:64;not null;default:''" json:"name"`
	NameSpaces string         `gorm:"column:namespaces;size:64;not null;default:''" json:"namespaces"`
	Config     string         `gorm:"column:config" json:"config"`
	OperatorId int            `gorm:"column:operator_id;not null;default:0;type:int(10)"  json:"operator_id"`
	Status     int            `gorm:"column:status;size:1;not null;default:1" json:"status"`
	Ctime      timex.UnixTime `gorm:"column:ctime;not null;default:0;type:int(10)" json:"ctime"`
	Mtime      timex.UnixTime `gorm:"column:mtime;not null;default:0;type:int(10)" json:"mtime"`
}

func ClusterTableName() string {
	return "k8s_cluster"
}

func (c Cluster) TableName() string {
	return ClusterTableName()
}

func (c Cluster) CheckCanUse() (bool, error) {
	if c.Id == 0 {
		return false, errors.New("集群信息查找失败")
	}
	if c.Status != 1 {
		return false, errors.New("此集群已被禁用")
	}
	if c.ServerApi == "" {
		return false, errors.New("k8s远程服务器未配置")
	}
	if c.Config == "" {
		return false, errors.New("k8s鉴权信息未配置")
	}
	return true, nil
}

func (c Cluster) Available() (bool, error) {
	return c.CheckCanUse()
}
