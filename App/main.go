package main

import (
	"fmt"
	"os"

	"github.com/witchakornb/student-management-system/Database"
	"github.com/witchakornb/student-management-system/Entity/User"
	"github.com/witchakornb/student-management-system/Config"
)

func main() {
	// load environment variables
	if err := config.LoadEnv(); err != nil {
		fmt.Println(err)
		return
	}
		
	// config database
	cfg := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	// connect to database
	db, err := database.NewMySQL(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// migrate database
	if err := database.Migrate(db, &user.User{}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Database migrated successfully")
}