package db

import (
	"log"

	"github.com/14jasimmtp/notebook-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
) 

type DB struct{
	DB *gorm.DB
}

func Connection(url string) *gorm.DB {
	db,err:=gorm.Open(postgres.Open(url),&gorm.Config{})
	if err != nil {
		log.Fatal("error while connecting to db",err)
	}
	db.AutoMigrate(&models.Notebook{})

	return db
}
