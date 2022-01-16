在mac系统下通过`brew install util-linux`是没有nsenter命令的
但是可以通过docker的方式获取nsenter

1. 拉取镜像
    `docker pull justincormack/nsenter1`
2. 指定--pid=host获取主机进程，并指定--privileged 提升特权（这步操作谨慎用于生产环境）
    ```bash
   docker run --pid=host --privileged -it --rm justincormack/nsenter1
    ```
   这时会进入nsenter镜像
3. 在新的terminal启动httpserver
    ```bash
    docker run -p 80:32000 --name httpserver --rm simplezhao/go-httpserver:0.3
   ```
4. 获取容器的PID
    ```bash
    docker inspect -f '{{.State.Pid}}' httpserver
   ```
   比如返回1961
5. 通过nsenter进入PID为1961的进程
    ```bash
    nsenter -t 1961 -m -u -n -i bash
   ```
   返回
   ```bash
    1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
        link/ipip 0.0.0.0 brd 0.0.0.0
    3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
        link/tunnel6 :: brd ::
    8: eth0@if9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
        link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
        inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
        valid_lft forever preferred_lft forever
   ```
6. 通过`docker inspect httpserver | grep IPAddress`
   IP跟nsenter获取一致

![运行httpserver](https://oss.smart-lifestyle.cn/file/1mbmk.png)

![nsenter查看IP](https://oss.smart-lifestyle.cn/file/s7yf4.png)