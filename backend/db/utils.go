package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDSN(a struct {
	Addr string
	User string
	Pass string
	DB   string
}) string {
	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", a.User, a.Pass, a.Addr, a.DB)
}

func CreateDB(a struct {
	Addr string
	User string
	Pass string
	DB   string
}) (*gorm.DB, error) {
	fmt.Println(CreateDSN(a))
	return gorm.Open(mysql.Open(CreateDSN(a)), &gorm.Config{PrepareStmt: true})
}