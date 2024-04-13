package main

import (
	"fmt"
	"rabbit/RabbitMQ"
)

func main() {
	//rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
	//	"imoocSimple")
	//rabbitmq.PublishSimple("Hello imooc!")
	//fmt.Println("发送成功！")

	rabbitmq := RabbitMQ.NewRabbitMQSimple("test_client1")
	rabbitmq.PublishSimple("Hello Tim!")
	fmt.Println("发送成功！")
}
