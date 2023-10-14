package database

import (
	"goApiStartetProject/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB(env *config.Env) *gorm.DB {

	db, err := gorm.Open(postgres.Open(env.DBUrl+"&application_name=$ "+env.DBName), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to Database")
	}
	log.Println("Connected to Database!")

	return db
}
