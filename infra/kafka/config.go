package kafka

import (
	"os"
)

type KafaConfig struct {
	topic    string
	brokers  string
	clientId string
}

func LoadKafaConfig() *KafaConfig {
	return &KafaConfig{
		topic:    "go-task-notification-topic",
		brokers:  os.Getenv("KAFKA_BROKERS"),
		clientId: "notification",
	}
}
