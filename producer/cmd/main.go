package main

import (
	"fmt"
	"producer/config"
	"producer/pkg/rabbitmq"
	"producer/type/error"
	"time"
)

func main() {
	cfg := config.GetConfig()
	conStr := "amqp://" + cfg.RabbitUsername + ":" + cfg.RabbitPassword + "@" + cfg.RabbitUri + "/"
	conn := rabbitmq.Connect(conStr)
	defer conn.Close()
	channel := rabbitmq.OpenChannel(conn)
	defer channel.Close()
	qName := cfg.ErrorQName
	queue := rabbitmq.DeclareAQueue(channel, qName)

	val := rabbitmq.PublishToQueue(channel, queue, error.Error{
		AppName:     "Producer",
		Operation:   "Publish",
		Code:        914,
		CreatedDate: time.Now(),
	})

	if val {
		fmt.Println("Message successfully published.")
	}
}
