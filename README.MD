# GETTING START WITH KAFKA

## WHAT IS KAFKA
Kafka is not only a simple message broker, kafka is a powerful message broker thar provides fault tolerance and high scalability with your cluster architecture and also provide ETL capatibilites and also mantain the stream data as long as you want
## PRODUCER
First install kafka and zookeeper flowing the documetation steps.
next create a topic and use the command `export KAFKA_HOME=/usr/local/kafka/bin` to configure the location of your kafka instalation
whith kafka runnig use the comand below to start to write messages to kafka topic `TOPIC_NAME=teste-kafka ./scripts/producer.sh`

## CONSUMER
In another tab run the command `TOPIC_NAME=teste-kafka ./scripts/consumer.sh` to receive the messages sent with consumer