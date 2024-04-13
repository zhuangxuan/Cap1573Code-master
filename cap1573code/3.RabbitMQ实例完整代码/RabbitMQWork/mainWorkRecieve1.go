package main

import "rabbit/RabbitMQ"

func main() {

	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rabbitmq.ConsumeSimple()
}
