package db

import "gorm.io/gorm"

func MySQL() *gorm.DB {
	return dbService.MySQL
}
