package repository

import (
	"errors"

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

func (r *TaskRepository) Create(task *entity.Task) (uuid.UUID, error) {
	return task.Id, r.db.Create(task).Error
}

func (r *TaskRepository) UpdateStatus(taskId, userId, status string) (entity.Task, error) {
	var task entity.Task
	data := r.db.First(&task, "id = ?", taskId).Where("user_id = ?", userId)
	if data.Error != nil {
		if errors.Is(data.Error, gorm.ErrRecordNotFound) {
			return entity.Task{}, errors.New("task not found")
		}
		return entity.Task{}, data.Error
	}

	data.Update("status", status)

	return task, nil
}

func (r *TaskRepository) GetById(taskId, userId string) (entity.Task, error) {
	var task entity.Task
	err := r.db.First(&task, "id = ?", taskId).Where("user_id = ?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Task{}, errors.New("task not found")
		}
		return entity.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) GetAll() ([]entity.Task, error) {
	var tasks []entity.Task
	err := r.db.Find(&tasks).Error

	return tasks, err

}

func (r *TaskRepository) GetByUserId(userId string) ([]entity.Task, error) {
	var tasks []entity.Task
	err := r.db.Where("user_id = ?", userId).Order("create_at  asc").Find(&tasks).Error

	return tasks, err
}
