package database

import (
	"github.com/Deepanshu276/AccuKnox-Task/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("Note.db"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&model.User{}, &model.Note{})
}