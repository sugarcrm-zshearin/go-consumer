package kafka

import (
	"github.com/Shopify/sarama"
)

func GeSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	//producer, err := sarama.NewSyncProducer(brokers, config)
	//return producer, err
	return sarama.NewSyncProducer(brokers, config)

}

func GetMessage(message string, topicName string, partition int32) *sarama.ProducerMessage {
	payload := message
	return &sarama.ProducerMessage{
		Topic:     topicName,
		Partition: partition,
		Value:     sarama.ByteEncoder(payload),
	}

}
