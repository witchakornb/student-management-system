package main

import (
	"fmt"

	"github.com/witchakornb/student-management-system/Database"
	"github.com/witchakornb/student-management-system/Entity/User"
	"github.com/witchakornb/student-management-system/Entity/Student"
	"github.com/witchakornb/student-management-system/Config"
)

func main() {
	// load environment variables
	if err := config.LoadEnv(); err != nil {
		fmt.Println(err)
		return
	}
		
	// config database
	cfg := database.ConfigDatabase()

	// connect to database
	db, err := database.NewMySQL(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// migrate database
	if err := database.Migrate(db, &user.User{}, &student.Student{}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Database migrated successfully")

	// Close database connection
	defer func() {
		database.Close(db)
	}()

	// Select all users
	users := []user.User{}
	if err := db.Find(&users).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Users:" , users)

	// Select all students
	students := []student.Student{}
	if err := db.Find(&students).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Students:" , students)

	// Select student by student_id
	retrievedStudent := student.Student{}
	if err := db.Preload("User").First(&retrievedStudent, "student_id = ?", "12345").Error; err != nil {
		fmt.Println("Error retrieving student:", err)
		return
	}
	fmt.Printf("Student: %+v\n", retrievedStudent)
	fmt.Printf("Student's User Email: %s\n", retrievedStudent.User.Email)
	

}