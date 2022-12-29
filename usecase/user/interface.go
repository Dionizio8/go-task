package user

import (
	"github.com/Dionizio8/go-task/entity"
)

type Repository interface {
	Create(user *entity.User) error
	GetById(id string) (entity.User, error)
}

type UseCase interface {
	CreateUser(user entity.User) error
	FindUserById(id string) (entity.User, error)
}
