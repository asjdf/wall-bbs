package user

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
	"wall-bbs/service/db"
)

func GetUserInfoWithUid(uid uint) (*User, error) {
	user := new(User)
	result := db.MySQL().Model(User{}).Where("uid = ?", uid).First(user)
	return user, result.Error
}

func GetUserInfoWithTel(tel string) (*User, error) {
	user := new(User)
	result := db.MySQL().Model(User{}).Where("tel = ?", tel).First(user)
	return user, result.Error
}

func NewUser(name string, tel string, password string, right string) (uid uint, err error) {
	user := User{
		Uuid:     uuid.NewV4().String(),
		Name:     name,
		Tel:      tel,
		Password: password,
		Right:    right,
	}
	result := db.MySQL().Create(&user)
	if result.Error != nil {
		if result.Error.Error()[:10] == "Error 1062" {
			return 0, errors.New("用户已存在")
		}
	}
	if result.RowsAffected == 1 {
		return user.Uid, nil
	}

	return 0, errors.New("新建用户失败")
}

func VerifyCardCode(name string, cardCode string) bool {
	cardInfo := new(CardInfo)
	result := db.MySQL().Model(CardInfo{}).Where("card_code = ?", cardCode).First(&cardInfo)
	if result.RowsAffected == 1 {
		if cardInfo.RealName == name {
			return true
		}
	} else {
		return false
	}
	return false
}

func UpdateCardInfo() []error {
	request := gorequest.New()
	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; U; Android 2.3.6; zh-cn; GT-S5660 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 MicroMessenger/4.5.255")
	_, body, errs := request.Get("http://xfxt.ptlz.com.cn/wxyjt/online/selUser").End()
	if errs != nil {
		return errs
	}

	cardInfoT := CardInfoT{}
	err := json.Unmarshal([]byte(body), &cardInfoT)
	if err != nil {
		return []error{err}
	}

	err = db.MySQL().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"card_code", "balance"}),
	}).Create(&cardInfoT.Data).Error
	if err != nil {
		return []error{err}
	}

	return nil
}
