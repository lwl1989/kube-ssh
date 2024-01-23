# 多集群web-kubectl

# 初始化

## go项目init

```go
cd backend
go mod tidy
```

## nodejs

```shell
cd front
npm install
npm run build:prod
```

## 手动覆盖前端页面

cp -rf front/dist/* backend/dist/

## 导入db(测试数据)
```go
cd backend/script/init
go run main.go
```

## sso对接

实现backend/internal/middleware/types/IOaUser接口,并调用backend/internal/middleware.SetOa()
```go
type IOaUser interface {
	GetUser(r *http.Request) (OaUserInfo, error)
	GetUserBySign(sign, userAgent string) (OaUserInfo, error)
	GetUserInfoByDepId(r *http.Request, depId int) (users []OaUserInfo, err error)
}
```

# 效果预览

![集群列表](https://www.github.com/lwl1989/kube-ssh/docs/clusters.png)

![pod列表](https://www.github.com/lwl1989/kube-ssh/docs/pods.png)

![命令行UI](https://www.github.com/lwl1989/kube-ssh/docs/terminal.png)

![img.png](https://www.github.com/lwl1989/kube-ssh/docs/manager.png)