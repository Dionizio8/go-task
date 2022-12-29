package task

import (
	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
)

type Repository interface {
	Create(task *entity.Task) error
	Update(task *entity.Task) (entity.Task, error)
	GetById(id string) (entity.Task, error)
	GetAll() ([]entity.Task, error)
	GetByUserId(userId string) ([]entity.Task, error)
}

type UseCase interface {
	CreateTask(title, description string, userId, managerId uuid.UUID, status string) error
	EditTask(*entity.Task) (entity.Task, error)
	FindTasktById(id string) (entity.Task, error)
	List() ([]entity.Task, error)
	FindTaskByUserId(userId string) ([]entity.Task, error)
}
