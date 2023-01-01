package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := NewUser("Dionizio", GetUserRoleTechnician())
	assert.Equal(t, user.Name, "Dionizio")
	assert.Equal(t, user.Role, GetUserRoleTechnician())
	assert.NotNil(t, user.Id)
}

func TestUserRole(t *testing.T) {
	assert.Equal(t, GetUserRoleManager(), "MANAGER")
	assert.Equal(t, GetUserRoleTechnician(), "TECHNICIAN")
}
