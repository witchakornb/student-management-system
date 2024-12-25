package user

import (
	"time"

	"gorm.io/gorm"
)

type (
	// User struct
	User struct {
		UserID      string         `json:"user_id" gorm:"primaryKey"`       // Explicit primaryKey tag
		Username    string         `json:"username" gorm:"not null;unique"` // Explicit not null and unique constraints
		Password    string         `json:"password"`
		Role        string         `json:"role" gorm:"default:'student'"` // Default value for the 'role' field
		FirstNameTH string         `json:"first_name_th"`                 // JSON keys in camelCase
		LastNameTH  string         `json:"last_name_th"`
		FirstNameEN string         `json:"first_name_en"`
		LastNameEN  string         `json:"last_name_en"`
		Email       string         `json:"email" gorm:"not null;unique"` // Explicit not null and unique constraints
		Phone       string         `json:"phone"`
		Department  string         `json:"department"`
		CreatedAt   time.Time      `json:"created_at"` // Time fields are mapped by GORM automatically
		UpdatedAt   time.Time      `json:"updated_at"`
		DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"` // Use 'DeletedAt' for soft delete
	}
)
