package db

import (
	"log"

	model "github.com/Amir-Hossein-shamsi/fiber-resapi/src/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	return db
}
