package user

import (
	"testing"
	"time"

	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newFixtureUser() *entity.User {
	return &entity.User{
		Id:       uuid.New(),
		Name:     "Gabriel",
		Role:     entity.GetUserRoleTechnician(),
		CreateAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureUser()
	_, err := m.CreateUser(u.Name, u.Role)
	assert.Nil(t, err)
	assert.False(t, u.CreateAt.IsZero())
	assert.True(t, u.UpdateAt.IsZero())
}

func Test_GetById(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2.Name = "Dionizio"

	uud1, _ := m.CreateUser(u1.Name, u1.Role)
	uud2, _ := m.CreateUser(u2.Name, u2.Role)

	r1, err := m.FindUserById(uud1.String())
	assert.Nil(t, err)
	assert.Equal(t, "Gabriel", r1.Name)

	r2, err := m.FindUserById(uud2.String())
	assert.Nil(t, err)
	assert.Equal(t, "Dionizio", r2.Name)
}
