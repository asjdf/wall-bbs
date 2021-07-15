package posts

import (
	"database/sql"
	"time"
)

//板块、文章及评论
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
