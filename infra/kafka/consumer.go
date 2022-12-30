package kafka

import (
	"context"
	"strings"

	"github.com/segmentio/kafka-go"
)

type KafkaMessageConsumer struct {
	reader *kafka.Reader
}

func NewKafkaMessageConsumer(reader *kafka.Reader) *KafkaMessageConsumer {
	return &KafkaMessageConsumer{
		reader: reader,
	}
}

func (k *KafkaMessageConsumer) Pull(parent context.Context) (kafka.Message, error) {
	return k.reader.ReadMessage(parent)
}

func InitKafkaConsumer() (w *kafka.Reader) {
	config := LoadKafaConfig()

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(config.brokers, ","),
		Topic:   config.topic,
		GroupID: config.clientId,
	})
}
