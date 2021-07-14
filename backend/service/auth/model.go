package auth

import (
	"gorm.io/gorm"
	"time"
)

type authToken struct {
	gorm.Model
	Token      string `gorm:"uniqueIndex;size:36"`
	Uuid       string `gorm:"index"`
	UID        uint `gorm:"index"`
	Right      string
	Tel        string `gorm:"index"`
	ExpireTime time.Time
}

type authResp struct {
	Found    bool
	Token    *authToken
	Err      error
	ErrCode  int
	HTTPCode int
}

func (a authResp) Valid() (status bool) {
	if a.Err == nil && a.Token != nil && a.Token.UID != 0 {
		status = true
	}
	return
}
