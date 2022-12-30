package kafka

import "time"

type TaskMessage struct {
	UserId    string
	ManagerId string
	TaskId    string
	Status    string
	Date      time.Time
}
