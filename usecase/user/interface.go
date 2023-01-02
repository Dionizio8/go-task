package user

import (
	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
)

type Repository interface {
	Create(user *entity.User) (uuid.UUID, error)
	GetById(id string) (*entity.User, error)
}

type UseCase interface {
	CreateUser(user entity.User) (uuid.UUID, error)
	FindUserById(id string) (*entity.User, error)
}
