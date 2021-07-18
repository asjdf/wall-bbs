package user

import (
	"database/sql"
	"time"
	"wall-bbs/service/db"
)

func initDB() error {
	return db.MySQL().AutoMigrate(&User{}, &CardInfo{})
}

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Uuid      string       `gorm:"uniqueIndex;size:36"`
	Uid       uint         `gorm:"autoIncrement;index"`
	Name      string
	Tel       string `gorm:"uniqueIndex;size:11"`
	RealName  string
	CardCode  string `gorm:"size:20"`
	Password  string
	Right     string
}

type RegisterForm struct {
	Tel      string `json:"tel"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RealName string `json:"real_name"`
	CardCode string `json:"card_code"`
}

type LoginForm struct {
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

type CardInfoT struct {
	Msg    string     `json:"msg"`
	Data   []CardInfo `json:"data"`
	Status int        `json:"status"`
}

type CardInfo struct {
	UserNum    string  `json:"user_num"`
	CardCode   string  `json:"card_code" gorm:"index"`
	Balance    float64 `json:"balance"`
	UserId     int     `json:"user_id" gorm:"uniqueIndex"`
	GroupName  string  `json:"group_name"`
	SchoolName string  `json:"school_name"`
	RealName   string  `json:"real_name"`
}
