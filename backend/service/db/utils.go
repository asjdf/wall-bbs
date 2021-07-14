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
	ConnectTimeout uint
}) (*gorm.DB, error) {
	//fmt.Println(CreateDSN(a))
	config := struct {
		Addr string
		User string
		Pass string
		DB   string
	}{a.Addr,a.User,a.Pass,a.DB}
	DB,err := gorm.Open(mysql.Open(CreateDSN(config)), &gorm.Config{PrepareStmt: true})
	return DB,err
}