package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaMessageExecutor struct {
	writer *kafka.Writer
}

func NewKafkaMessageExecutor(writer *kafka.Writer) *KafkaMessageExecutor {
	return &KafkaMessageExecutor{
		writer: writer,
	}
}

func (k *KafkaMessageExecutor) Push(parent context.Context, key string, value string) (err error) {
	message := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
		Time:  time.Now(),
	}

	return k.writer.WriteMessages(parent, message)
}

func InitKafkaProducer() (w *kafka.Writer) {
	config := LoadKafaConfig()

	return &kafka.Writer{
		Addr:                   kafka.TCP(config.brokers),
		Topic:                  config.topic,
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.LeastBytes{},
	}
}
