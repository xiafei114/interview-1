​	

1. consul支持多数据中心,  consul分为Server和Client两种节点, Server保存节点数据，Client负责健康检查以及转发数据请求到Server, Server节点有1个Leader和n个Follower，Leader节点会将数据同步到Follower。
2. 集群内的consul节点通过gossip(流言)协议维护成员关系

> 可以通过consul members 进行查看集群agent, 但是由于节点之间通过gossip协议维护成员关系, 所以得到的成员关系不一定正确，可通过读取server端的数据进行确认
>
> curl 127.0.0.1:8500/v1/catalog/nodes

3. 注册服务

> consul agent -config-dir ~/consul.d
>
> curl 127.0.0.1:8500/v1/catalog/service/web 查看所有的服务
>
> curl http://localhost:8500/v1/catalog/service/web?passing  检查所有健康的服务
>
> curl -X PUT http://localhost:8500/v1/agent/service/deregister/{实例id} 删除无效服务
>
> curl -X PUT  http://localhost:8500/v1/agent/force-leave/{节点id} 删除无效节点

通过给agent发送`SIGHUP` 的信号来进行热更新

4. 双server节点的集群下，一个server推出，讲产生无Leader模式，在三server节点的集群下，leader退出，其余两个会协商在产生一个新的leader
5. 