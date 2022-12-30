package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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

		var data kafka.TaskMessage

		err = json.Unmarshal(msg.Value, &data)
		if err != nil {
			log.Println(err.Error())
		}

		fmt.Printf("The tech %v performed the task %v on %v\n", data.UserId, data.TaskId, data.Date.Format("2006-01-02 15:04:05"))
	}
}
