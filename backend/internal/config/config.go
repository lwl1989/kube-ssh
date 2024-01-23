package config

import (
	"github.com/go-libraries/ormx"
	"github.com/zeromicro/go-zero/rest"
	"time"
)

var GlobalConfig Config

const (
	DbDefault = "default"
)

type Config struct {
	rest.RestConf
	DB                   DBConfig
	Cors                 string `json:",default=*"` // 跨域配置 ,逗号拼接
	Env                  string `json:",default=dev"`
	TokenExpiresDuration int    `json:",default=300"`
	StaticDir            string `json:",default=dist"`
	MaxWsConnection      int    `json:",default=50"`
	WSOrigin             string `json:",default=*"`
	WhiteDepId           int    `json:",default=3247"`
	ItUserDepId          int    `json:",default=3195"`
	OaHost               string `json:",default=test.local"`
}

var TimeLocation *time.Location

func init() {
	TimeLocation, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = TimeLocation
}

type DBConfig struct {
	DefaultMysql ormx.MysqlConfig // mysql链接地址，满足 $user:$password@tcp($ip:$port)/$db?$queries 格式即可
}
