1. Kafka为什么要使用分区的概念而不是直接使用主题

   提高并发能力, 同一个消费组下的消费者读取不同分区下的数据并发消费。不同的分区分布在不同的broker上, 每一台机器都可以独立的进行读写, 增强集群的高可用

2. 分区策略
   	- 轮训分区
   	- 随机分区
   	- Key-Order分区(自定义分区): 按照业务属性进行进去, 当业务需要保证一定的顺序处理时, 同一个业务属性下的数据进入到同一个分区 
   	- 自定义分区

3. 数据压缩

   producer 压缩, broker 保持, comsumer 消费

   	- 当producer和broker设置的压缩算法不一致时, broker会对数据进行重新压缩
   	- 当同一个kafka使用了不同版本的消息时, 为了兼容老版本, broker会将数据从新版本转化成老版本
   	- broker解压缩会进行数据的校验

![img](https://static001.geekbang.org/resource/image/cf/68/cfe20a2cdcb1ae3b304777f7be928068.png)



**数据压缩的最佳实践**

- 数据压缩非常吃CPU资源, 当CPU资源紧缺时，尽量少开启压缩
- CPU充足时, 考虑打开压缩, 减少带宽的损耗
- 尽量避免kafka消息协议的版本不一致, 减少broker端对数据的解压缩以及协议的转化



**位移主题__consumer_offsets**

对于老版本的kafka来说， consumer的位移信息被保存在了zookeeper上, kafka作为一个无状态的服务,能够最大程度上的支持kafka的扩展, 后来在使用的过程中发现对于主题存在着大量的写, zookeeper不适合这样的工作, 新版本将位移的信息保存到了broker上面, 也就是__consumer_offsets 这个主题当中. 



位移的提交方式:

​	位移的提交方式可以根据配置`enable.auto.commit` 来控制

 - 自动提交位移:  当`enable.auto.commit=true`时，提交间隔根据`auto.commit.intervals.ms` 控制, 举一个极端的例子, 某些场景下, 没有consumer消费数据了， 仍然设置自动提交的话，位移永远保持在某一个位置上，consumer_offset的主题中会存在大量相同位移点的数据, kafka会采用compact的策略来整合相同的数据(<consumer_group_id, topic, partitionId>：value).

   >  kafka的后台提供了Log Cleaner的线程针对compac的主图来定时清理满足条件的可删除数据

	- 主动提交位移



**重平衡Rebalance**

​	Rebalance本质上是一种协议, 规定了一个consumer group下的consumer如何达成一致，来分配订阅Topic的每个分区。Rebalance的过程中会发生STW(所有consumer实例不能消费消息)

   Rebalance发生的条件:

  - 组内成员数发生改变:  consumer group中consumer数减少的场景中，有两类是不必要的
    1. 由于心跳不及时，导致被coordinator踢出组, consumer当中有`session.timeout.ms`配置，该配置下coordinator认为在n秒内没有收到consumer的心跳，就会认为该consumer挂了，从而提出组，consumer还有个 `heartbeat.interval.ms`表示心跳的发送频率。
    2.  `max.poll.interval.ms`表示如果一定时间内consumer没有消费完poll返回的消息, consumer会主动发起离开组的请求，导致Rebalance。
		-  订阅的主题数发生改变
		- 订阅主题的分区数发生改变



Coordinator专门为consumer Group 服务，负责为 Group 执行 Rebalance 以及提供位移管理和组成员管理等。

