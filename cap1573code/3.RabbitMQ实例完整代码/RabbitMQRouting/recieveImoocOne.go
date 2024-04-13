package main

import "rabbit/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_one")
	imoocOne.RecieveRouting()
}
