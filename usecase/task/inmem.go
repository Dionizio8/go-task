package task

import (
	"errors"

	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
)

type inmem struct {
	m map[uuid.UUID]*entity.Task
}

func newInmem() *inmem {
	var m = map[uuid.UUID]*entity.Task{}
	return &inmem{
		m: m,
	}
}

func (r *inmem) Create(task *entity.Task) (uuid.UUID, error) {
	r.m[task.Id] = task
	return task.Id, nil
}

func (r *inmem) GetById(taskId, userId string) (*entity.Task, error) {
	uuid := uuid.MustParse(taskId)

	if r.m[uuid] == nil {
		return &entity.Task{}, errors.New("user not found")
	}
	return r.m[uuid], nil
}

func (r *inmem) UpdateStatus(taskId, userId, status string) (*entity.Task, error) {
	task, err := r.GetById(taskId, userId)
	if err != nil {
		return &entity.Task{}, err
	}

	task.Status = status

	r.m[task.Id] = task
	return task, nil
}

func (r *inmem) GetAll() ([]*entity.Task, error) {
	var t []*entity.Task
	for _, j := range r.m {
		t = append(t, j)
	}
	return t, nil
}

func (r *inmem) GetByUserId(userId string) ([]*entity.Task, error) {
	var t []*entity.Task
	uuid := uuid.MustParse(userId)
	for _, j := range r.m {
		if j.UserId == uuid {
			t = append(t, j)
		}
	}
	return t, nil
}
