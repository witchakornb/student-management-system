package student_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/witchakornb/student-management-system/Entity/User"
	"github.com/witchakornb/student-management-system/Entity/Student"
	"github.com/witchakornb/student-management-system/Database"
)

func TestStudentCRUD(t *testing.T) {
	// Connecting to database without modification
	cfg := database.ConfigDatabaseWithPath("../../.env")
	db, err := database.NewMySQL(cfg)
	assert.NoError(t, err, "failed to connect to database")
	defer func() {
		database.Close(db)
	}()

	t.Run("Insert Student", func(t *testing.T) {
		assert := assert.New(t)

		// Create test user
		newUser := user.User{
			UserID:      "12345",
			Username:    "testuser",
			Password:    "password",
			Role:        "student",
			FirstNameTH: "Test",
			LastNameTH:  "User",
			FirstNameEN: "Test",
			LastNameEN:  "User",
			Email:       "test@example.com",
			Phone:       "1234567890",
			Department:  "Computer Science",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		assert.NoError(db.Create(&newUser).Error, "failed to create user")

		// Create student linked to user
		newStudent := student.Student{
			StudentID:        newUser.UserID,
			StudentCreatedAt: time.Now(),
			StudentUpdatedAt: time.Now(),
			User:             newUser,
		}
		assert.NoError(db.Create(&newStudent).Error, "failed to create student")
	})

	t.Run("Retrieve Student", func(t *testing.T) {
		assert := assert.New(t)

		// Retrieve student with user data
		retrievedStudent := student.Student{}
		assert.NoError(db.Preload("User").First(&retrievedStudent, "student_id = ?", "12345").Error, "failed to retrieve student")
		assert.Equal("testuser", retrievedStudent.User.Username, "user data mismatch")
	})

	t.Run("Update Student", func(t *testing.T) {
		assert := assert.New(t)

		// Update student
		updatedTime := time.Now().Add(24 * time.Hour)
		assert.NoError(db.Model(&student.Student{}).Where("student_id = ?", "12345").Update("student_updated_at", updatedTime).Error, "failed to update student")

		// Update user
		updatedEmail := "updated@example.com"
		assert.NoError(db.Model(&user.User{}).Where("user_id = ?", "12345").Update("email", updatedEmail).Error, "failed to update user")

		// Verify updates
		retrievedStudent := student.Student{}
		assert.NoError(db.Preload("User").First(&retrievedStudent, "student_id = ?", "12345").Error, "failed to retrieve updated student")
		assert.Equal(updatedTime.Unix(), retrievedStudent.StudentUpdatedAt.Unix(), "student updated_at mismatch")
		assert.Equal(updatedEmail, retrievedStudent.User.Email, "user email mismatch")
	})

	t.Run("Delete Student", func(t *testing.T) {
		assert := assert.New(t)

		// Delete student
		assert.NoError(db.Delete(&student.Student{}, "student_id = ?", "12345").Error, "failed to delete student")

		// Delete user
		assert.NoError(db.Delete(&user.User{}, "user_id = ?", "12345").Error, "failed to delete user")

		// Verify deletion
		retrievedStudent := student.Student{}
		assert.Error(db.First(&retrievedStudent, "student_id = ?", "12345").Error, "expected error when retrieving deleted student")
		retrievedUser := user.User{}
		assert.Error(db.First(&retrievedUser, "user_id = ?", "12345").Error, "expected error when retrieving deleted user")
	})
}
