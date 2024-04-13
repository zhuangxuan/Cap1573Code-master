package main

import "rabbit/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_two")
	imoocOne.RecieveRouting()
}
