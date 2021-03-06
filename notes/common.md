## redis

1. 如果有大量的key在同一时间内失效, 一般需要在这个时间上加上一个随即值, 避免在同一时间内失效(缓存雪崩)
2. Redis 的分布式缓存, set key value NX EX 10000
3. redis作为消息队列, 一般使用list作为队列, rpush 生产数据, lpop消费数据, list 有一个叫做blpop的指令, 可以阻塞直到有消息到来
4. 查找出redis当中key中以某些固定开头的 eq: keys hello*
5. 如果redis 正在给线上的业务提供服务, 使用keys 会有什么问题? 由于redis是单进程, 使用keys 或者 smembers的指令在查找大批量数据可能会阻塞
服务器较长的时间。 可以考虑选用scan, 增量式迭代命令,  因为在对键进行增量式迭代的过程中， 键可能会被修改， 所以增量式迭代命令只能对被返回的元素提供有限的保证

```
    scan cursor [MATCH match] [COUNT count]
    SCAN 命令是一个基于游标的迭代器（cursor based iterator）： SCAN 命令每次被调用之后， 都会向用户返回一个新的游标， 用户在下次迭代时需要使用这个新游标作为 SCAN 命令的游标参数， 以此来延续之前的迭代过程。
    
    当 SCAN 命令的游标参数被设置为 0 时， 服务器将开始一次新的迭代， 而当服务器向用户返回值为 0 的游标时， 表示迭代已结束

    由于是增量式更新，一直存在在g数据中的键值一定能查询到, 但是新增或者删除的数据无法做保证
     COUNT 选项只是对增量式迭代命令的一种提示
    在迭代一个编码为整数集合（intset，一个只由整数值构成的小集合）、 或者编码为压缩列表
   （ziplist，由不同值构成的一个小哈希或者一个小有序集合）时， 
    增量式迭代命令通常会无视 COUNT 选项指定的值， 在第一次迭代就将数据集包含的所有元素都返回给用户。

    ZSCAN/HSCAN/SSCAN
```
6. 队列的一次生产， 多次消费? pub/sub的订阅者模式, 消费者下线后, 生产者的数据会丢失
7. 如何实现redis的延时队列?
    - 使用有序集合, 消息内容作为key, 时间戳作为score, 使用zadd来生产数据, 消费者用zrangebyscore 来消费数据, 消费完成后
        使用zrem 移除数据
    - 键空间通知(接收键的名称), 配置中修改 notify-keyspace-events: E(开启键空间事件)x(开启过期事件) `PUBLISH __keyevent@0__:del mykey`
    ps: 键事件通知: 接收事件的名称   `PUBLISH __keyspace@0__:mykey del`
    

# 8. Redis怎么做持久化? 服务主从数据怎么交互
RDB/AOF
RDB: 快照, 当符合一定条件时, redis会将内存中的数据以二进制的形式存储在硬盘上, redis重启时, 且未开启AOF, redis会读取RDB持久化生成的二进制文件
触发条件:
    - 客户端执行命令save和bgsave会生成快照
    - 根据配置文件save m n规则进行自动快照 (save second count)
    - 主从复制时，从库全量复制同步主库数据，此时主库会执行bgsave命令进行快照
    - 客户端执行数据库清空命令FLUSHALL时候，触发快照(bgsave)
    - 客户端执行shutdown关闭redis时，触发快照
    (bgsave fork 子进程, fork过程中主进程阻塞)

AOF: redis将每一条写命令以redis通讯协议添加至缓冲区aof_buf,这样的好处在于在大量写请求情况下，
采用缓冲区暂存一部分命令随后根据策略一次性写入磁盘，这样可以减少磁盘的I/O次数，提高性能   

同步命令到硬盘:
    - no: 依赖操作系统write函数去同步,一般30s同步一次, 在大量写操作指令下, aof_buf的空间会大量堆积,如果redis服务崩溃,数据丢失严重
    - always: 每次写操作会调用fsync 强制写入磁盘
    - everysec: 数据将使用调用操作系统write写入文件，并使用fsync每秒一次从内核刷新到磁盘
    
