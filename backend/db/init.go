package db

import "gorm.io/gorm"
import "wall-bbs/config"

var s *service

type service struct {
	DB *gorm.DB
}

func init() {
	s = new(service)
	db, err := CreateDB(config.Get().Conf.DB)
	if err != nil {
		panic(err)
	}
	s.DB = db
}

func Get() *service {
	return s
}