package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	task := NewTask("Title Task", "Desc Task", uuid.New(), uuid.New(), GetTaskInitialState())
	assert.Equal(t, task.Title, "Title Task")
	assert.Equal(t, task.Description, "Desc Task")
	assert.Equal(t, task.Status, GetTaskInitialState())
	assert.NotNil(t, task.Id)
	assert.NotNil(t, task.UserId)
	assert.NotNil(t, task.ManagerUserId)
}

func TestGetTaskState(t *testing.T) {
	assert.Equal(t, GetTaskInitialState(), "IN_PROGRESS")
	assert.Equal(t, GetTaskFinalState(), "CONCLUDED")
}
