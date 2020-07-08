# go-kit example
> 从官网提取出来的 go-kit 演示项目


```bash
# 1. 通过 docker-compose 启动 consule
docker-compose up -d

# 2. 执行服务和 api 网关
go-kit-example/cmd/addsvc/cmd/addsvc -debug.addr :7080 -http-addr :7081 -grpc-addr :7082 -thrift-addr :7083 -jsonrpc-addr :7084
go-kit-example/cmd/addsvc/cmd/stringsvc -listen :8081
go-kit-example/cmd/addsvc/cmd/apigateway -consul.addr "127.0.0.1:8500"

# 3. 相关接口信息在 tests/ 文件夹下
```