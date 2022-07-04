package kafka

import (
	"github.com/Shopify/sarama"
)

type KafkaConsumer interface {
	Consume(topic string, c chan *sarama.ConsumerMessage)
}

type KafkaConsumerImpl struct {
	consumer sarama.Consumer
}

func NewKafkaConsumer(url string) KafkaConsumer {
	consumer, err := sarama.NewConsumer([]string{url}, nil)
	if err != nil {
		panic(err)
	}

	return &KafkaConsumerImpl{
		consumer: consumer,
	}
}

func (kafkaConsumer KafkaConsumerImpl) Consume(topic string, chanMessage chan *sarama.ConsumerMessage) {
	partitionList, err := kafkaConsumer.consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}
	for _, partition := range partitionList {
		pc, err := kafkaConsumer.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				chanMessage <- message
			}
		}(pc)
	}
}
