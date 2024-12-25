package user_test

import (
	"github.com/witchakornb/student-management-system/Entity/User"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserStruct(t *testing.T) {
	now := time.Now()
	deletedAt := gorm.DeletedAt{Time: now, Valid: true}

	u := user.User{
		UserID:      "123",
		Username:    "testuser",
		Password:    "password",
		Role:        "student",
		FirstNameTH: "Test",
		LastNameTH:  "User",
		FirstNameEN: "Test",
		LastNameEN:  "User",
		Email:       "testuser@example.com",
		Phone:       "1234567890",
		Department:  "Computer Science",
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   deletedAt,
	}

	assert.Equal(t, "123", u.UserID)
	assert.Equal(t, "testuser", u.Username)
	assert.Equal(t, "password", u.Password)
	assert.Equal(t, "student", u.Role)
	assert.Equal(t, "Test", u.FirstNameTH)
	assert.Equal(t, "User", u.LastNameTH)
	assert.Equal(t, "Test", u.FirstNameEN)
	assert.Equal(t, "User", u.LastNameEN)
	assert.Equal(t, "testuser@example.com", u.Email)
	assert.Equal(t, "1234567890", u.Phone)
	assert.Equal(t, "Computer Science", u.Department)
	assert.Equal(t, now, u.CreatedAt)
	assert.Equal(t, now, u.UpdatedAt)
	assert.Equal(t, deletedAt, u.DeletedAt)
}