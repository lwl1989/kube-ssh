package model

import (
	"errors"
	"github.com/lwl1989/kube-ssh/backend/internal/types/timex"
)

type UserManager struct {
	Id         int            `gorm:"column:id;primaryKey;type:int(10)" json:"id"`
	UserId     int            `gorm:"column:user_id;not null;default:0;type:int(10)"  json:"user_id"`
	Role       int            `gorm:"column:role;size:1;not null;default:1"  json:"role"`
	OperatorId int            `gorm:"column:operator_id;not null;default:0;type:int(10)"  json:"operator_id"`
	Status     int            `json:"status;size:1;not null;default:1" gorm:"column:status"`
	Ctime      timex.UnixTime `gorm:"column:ctime;not null;default:0;type:int(10)" json:"ctime"`
	Mtime      timex.UnixTime `gorm:"column:mtime;not null;default:0;type:int(10)" json:"mtime"`
}

func UserManagerTableName() string {
	return "user_manager"
}

func (c UserManager) TableName() string {
	return UserManagerTableName()
}

func (c UserManager) Available() (bool, error) {
	if c.Id == 0 {
		return false, errors.New("非管理员无法操作")
	}
	if c.Status != 1 {
		return false, errors.New("此管理员已被禁用")
	}
	return true, nil
}
