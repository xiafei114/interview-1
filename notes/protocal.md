1. 建立连接, 是客户端和服务端为了维护状态，而建立一定的数据结构来维护双方交互的状态



**udp 协议**

  特点: 

- UDP继承IP的特性，	不保证不丢失，不保证数据顺序到达
- 基于数据报发送，一个一个发，一个一个接收
- 无拥塞控制
- 无状态服务

![img](https://static001.geekbang.org/resource/image/6d/bf/6d1313f51b9dfd7ab454b2cef1cb37bf.jpg)

应用场景:

	- 需要的资源少, 在网络情况比较好的内网，或者对丢包不敏感的应用
	- 不需要建立一对一连接， 广播应用
	- 处理速度快，时延低，容许一定程度的丢包，网络拥塞时，也要正常的流速发送数据报

实际应用:

	- 网页或者APP访问(QUIC Quick UDP Internet Connection)
	- 流媒体(直播类)
	- 实时游戏, 一般情况下游戏需要较为严格的数据接收顺序和数据的完整性, 但是由于TCP协议需要在服务器上维护连接的数据结构，在大数据量下内存容量有效，因此会有些场景下需要自定义UDP协议，实现重传等策略来保证数据的完整以及降低数据传输时延
	- IoT物联网, 物联网芯片经常是内存容量较小的嵌入式芯片，没有足够的内存维护TCP协议的数据结构
	- 移动通信领域



思考题:

1. TCP是面向连接的, 对于计算机来说，怎么样才算是一个连接呢?

   连接: 自己监听的端口接收到来自远端的连接请求, 从数据包中解析出IP,端口等数据信息来维护一份数据结构来表示同一份来自远端的请求。

   断开: 四次挥手后确保对方的连接已经断开,删除本地维护的远端的数据结构





**TCP协议**

- 面向字节流
- 提供可靠交付
- 拥塞控制
- 有状态服务
- 流量控制

![img](https://static001.geekbang.org/resource/image/a7/bf/a795461effcce686a43f48e094c9adbf.jpg)

双方建立连接后相互传输的序号不是顺序序号，是随着时间变化的，可以看成一个32位的计数器，每4ms+1，每4个多小时有可能会发生重复, 某一方在这个时间之后也不会接受这个异常包



**TCP三次握手**

> TTL:time to live 数据包在网络中的存活时间, 最大值为255, 一般设置为64, 实际上TTL是IP数据包在计算机网络中可以转发的最大跳数, 如果在某个路由器上接收到TTL=0的数据报,将会向发送者发送ICMP time exceeded的消息
>
> linux 修改ttl值：
>
> 1. /proc/sys/net/ipv4/ip_default_ttl
> 2. 在 /etc/sysctl.conf 中增加一行 net.ipv4.ip_default_ttl = X



![img](https://static001.geekbang.org/resource/image/66/a2/666d7d20aa907d8317af3770411f5aa2.jpg)

一开始，双方服务器都处于CLOSED的状态，客户机想服务器发送SYN请求之后，处于SYN-SENT的状态, 服务端接收到SYN的请求后，状态转变为SYN_RCVD的状态，同时返回SYN + ACK的请求,客户机接收到消息之后建立ESTABLISHED的状态，返回ACK, 服务器状态也转变为ESTABLISHED状态，双方开始通信

**TCP四次挥手**

![img](https://static001.geekbang.org/resource/image/1f/11/1f6a5e17b34f00d28722428b7b8ccb11.jpg)

当客户端在FIN-WAIT-1状态并收到服务器发来的ACK时,状态转移到FIN-WAIT-2, 假如此时服务器不发送FIN+ACK的信号, 客户端会一直卡在FIN-WAIT2的状态,linux下可通过调整`tcp_fin_timeout`的参数设置超时时间. 

当客户端在FIN-WAIT2下收到了FIN+ACK的信号, 像服务器发送ACK的信号, 同时状态转移称TIME-WAIT(等待2MSL), 该状态防止服务器收不到ACK的信号时,服务端会重新发送FIN+ACK的信号，A不能直接关闭还有个原因在于假如A的端口直接关闭, 并被其他应用占用，服务器如果还有包没到达会导致异常

> MSL: Maxinum Segment Lifetime 报文最大生存时间

![img](https://static001.geekbang.org/resource/image/da/ab/dab9f6ee2908b05ed6f15f3e21be88ab.jpg)

在这个图中，加黑加粗的部分，是上面说到的主要流程，其中阿拉伯数字的序号，是连接过程中的顺序，而大写中文数字的序号，是连接断开过程中的顺序。加粗的实线是客户端 A 的状态变迁，加粗的虚线是服务端 B 的状态变迁



思考题:

1.TCP 的连接有这么多的状态，你知道如何在系统中查看某个连接的状态吗？

`netstat -an | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'`



**TCP累计应答/确认**

发送端

![img](https://static001.geekbang.org/resource/image/16/7b/16dcd6fb8105a1caa75887b5ffa0bd7b.jpg)

第一部分: 发送已确认, 接受方接收到消息，并返回ack确认(等待删除)

第二部分: 发送未确认，发送方消息发出，不确认接收方是否已经接收成功

第三部分: 未发送可发送, 根绝接收方的接收能力，数据已经准备好，随时可以发送

第四部分, 接收方接收能力已经达到了上限,无法继续接收

> 评估工作能力，避免工作量过大，无法完成工作或者工作过于轻松导致工作不饱和



接收端:

![img](https://static001.geekbang.org/resource/image/f7/a4/f7b1d3bc6b6d8e55f0951e82294c8ba4.jpg)

**超时重传**

快速重传: 当接收方发现丢了一个中间包后，发送三次前一个包的ACK, 发送端就不回等待超时重传，快速重传丢失的包

**流量控制**: 防止网络流量过大, 接收方处理不过来发送的消息(接收方缓存空间塞满)

**拥塞控制**: 

慢启动, 拥塞避免, bbr优化

tcp开始发送报文时,cwnd=1, 发送端接收到ack确认后，报文的大小指数级增长, 当cwnd > ssthreshold, 转为拥塞避免，每个RTT后cwnd+1, 当出现超时后, 设置cwnd=1, 重新开始慢启动， 同时ssthreshold降为发送窗口的一半


>cwnd 拥塞窗口
>
>ssthresh(slow start threshold[临界值]) 慢开始阈值 