package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"github.com/witchakornb/student-management-system/Entity/User"
	"github.com/witchakornb/student-management-system/Config"
)

func TestNewMySQL(t *testing.T) {
	
	// load environment variables
	if err := config.LoadEnvWithPath("../.env"); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name: "Valid configuration",
			cfg: Config{
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
				DBName:   os.Getenv("DB_NAME"),
			},
			wantErr: false,
		},
		{
			name: "Invalid configuration",
			cfg: Config{
				Host:     "invalidhost",
				Port:     "3306",
				Username: "root",
				Password: "example",
				DBName:   "example_db",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := NewMySQL(tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, db)
			}
		})
	}
}

// TestMigrate function
func TestMigrate(t *testing.T) {

	// load environment variables
	if err := config.LoadEnvWithPath("../.env"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		db     *gorm.DB
		models []interface{}
	}
	var db *gorm.DB
	var err error
	db, err = NewMySQL(Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	})
	assert.NoError(t, err)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid database connection",
			args: args{
				db:     db,
				models: []interface{}{&user.User{}},
			},
			wantErr: false,
		},
		{
			name: "Invalid database connection",
			args: args{
				db:     nil,
				models: []interface{}{&user.User{}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Migrate(tt.args.db, tt.args.models...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
