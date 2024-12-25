package students

import (
	"github.com/witchakornb/student-management-system/Entity/User"
)

type (
	// Student struct
	Student struct {
		StudentID string `json:"student_id" gorm:"primary_key"`
		Student   user.User `json:"student" gorm:"foreignkey:StudentID"`
	}
)
