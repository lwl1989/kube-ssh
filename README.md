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