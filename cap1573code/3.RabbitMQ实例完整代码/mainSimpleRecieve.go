package main

import "rabbit/RabbitMQ"

func main() {
	//rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
	//	"imoocSimple")
	//rabbitmq.ConsumeSimple()

	rabbitmq := RabbitMQ.NewRabbitMQSimple("test_client1")
	rabbitmq.ConsumeSimple()
}