aof 文件重写(bgrewriteaof)   重复或者无效命令不写入文件, 过期命令不写，多条命令合并
为了防止aof重写文件的过程中新的aof文件覆盖旧的aof文件, 导致fork之后的数据丢失, 数据除了进入到aof_buf, 也会同步写入
到aof_rewrite_buf当中，当aof重写完成后, aof_rewrite_buf的文件刷回到新的文件当中

开启aof之后, redis服务启动优先通过aof回复数据
    
# 9. RDB原理, AOF和RDB的优缺点
优点：
**RDB**是一个非常紧凑（compact）的文件，体积小，因此在传输速度上比较快，因此适合灾难恢复。 
RDB 可以最大化 Redis 的性能：父进程在保存 RDB 文件时唯一要做的就是 fork 出一个子进程，然后这个子进程就会处理接下来的所有保存工作，父进程无须执行任何磁盘 I/O 操作。
RDB 在恢复大数据集时的速度比 AOF 的恢复速度要快。
缺点：
RDB是一个快照过程，无法完整的保存所以数据，尤其在数据量比较大时候，一旦出现故障丢失的数据将更多。
当redis中数据集比较大时候，RDB由于RDB方式需要对数据进行完成拷贝并生成快照文件，fork的子进程会耗CPU，并且数据越大，RDB快照生成会越耗时。
RDB文件是特定的格式，阅读性差，由于格式固定，可能存在不兼容情况。


**AOF**
优点：
数据更完整，秒级数据丢失(取决于设置fsync策略)。
兼容性较高，由于是基于redis通讯协议而形成的命令追加方式，无论何种版本的redis都兼容，再者aof文件是明文的，可阅读性较好。
缺点：
数据文件体积较大,即使有重写机制，但是在相同的数据集情况下，AOF文件通常比RDB文件大。
相对RDB方式，AOF速度慢于RDB，并且在数据量大时候，恢复速度AOF速度也是慢于RDB。
由于频繁地将命令同步到文件中，AOF持久化对性能的影响相对RDB较大，但是对于我们来说是可以接受的。

**混合持久化**
优点：
混合持久化结合了RDB持久化 和 AOF 持久化的优点, 由于绝大部分都是RDB格式，加载速度快，同时结合AOF，增量的数据以AOF方式保存了，数据更少的丢失。
缺点：
兼容性差，一旦开启了混合持久化，在4.0之前版本都不识别该aof文件，同时由于前部分是RDB格式，阅读性较差

10. pipeline的好处
11. 同步机制
12. Redis的集群, 如何保证高可用
13. Redis 与 Memcache的对比?
    Redis
    - k, v存储以及丰富的数据结构
    - 内存操作
    - 单线程服务, 才用epoll的网络模型
    - 去中心化的分布式集群
    - 丰富的网络编程接口
    - 支持事务, 发布/订阅, 消息队列等功能
    - 数据持久化 (RDB/AOF)
    

   Memcahe
    - 不提供数据的持久化
    - 数据存储基于LRU
    - 多线程
    - 仅支持k,v 存储
    - 多线程, 非阻塞的网络IO模型
	- key不能超过250字节, value 不能超过1M
	- key的最大失效时间不能超过30天
	- 只支持了k-v结构,不支持持久话和主从同步


14. 缓存雪崩, 击穿, 穿透
雪崩: 同一时间内, 大规模的key同时失效, 大量的请求过来之后, 全部落到DB上, 导致一系列的服务不可用, 可通过过期时间+随机数, 将失效
时间分割到不同的时间段内
穿透: 数据库ID都是从1往上递增的,  如果发起大量的id不存在或者<0的请求, 缓存和数据库都不存在的数据, 短时间内就会导致服务不可用
击穿: 大量的流量集中对某一个key进行访问, 失效的瞬间会导致服务不可用



