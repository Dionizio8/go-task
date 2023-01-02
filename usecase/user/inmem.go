package user

import (
	"errors"

	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
)

type inmem struct {
	m map[uuid.UUID]*entity.User
}

func newInmem() *inmem {
	var m = map[uuid.UUID]*entity.User{}
	return &inmem{
		m: m,
	}
}

func (r *inmem) Create(user *entity.User) (uuid.UUID, error) {
	r.m[user.Id] = user
	return user.Id, nil
}

func (r *inmem) GetById(id string) (*entity.User, error) {
	uuid := uuid.MustParse(id)
	if r.m[uuid] == nil {
		return &entity.User{}, errors.New("user not found")
	}
	return r.m[uuid], nil
}
