package auth

import "gorm.io/gorm"

type accessToken struct {
	gorm.Model
	AccessToken string `gorm:unique_index`
	UID string `gorm:index`
	Right string
}
