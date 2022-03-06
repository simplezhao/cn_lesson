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

#### part1
```text
作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

- 优雅启动
- 优雅终止
- 资源需求和 QoS 保证
- 探活
- 日常运维需求，日志等级
- 配置和代码分离

```

#### part2
```text
除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧。
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）：

 - Service
 - Ingress
可以考虑的细节

 - 如何确保整个应用的高可用。
 - 如何通过证书保证 httpServer 的通讯安全。
```

