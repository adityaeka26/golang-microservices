package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/adityaeka26/golang-microservices/notification/logger"
	"go.uber.org/zap"
)

type KafkaConsumer interface {
	Consume(topic string, c chan sarama.ConsumerMessage)
}

type KafkaConsumerImpl struct {
	consumer sarama.Consumer
	logger   logger.Logger
}

func NewKafkaConsumer(url string, logger logger.Logger) KafkaConsumer {
	consumer, err := sarama.NewConsumer([]string{url}, nil)
	if err != nil {
		panic(err)
	}

	return &KafkaConsumerImpl{
		consumer: consumer,
		logger:   logger,
	}
}

func (kafkaConsumer KafkaConsumerImpl) Consume(topic string, chanMessage chan sarama.ConsumerMessage) {
	context := "kafkaConsumer-Consume"

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
				chanMessage <- *message

				kafkaConsumer.logger.GetLogger().Info(
					"Consume message success",
					zap.String("context", context),
					zap.String("topic", topic),
					zap.ByteString("data", message.Value),
				)
			}
		}(pc)
	}
}
