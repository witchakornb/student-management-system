package database

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"github.com/witchakornb/student-management-system/Config"
)

// Config struct
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// Function Config Database
func ConfigDatabase() Config {

	// load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	cfg := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return cfg
}

// Function Config Database with path
func ConfigDatabaseWithPath(path string) Config {
	
	// load environment variables
	if err := config.LoadEnvWithPath(path); err != nil {
		log.Fatal(err)
	}

	cfg := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return cfg
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

// Close function
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return errors.New("failed to get database instance: " + err.Error())
	}
	return sqlDB.Close()
}