15. 缓存的双写一致性 
 - 先读缓存, 缓存不存在的时候去读取数据库, 然后取出数据放入到缓存当中, 同时返回响应
 - 更新缓存的时候,先更新数据库, 然后删除缓存
 删除缓存而不是更新缓存是因为, 缓存值往往是需要一系列参数进行计算得到的,仅仅可能因为某些字段的更新,而耗费时间进行新的缓存值的计算是不值得的, 同时,更新缓存的代价时较为高昂的, 某些复杂的计算场景下,缓存所依赖的字段更新较为频繁,但是这个缓存是否是要被频繁计算的? 


16. redis的线程模型 
	IO多路复用, reactor的IO模型
	ps -T -p 端口号 查找该进程下的线程id

	- 纯内存操作 
	- 核心是基于非阻塞的IO多路复用机制
	- 单线程避免了多线程上下文切换的开销


17. redis并发场景下的读写
 - 分布式锁+ 时间戳: 
		- 消息队列, 串行化读写,  比如将通过hash等方式将对某商品Id的读写塞入到同一个队列当中, 之后的写操作通过串行化的顺序进行操作.	

18. redis主从同步
	
- sentinel, 分布式架构, raft算法
	
19. redis过期机制
	- 定时删除
	- 惰性删除

20. 缓存的类型 
	-  本地缓存
	-  分布式缓存
	-  多级缓存
	
	noevicition: 返回异常
	allkeys-lru: 回收最小使用的
	volatile-lru: 回收过期的最小使用键
	allkeys-random: 随机回收
	volatile-random: 随机回收过期键
	volatile-ttl: 回收过期键, 并且优先回收存活时间较短的键


21. redis 数据结构
	- string
	- list
	- set
	- zset
	- hash
	- bitmap
	- HyperLogLog: 不精确的去重计数功能，比较适合用来做大规模数据的去重统计，例如统计 UV
	- Geospatial: 可以用来保存地理位置，并作位置距离计算或者根据半径计算位置等


1. TCP
20字节的固定首部
源端口号(2字节) 目的端口(2字节)
序列号(4字节)报文段第一个字节的数据编号
确认号(4字节)期待收到的下一个报文段的第一个字节的数据编号[当前报文段的最后一个字节的编号+1]
数据偏移,保留,TCP flag(2字节) 窗口(2字节)


URG,ACK,PSH,RST,SYN,FIN(TCP flag)
URG:紧急指针， =1表示某一位需要被紧急处理
ACK:确认号是否有效
PSH:提示接收端应用程序立即从TCP缓冲区把数据读走
RST:对方要求重新建立连接, 复位
SYN:请求建立连接, 并在其序列号的字段进行序列号的初始值设定。建立连接，设置为1
FIN:希望断开连接

MSL: 只一个数据片段在网络内的最大存活时间, 2MSL就是一个发送和回复的最大请求时间


TCP没有IP地址


- 建立连接是SYN超时: 建立连接时SYN超时, 如果server端接收到SYN, 返回SYN+ACK后client掉线了, server端没有接收到client返回的ACK,
server会在一定时间内重发, 默认重发次数是5次，1+2+4+8+16+32 = 63s, 第五次发出后等待32s


tcp_syncookies: 当sync连接队列满了, TCP会通过源端口号, 目标端口号，时间戳, 建立一个特殊的syn, 如果是攻击者则不会有响应，如果是正常连接，则会把这个 SYN Cookie发回来, 请先千万别用tcp_syncookies来处理正常的大负载的连接的情况

/proc/sys/net/ipv4 修改ipv4对应的配置

window(滑动窗口): 接收端告诉发送端自己还有多少	缓冲区可以用来接收数据






