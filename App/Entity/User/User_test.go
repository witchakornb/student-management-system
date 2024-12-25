package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user := User{
		UserID:      "123",
		Username:    "testuser",
		Password:    "password",
		Role:        "admin",
		FirstNameTH: "ชื่อ",
		LastNameTH:  "นามสกุล",
		FirstNameEN: "FirstName",
		LastNameEN:  "LastName",
		Email:       "testuser@example.com",
		Phone:       "1234567890",
		Department:  "IT",
	}

	assert.Equal(t, "123", user.UserID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "password", user.Password)
	assert.Equal(t, "admin", user.Role)
	assert.Equal(t, "ชื่อ", user.FirstNameTH)
	assert.Equal(t, "นามสกุล", user.LastNameTH)
	assert.Equal(t, "FirstName", user.FirstNameEN)
	assert.Equal(t, "LastName", user.LastNameEN)
	assert.Equal(t, "testuser@example.com", user.Email)
	assert.Equal(t, "1234567890", user.Phone)
	assert.Equal(t, "IT", user.Department)
}