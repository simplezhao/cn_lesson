### 要求
1. 构建本地镜像
2. 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
3. 将镜像推送至 docker 官方镜像仓库
4. 通过 docker 命令本地启动 httpserver
5. 通过 nsenter 进入容器查看 IP 配置

### Server
`http://127.0.0.1:32000`
### 接口
参考lesson2

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

### nsenter
[nsenter使用](./nsenter使用.md)