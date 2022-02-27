## 目录内容
### lesson2
[模块2作业](./lesson2)
```
1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
```

### lesson3
[模块3作业](./lesson3)
```
1. 构建本地镜像
2. 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
3. 将镜像推送至 docker 官方镜像仓库
4. 通过 docker 命令本地启动 httpserver
5. 通过 nsenter 进入容器查看 IP 配置
```

### lesson8
[模块8作业](./lesson8)
作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

- 优雅启动
- 优雅终止
- 资源需求和 QoS 保证
- 探活
- 日常运维需求，日志等级
- 配置和代码分离
