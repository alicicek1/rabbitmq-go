package main

import (
	"consumer/config"
	"consumer/pkg/rabbitmq"
	"log"
	"sync"
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
	msgS := rabbitmq.ConsumeMessage(channel, queue)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for d := range msgS {
			log.Printf("Received a message: %s", d.Body)
		}
		wg.Wait()
	}()

	wg.Done()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
