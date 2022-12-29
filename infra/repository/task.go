package repository

import (
	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
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

func (t *TaskRepository) Create(task *entity.Task) (uuid.UUID, error) {
	return task.Id, t.db.Create(task).Error
}

func (t *TaskRepository) Update(task *entity.Task) (entity.Task, error) {
	return entity.Task{}, nil
}

func (t *TaskRepository) GetById(id string) (entity.Task, error) {
	return entity.Task{}, nil

}

func (t *TaskRepository) GetAll() ([]entity.Task, error) {
	var tasks []entity.Task

	err := t.db.Find(&tasks).Error

	return tasks, err

}

func (t *TaskRepository) GetByUserId(userId string) ([]entity.Task, error) {
	var tasks []entity.Task

	err := t.db.Where("user_id = ?", userId).Order("create_at  asc").Find(&tasks).Error

	return tasks, err
}
