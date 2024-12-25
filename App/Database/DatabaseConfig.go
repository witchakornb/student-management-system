package database

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config struct
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// NewMySQL function
func NewMySQL(cfg Config) (*gorm.DB, error) {
	// Create a DSN (Data Source Name) for the MySQL connection
	dsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to open database connection: " + err.Error())
	}

	// Check the database connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New("failed to get database instance: " + err.Error())
	}

	// Ping the database to ensure the connection is valid
	if err := sqlDB.Ping(); err != nil {
		return nil, errors.New("failed to ping database: " + err.Error())
	}

	return db, nil
}

// Migrate function
func Migrate(db *gorm.DB, models ...interface{}) error {
	if db == nil {
		return errors.New("database connection is nil")
	}
	return db.AutoMigrate(models...)
}