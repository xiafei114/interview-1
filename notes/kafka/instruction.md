- 查看所有的topic列表:  `./bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --list `

- 查看topic的详细信息: `./bin/kafka-topics.sh --zookeeper 127.0.0.1:2181 --topic [topic] --describe`

- 启动zookeeper: `./bin/zookeeper-server-start.sh config/zookeeper.properties`

- 启动kafka `./bin/kafka-server-start.sh  config/kafka.properties`

- 启动生产者 `./bin/kafka-console-producer.sh --broker-list 127.0.0.1:2181 --topic [topic]`

- 启动消费者 `./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic cctv1 --from-beginning`

  

  

  