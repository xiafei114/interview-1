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




