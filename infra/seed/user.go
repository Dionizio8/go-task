package seed

import (
	"github.com/Dionizio8/go-task/entity"
	"gorm.io/gorm"
)

type Seed struct {
}

var users = []entity.User{
	*entity.NewUser("Gabriel", "MANAGER"),
	*entity.NewUser("Dionizio", "TECHNICIAN"),
}

func (s *Seed) Load(db *gorm.DB) {}
