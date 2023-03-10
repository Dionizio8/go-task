package user

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

func (s *Service) CreateUser(name, role string) (uuid.UUID, error) {
	user := entity.NewUser(name, role)
	return s.repo.Create(user)
}

func (s *Service) FindUserById(id string) (*entity.User, error) {
	return s.repo.GetById(id)
}
