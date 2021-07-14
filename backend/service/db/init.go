package db

import (
	"gorm.io/gorm"
	"wall-bbs/service/config"
)

var dbService *service

type service struct {
	MySQL *gorm.DB
	config config.MySQL
}

func Init() {
	dbService = new(service)
	db, err := CreateDB(config.GetMySQLConfig())
	if err != nil {
		panic(err)
	}
	dbService.MySQL = db
}
