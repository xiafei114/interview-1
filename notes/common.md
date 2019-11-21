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
     

8. Redis怎么做持久化? 服务主从数据怎么交互
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
    
    
9. RDB原理, AOF和RDB的优缺点

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
   