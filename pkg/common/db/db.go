package db

import (
	"gin_test_prjct/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln("Не удалось подключиться к БД", err)
	}

	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatalf("Ошибка миграции Book: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Ошибка миграции User: %v", err)
	}

	return db
}
