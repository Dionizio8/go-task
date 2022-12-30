package worker

import (
	"context"
	"fmt"

	"github.com/Dionizio8/go-task/infra/kafka"
)

type NotificationHandler struct {
	msg kafka.KafkaMessageConsumer
}

func NewNotificationHandler(msg kafka.KafkaMessageConsumer) *NotificationHandler {
	return &NotificationHandler{
		msg: msg,
	}
}

func (n *NotificationHandler) Execute(ctx context.Context) {
	for {
		msg, err := n.msg.Pull(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		//TODO: “The tech X performed the task Y on date
		fmt.Println("received:", string(msg.Value))
	}
}
