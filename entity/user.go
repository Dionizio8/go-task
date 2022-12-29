package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	roleManager    = "MANAGER"
	roleTechnician = "TECHNICIAN"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name     string    `json:"name"`
	Role     string    `json:"role" gorm:"index"`
	CreateAt time.Time `json:"create_at" gorm:"index"`
	UpdateAt time.Time `json:"update_at"`
}

func NewUser(userName, role string) *User {
	return &User{
		Id:       uuid.New(),
		Name:     userName,
		Role:     role,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

func GetUserRoleManager() string {
	return roleManager
}

func GetUserRoleTechnician() string {
	return roleTechnician
}
