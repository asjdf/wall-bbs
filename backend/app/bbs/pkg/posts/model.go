package posts

import (
	"database/sql"
	"time"
)

//文章及评论
type post struct {
	CreatedAt  time.Time `gorm:"index"`
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
	PostUuid   string       `gorm:"uniqueIndex;size:36"`
	Pid        uint         `gorm:"autoIncrement;index"`
	Uuid       string       `gorm:"index;size:36"`
	ParentUuid string       `gorm:"index;size:36"`
	Content    string
	Statue     uint `gorm:"default:1"`
}

//板块
type board struct {
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Uuid      string       `gorm:"uniqueIndex;size:36"`
	ID        uint         `gorm:"autoIncrement;index"`
	Pic       string
	Name      string
	Content   string
}