网络IO模型(https://blog.csdn.net/li1914309758/article/details/82081899)
1. blocking IO
改进方案
	1. 使用多进程(多线程):	传统意义上, 进程的开销是要大于线程, 所以如果要为较多的客户提供服务, 使用线程, 
	如果单个服务执行需要消耗较多的CPU资源，则使用进程


2. no-blocking IO
	非阻塞IO中, 用户进程是在不断轮询kernal的数据准备好没有(不推荐，循环调用recv()接口将大幅度推高CPU的占用率)

3. IO multiplexing(事件驱动IO)
单个process处理多个socket连接，用select以后最大的优势是用户可以在一个线程内同时处理多个socket的IO请求。用户可以注册多个socket，然后不断地调用select读取被激活的socket，即可达到在同一个线程内同时处理多个IO请求的目的。而在同步阻塞模型中，必须通过多线程的方式才能达到这个目的。（多说一句：所以，如果处理的连接数不是很高的话，使用select/epoll的web server不一定比使用multi-threading + blocking IO的web server性能更好，可能延迟还更大。select/epoll的优势并不是对于单个连接能处理得更快，而是在于能处理更多的连接。）
缺点:
	1. 首先select()接口并不是实现“事件驱动”的最好选择。因为当需要探测的句柄值较大时，select()接口本身需要消耗大量时间去轮询各个句柄
	2. 模型将事件探测和事件响应夹杂在一起，一旦事件响应的执行体庞大，则对整个模型是灾难性的

4. signal driven IO



5. asynchronous IO(异步IO)


差异: 异步与阻塞的差异, 阻塞IO最终会在recvfrom函数处阻塞，等待数据从内核空间copy到用户空间, 异步io不会在这一步阻塞，等到所有事情处理完成之后，内核才会通知socket数据处理完成



CDN回源策略

回源域名

浏览器缓存->边缘层->区域层->中心层

CDN节点数据更新不及时







prometheus 学习笔记

pagerduty

1. 监控数据的精细程度


采集频率过高后会导致，数据量过大


dau: 日活

优质特性

1. 基于时间序列模型, 通常是时间间隔的采样数据
2. 基于K/V的数据模型
3. 完全基于数学运算
4. 采用 pull/push 两种对应的数据采集传输方式
5. 开源，大量的社区成品插件
6. 本身自带图形调试
7. 精细的数据采样


缺点
1. 不支持集群化
2. 监控集群过大，性能瓶颈
3. 数学要求高
4. 监控的经验要有很高的要求, 很多监控项需要很细的定制

guages
counters
histogram 统计数据的分布情况(近似百分比估算数值)



Cpu状态 八种状态
cpu使用率 = (所有非空闲状态的Cpu使用时间总和) / 所有状态的CPU使用时间的总和



K/V metrics

gauag 瞬时变化指标




increase() 针对Counter, 截取一段时间内的增量



2. 5种网络I/O模型，阻塞、非阻塞、I/O多路复用、信号驱动IO、异步I/O。从数据从I/O设备到内核态，内核态到进程用户态分别描述这5种的区别
3. 从select的机制，以及select的三个缺点，讲解epoll机制，以及epoll是如何解决select的三个缺点的。还会讲到epoll中水平触发和边沿触发的区别。
4. leetcode 3道数组题





1. 四次挥手中TIME_WAIT状态存在的目的是什么: 用来重发可能丢失的ACK报文
2. 要求熟悉三次握手和四次挥手的机制，要求画出状态图 (https://blog.csdn.net/qq_38950316/article/details/81087809)


   netstat 
   -a: 列出所有的连接
   -n: 不解析名称(禁用域名解析功能)
   -p: 显示程序名称
   -t: 显示tcp/udp 
   -l: 监听中的连接
   -s: 打印统计数据
   -c: 持续输出
   -e: 扩展信息, 用户等 



** 全连接队列/半连接队列 ** 
半连接队列: client 向 server发出sync请求成功后,该链接被放入到半连接队列, 同时向client回复ACK+SYN
全连接队列: 
tcp_abort_on_overflow 0: 全连接队列满了之后, server会扔掉client发来的ack. 1: 全连接满了之后,server发送一个reset包给client, 表示废掉这个过程和这个连接



