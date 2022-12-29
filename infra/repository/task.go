package repository

import (
	"github.com/Dionizio8/go-task/entity"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) Create(task *entity.Task) error {
	return t.db.Create(task).Error
}

func (t *TaskRepository) Update(task *entity.Task) (entity.Task, error) {
	return entity.Task{}, nil
}

func (t *TaskRepository) GetById(id string) (entity.Task, error) {
	return entity.Task{}, nil

}

func (t *TaskRepository) GetAll() ([]entity.Task, error) {
	return []entity.Task{}, nil

}

func (t *TaskRepository) GetByUserId(userId string) ([]entity.Task, error) {
	return []entity.Task{}, nil
}
