package model

import (
	"errors"
	"github.com/go-libraries/kube-manager/backend/internal/types/timex"
)

type UserWhite struct {
	Id         int            `gorm:"column:id;primaryKey;type:int(10)" json:"id"`
	UserId     int            `gorm:"column:user_id;not null;default:0;type:int(10)"  json:"user_id"`
	ClusterId  int            `gorm:"column:cluster_id;not null;default:0;type:int(10)"  json:"cluster_id"`
	OperatorId int            `gorm:"column:operator_id;not null;default:0;type:int(10)"  json:"operator_id"`
	Status     int            `gorm:"column:status;size:1;not null;default:1" json:"status"`
	Ctime      timex.UnixTime `gorm:"column:ctime;not null;default:0;type:int(10)" json:"ctime"`
	Mtime      timex.UnixTime `gorm:"column:mtime;not null;default:0;type:int(10)" json:"mtime"`
}

func UserWhiteTableName() string {
	return "user_white"
}

func (c UserWhite) TableName() string {
	return UserWhiteTableName()
}

func (c UserWhite) CheckCanUse() (bool, error) {
	if c.Id == 0 {
		return false, errors.New("未授权进入")
	}
	if c.Status != 1 {
		return false, errors.New("此授权已被禁用")
	}
	return true, nil
}

func (c UserWhite) Available() (bool, error) {
	return c.CheckCanUse()
}
