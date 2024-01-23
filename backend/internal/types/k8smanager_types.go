// Code generated by goctl. DO NOT EDIT.
package types

import "github.com/lwl1989/kube-ssh/backend/internal/types/timex"

type ClusterItem struct {
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	NameSpaces string         `json:"namespaces"`
	ClusterId  string         `json:"cluster_id"`
	HasConfig  bool           `json:"has_config"` //是否配置过配置文件
	Server     string         `json:"server"`     //管理端地址（内网）
	Status     int            `json:"status"`
	Ctime      timex.UnixTime `gorm:"column:ctime" json:"ctime"`
	Mtime      timex.UnixTime `gorm:"column:mtime" json:"mtime"`
}

type ClusterItemDetail struct {
	Id         int    `json:"id"` //大于0为编辑
	Name       string `json:"name"`
	ServerApi  string `json:"server_api"`
	NameSpace  string `gorm:"column:namespaces" json:"namespaces"`
	Config     string `gorm:"column:config" json:"config"`
	OperatorId int    `gorm:"column:operator_id"  json:"operator_id"`
	Status     int    `json:"status" gorm:"column:status"`
}

type ClustersResponse struct {
	List []ClusterItem `json:"list"`
}

type PodItem struct {
	Name     string `json:"name"`
	Restarts int    `json:"restarts"`
	Status   string `json:"status"`
	Ready    string `json:"ready"`
}

type PodsResponse struct {
	List []PodItem `json:"list"`
}

type RequestWithId struct {
	Id int `json:"id" validate:"required,gt=0"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ListResponse struct {
	List  []any `json:"list"`
	Count int64 `json:"count"`
}

type RequestStatusChange struct {
	Id     int `json:"id" validate:"required,gt=0"`
	Status int `json:"status"`
}

type ManagerUpsertReq struct {
	UserId int `json:"user_id"`
	Role   int `json:"role"`
}

type WhiteUpsertReq struct {
	UserId    int `json:"user_id"`
	ClusterId int `json:"cluster_id"`
}

type StatusChangeReq struct {
	Id     int `json:"id" validate:"required,gt=0"`
	Status int `json:"status,gt=0"`
}
