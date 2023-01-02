package task

import (
	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
)

type Repository interface {
	Create(task *entity.Task) (uuid.UUID, error)
	UpdateStatus(taskId, userId, status string) (*entity.Task, error)
	GetById(taskId, userId string) (*entity.Task, error)
	GetAll() ([]*entity.Task, error)
	GetByUserId(userId string) ([]*entity.Task, error)
}

type UseCase interface {
	CreateTask(title, description string, userId, managerId uuid.UUID, status string) (uuid.UUID, error)
	EditTaskStatus(taskId, userId, status string) (*entity.Task, error)
	FindTasktById(taskId, userId string) (*entity.Task, error)
	List() ([]*entity.Task, error)
	FindTaskByUserId(userId string) ([]*entity.Task, error)
}
