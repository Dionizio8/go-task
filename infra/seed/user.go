package seed

import (
	"errors"

	"github.com/Dionizio8/go-task/entity"
	"github.com/Dionizio8/go-task/infra/repository"
	"gorm.io/gorm"
)

type SeedUser struct {
	db *gorm.DB
}

func NewSeedUser(db *gorm.DB) *SeedUser {
	return &SeedUser{
		db: db,
	}
}

var users = []entity.User{
	*entity.NewUser("Manager", entity.GetUserRoleManager()),
	*entity.NewUser("Technician 1", entity.GetUserRoleTechnician()),
	*entity.NewUser("Technician 2", entity.GetUserRoleTechnician()),
}

func (s *SeedUser) Load() {
	if err := s.db.First(&entity.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		repo := repository.NewUserRepository(s.db)

		for _, user := range users {
			repo.Create(&user)
		}
	}
}
