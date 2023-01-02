package repository

import (
	"errors"

	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *entity.User) (uuid.UUID, error) {
	return user.Id, r.db.Create(user).Error
}

func (r *UserRepository) GetById(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &entity.User{}, errors.New("user not found")
		}
		return &entity.User{}, err
	}

	return &user, nil
}
