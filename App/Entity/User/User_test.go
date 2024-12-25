package user_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	userpkg "github.com/witchakornb/student-management-system/Entity/User"
	"github.com/witchakornb/student-management-system/Database"
)


// TestUserCreation function
func TestUserCreation(t *testing.T) {
	now := time.Now()
	user := userpkg.User{
		UserID:      "12345",
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
		DeletedAt:   gorm.DeletedAt{},
	}

	assert.Equal(t, "12345", user.UserID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "password", user.Password)
	assert.Equal(t, "student", user.Role)
	assert.Equal(t, "Test", user.FirstNameTH)
	assert.Equal(t, "User", user.LastNameTH)
	assert.Equal(t, "Test", user.FirstNameEN)
	assert.Equal(t, "User", user.LastNameEN)
	assert.Equal(t, "testuser@example.com", user.Email)
	assert.Equal(t, "1234567890", user.Phone)
	assert.Equal(t, "Computer Science", user.Department)
	assert.Equal(t, now, user.CreatedAt)
	assert.Equal(t, now, user.UpdatedAt)
	assert.Equal(t, gorm.DeletedAt{}, user.DeletedAt)
}

// TestUserUniqueConstraints function
func TestUserUniqueConstraints(t *testing.T) {
	user1 := userpkg.User{
		UserID:   "12345",
		Username: "testuser",
		Password: "password",
		Email:    "testuser@example.com",
	}

	user2 := userpkg.User{
		UserID:   "12346",
		Username: "testuser", // Same username as user1
		Password: "password",
		Email:    "testuser2@example.com",
	}

	assert.NotEqual(t, user1.UserID, user2.UserID)
	assert.Equal(t, user1.Username, user2.Username)
	assert.NotEqual(t, user1.Email, user2.Email)
}

// Test Insert User to Database
func TestInsertUser(t *testing.T) {
	t.Run("Insert User", func(t *testing.T) {
		assert := assert.New(t)
		now := time.Now()
		user := userpkg.User{
			UserID:      "12345",
			Username:    "testuser",
			Password:    "password",
			Role:        "student",
			FirstNameTH: "Test",
			LastNameTH:  "User",
			FirstNameEN: "Test",
			LastNameEN:  "User",
			Email:       "gg@gg.com",
			Phone:       "1234567890",
			Department:  "Computer Science",
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedAt:   gorm.DeletedAt{},
		}

		// config database
		cfg := database.ConfigDatabaseWithPath("../../.env")

		// connect to database
		db, err := database.NewMySQL(cfg)
		assert.NoError(err, "failed to connect to database")

		// Close database connection
		defer func() {
			database.Close(db)
		}()

		// Insert user to database
		err = db.Create(&user).Error
		assert.NoError(err, "failed to insert user")
	})
}

func TestRetrieveUser(t *testing.T) {
	t.Run("Retrieve User", func(t *testing.T) {
		assert := assert.New(t)

		// config database
		cfg := database.ConfigDatabaseWithPath("../../.env")
		db, err := database.NewMySQL(cfg)
		assert.NoError(err, "failed to connect to database")

		defer func() {
			database.Close(db)
		}()

		// Retrieve user from database
		retrievedUser := userpkg.User{}
		err = db.First(&retrievedUser, "user_id = ?", "12345").Error
		assert.NoError(err, "failed to retrieve user")
		assert.Equal("testuser", retrievedUser.Username, "username mismatch")
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Update User", func(t *testing.T) {
		assert := assert.New(t)

		// config database
		cfg := database.ConfigDatabaseWithPath("../../.env")
		db, err := database.NewMySQL(cfg)
		assert.NoError(err, "failed to connect to database")

		defer func() {
			database.Close(db)
		}()

		// Update user in database
		retrievedUser := userpkg.User{}
		err = db.First(&retrievedUser, "user_id = ?", "12345").Error
		assert.NoError(err, "failed to retrieve user for update")

		updatedEmail := "updated@gg.com"
		err = db.Model(&retrievedUser).Update("email", updatedEmail).Error
		assert.NoError(err, "failed to update user email")

		// Verify the update
		err = db.First(&retrievedUser, "user_id = ?", "12345").Error
		assert.NoError(err, "failed to retrieve user after update")
		assert.Equal(updatedEmail, retrievedUser.Email, "email mismatch after update")
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Delete User", func(t *testing.T) {
		assert := assert.New(t)

		// config database
		cfg := database.ConfigDatabaseWithPath("../../.env")
		db, err := database.NewMySQL(cfg)
		assert.NoError(err, "failed to connect to database")

		defer func() {
			database.Close(db)
		}()

		// Delete user from database
		retrievedUser := userpkg.User{}
		err = db.First(&retrievedUser, "user_id = ?", "12345").Error
		assert.NoError(err, "failed to retrieve user for deletion")

		err = db.Delete(&retrievedUser).Error
		assert.NoError(err, "failed to delete user")

		// Verify deletion
		err = db.First(&retrievedUser, "user_id = ?", "12345").Error
		assert.Error(err, "expected error when retrieving deleted user, got none")
	})
}
