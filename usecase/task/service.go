package task

import (
	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateTask(title, description string, userId, managerId uuid.UUID, status string) (uuid.UUID, error) {
	task := entity.NewTask(title, description, userId, managerId, status)
	return s.repo.Create(task)
}

func (s *Service) EditTask(*entity.Task) (entity.Task, error) {
	return entity.Task{}, nil
}

func (s *Service) FindTasktById(id string) (entity.Task, error) {
	return entity.Task{}, nil
}

func (s *Service) List() ([]entity.Task, error) {
	return s.repo.GetAll()
}

func (s *Service) FindTaskByUserId(userId string) ([]entity.Task, error) {
	return s.repo.GetByUserId(userId)
}
