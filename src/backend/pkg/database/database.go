package database

import (
	"log"
	"os"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"vue-gin-backend/internal/model"
)

var DB *gorm.DB

func Connect() {
	var err error
	
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "data.db"
	}
	
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	log.Println("Database connected successfully")
}

func Migrate() {
	err := DB.AutoMigrate(&model.User{}, &model.Micropost{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	
	log.Println("Database migration completed")
}

func SeedData() {
	var userCount int64
	DB.Model(&model.User{}).Count(&userCount)
	
	if userCount == 0 {
		users := []model.User{
			{Name: "Alice Johnson"},
			{Name: "Bob Smith"},
			{Name: "Charlie Brown"},
		}
		
		for _, user := range users {
			DB.Create(&user)
		}
		
		microposts := []model.Micropost{
			{Title: "Hello World!", UserID: 1},
			{Title: "Learning Go and Gin", UserID: 1},
			{Title: "GORM is awesome", UserID: 2},
			{Title: "Building APIs with Go", UserID: 2},
			{Title: "Vue.js frontend integration", UserID: 3},
		}
		
		for _, post := range microposts {
			DB.Create(&post)
		}
		
		log.Println("Sample data seeded successfully")
	}
}
