package task

import (
	"testing"
	"time"

	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newFixtureTask() *entity.Task {
	return &entity.Task{
		Id:            uuid.New(),
		Title:         "Title task",
		Description:   "Desc task",
		UserId:        uuid.New(),
		ManagerUserId: uuid.New(),
		Status:        entity.GetTaskInitialState(),
		CreateAt:      time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureTask()

	taskId, err := s.CreateTask(u.Title, u.Description, u.UserId, u.ManagerUserId, u.Status)
	assert.Nil(t, err)
	assert.NotNil(t, taskId)
	assert.False(t, u.CreateAt.IsZero())
	assert.True(t, u.UpdateAt.IsZero())
}

func Test_UpdateStatus(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	u := newFixtureTask()
	taskId, _ := s.CreateTask(u.Title, u.Description, u.UserId, u.ManagerUserId, u.Status)

	task, err := s.EditTaskStatus(taskId.String(), u.UserId.String(), entity.GetTaskFinalState())
	assert.Nil(t, err)
	assert.Equal(t, task.Status, entity.GetTaskFinalState())
}

func Test_GetById(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	t1 := newFixtureTask()
	t2 := newFixtureTask()
	t2.Title = "Title task 2"
	t2.UserId = t1.UserId

	taskId1, _ := s.CreateTask(t1.Title, t1.Description, t1.UserId, t1.ManagerUserId, t1.Status)
	taskId2, _ := s.CreateTask(t2.Title, t2.Description, t2.UserId, t2.ManagerUserId, t2.Status)

	r1, err := s.FindTasktById(taskId1.String(), t1.UserId.String())
	assert.Nil(t, err)
	assert.Equal(t, "Title task", r1.Title)

	r2, err := s.FindTasktById(taskId2.String(), t2.UserId.String())
	assert.Nil(t, err)
	assert.Equal(t, "Title task 2", r2.Title)
}

func Test_GetAll(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	t1 := newFixtureTask()
	t2 := newFixtureTask()
	t2.Title = "Title task 2"

	s.CreateTask(t1.Title, t1.Description, t1.UserId, t1.ManagerUserId, t1.Status)
	s.CreateTask(t2.Title, t2.Description, t2.UserId, t2.ManagerUserId, t2.Status)

	tasks, err := s.List()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(tasks))
	assert.Equal(t, "Title task", tasks[0].Title)
}

func Test_GetByUserId(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	t1 := newFixtureTask()
	t2 := newFixtureTask()
	t2.UserId = t1.UserId
	t3 := newFixtureTask()

	s.CreateTask(t1.Title, t1.Description, t1.UserId, t1.ManagerUserId, t1.Status)
	s.CreateTask(t2.Title, t2.Description, t2.UserId, t2.ManagerUserId, t2.Status)
	s.CreateTask(t3.Title, t3.Description, t3.UserId, t3.ManagerUserId, t3.Status)

	tasks, err := s.FindTaskByUserId(t1.UserId.String())
	assert.Nil(t, err)
	assert.Equal(t, 2, len(tasks))
	assert.Equal(t, "Title task", tasks[0].Title)
}
