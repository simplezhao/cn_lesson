### 要求
1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200

### Server
`http://127.0.0.1:32000`
### 接口
1. `/` home
2. `/index` 实现要求1， 2， 3
3. `/healthz` 实现要求4

### 本地运行
将在本地32000端口上运行
- make build
- make run

### 使用docker运行
#### 自行构建
```bash
# x86/amd64
docker buildx build --platform=linux/amd64   -t go-httpserver:latest .
docker run -p 80:32000 --name httpserver --rm go-httpserver:latest
```
#### 使用已有镜像
```bash
docker pull simplezhao/go-httpserver:0.3
docker run -p 80:32000 --name httpserver --rm simplezhao/go-httpserver:0.3
```