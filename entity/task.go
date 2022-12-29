package entity

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id            uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	UserId        uuid.UUID `json:"user_id" gorm:"type:char(36);"`
	ManagerUserId uuid.UUID `json:"manager_user_id" gorm:"type:char(36);"`
	Status        string    `json:"status" gorm:"index"`
	CreateAt      time.Time `json:"create_at" gorm:"index"`
	UpdateAt      time.Time `json:"update_at"`
}

func NewTask(title, description string, userId uuid.UUID, managerUserId uuid.UUID, status string) *Task {
	return &Task{
		Id:            uuid.New(),
		Title:         title,
		Description:   description,
		UserId:        userId,
		ManagerUserId: managerUserId,
		Status:        status,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
	}
}
