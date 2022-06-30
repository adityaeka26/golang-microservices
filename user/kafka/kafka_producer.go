package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

type KafkaProducer interface {
	SendMessage(topic string, msg string) error
}

type KafkaProducerImpl struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer() KafkaProducer {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	producers, err := sarama.NewSyncProducer([]string{"localhost:9092"}, kafkaConfig)
	if err != nil {
		panic(err)
	}

	return &KafkaProducerImpl{
		producer: producers,
	}
}

func (kafkaProducer KafkaProducerImpl) SendMessage(topic string, msg string) error {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	_, _, err := kafkaProducer.producer.SendMessage(kafkaMsg)
	if err != nil {
		return err
	}

	return nil
}
