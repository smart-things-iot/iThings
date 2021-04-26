# api文件编译方法
```shell script
goctl api go -api user.api  -dir ./
```

# 数据库文件生成
```shell script
goctl model mysql datasource -url="root:password@tcp(127.0.0.1:3308)/pet" -table="*" -dir ./model -c
```

# rpc文件编译方法
```shell script
goctl rpc proto -src rpc/user.proto  -dir rpc
```

#接口文档生成