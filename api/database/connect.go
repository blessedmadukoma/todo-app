package database

import (
	"github.com/blessedmadukoma/todo-vue-go/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=todo_vue_go port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	DB = database

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.Todo{}, &models.User{})
	fmt.Println("Database Migrated")
}
