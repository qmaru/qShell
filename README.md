# qShell

基于 gRPC 的远程脚本执行工具，降低 SSH 的暴露风险，将执行脚本集中在固定目录下，通过客户端发送参数调用。

## 基本功能

+ Ping 可用性测试
+ List 列出所有脚本
+ Run 执行脚本
+ Upload 上传文件

## 使用

在根目录创建 users.json 文件并添加用户和密码

```shell
{
    "username": "password"
}
```

启动服务

```shell
# 编译
go build

# 运行 server
qShell server

# 测试连接
qShell client ping
```

服务端

```shell
server control

Usage:
  qShell server [flags]

Flags:
  -h, --help              help for server
      --host string       Listen host (default "127.0.0.1")
  -p, --password string   Password
      --port string       Listen port (default "52500")
  -u, --username string   Username
```

客户端

```shell
client control

Usage:
  qShell client [flags]
  qShell client [command]

Available Commands:
  list        List all scripts
  ping        Ping Server
  run         Run a scripts
  upload      Upload file

Flags:
  -h, --help              help for client
      --host string       Listen host (default "127.0.0.1")
  -p, --password string   Password
      --port string       Listen port (default "52500")
      --timeout int       Timeout (default 15)
  -u, --username string   Username
```
