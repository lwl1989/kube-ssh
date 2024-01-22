package svc

import (
	"github.com/go-libraries/kube-manager/backend/internal/cache/token"
	"github.com/go-libraries/kube-manager/backend/internal/config"
	"github.com/go-libraries/kube-manager/backend/internal/library/tracex"
	"github.com/go-libraries/ormx"
	"github.com/zeromicro/go-zero/core/proc"
)

type ServiceContext struct {
	Config    config.Config
	DefaultDb *ormx.DbSplit
	Cache     token.Cache
}

var dbMap map[string]*ormx.DbSplit
var GlobalService *ServiceContext

func init() {
	GlobalService = &ServiceContext{}
	GlobalService.Cache = token.NewMemCache()
	dbMap = make(map[string]*ormx.DbSplit)
}

// GetDb 替换几率极低 不需要上锁
func GetDb(key string) *ormx.DbSplit {
	if v, ok := dbMap[key]; ok {
		return v
	}
	panic("no found this key with db:" + key)
}

func GetService() *ServiceContext {
	return GlobalService
}

func NewServiceContext(c config.Config) *ServiceContext {
	config.GlobalConfig = c
	dbInstance := ormx.NewDbSplit(c.DB.DefaultMysql)
	dbMap[config.DbDefault] = dbInstance
	dbInstance.GetLogger().SetTraceFunc(tracex.OrmTrace)
	GlobalService.Config = c
	GlobalService.DefaultDb = dbMap[config.DbDefault]
	proc.AddShutdownListener(func() {
		if dbInstance != nil {
			_ = dbInstance.Close()
		}
	})
	return GlobalService
}

// NewServiceContextNoDb 用于不需要连接db的单元测试
func NewServiceContextNoDb(c config.Config) *ServiceContext {
	config.GlobalConfig = c
	GlobalService.Config = c
	return GlobalService
}
