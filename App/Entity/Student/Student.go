package student

import (
	"github.com/witchakornb/student-management-system/Entity/User"
	"gorm.io/gorm"
	"time"
)

// Student struct
type Student struct {
	StudentID        string         `json:"student_id" gorm:"primaryKey"`
	StudentCreatedAt time.Time      `json:"student_created_at"`
	StudentUpdatedAt time.Time      `json:"student_updated_at"`
	StudentDeletedAt gorm.DeletedAt `json:"student_deleted_at" gorm:"index"`
	User             user.User      `json:"user" gorm:"foreignKey:StudentID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Correct foreign key relationship
}