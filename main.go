package main

import (
	"log"

	"github.com/sugarcrm-zshearin/go-consumer/kafka"
)

// Sarama configuration options
var (
	broker    = "192.168.99.112:30203" //This is minikube IP and kafka exposed port
	topicName = "my-topic"
	groupName = "my-group"
)

func main() {

	brokers := []string{broker}
	producer, err := kafka.GeSyncProducer(brokers)
	if err != nil {
		log.Fatalf("error creating producer", err)
	}

	message := kafka.GetMessage("I sent a message", topicName, 1)

	producer.SendMessage(message)

	kafka.CreateAndRunConsumer(broker, topicName, groupName, false)

}
