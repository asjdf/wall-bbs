package auth

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
	"wall-bbs/service/db"
)

func CreateAuthToken(Uuid string, uid uint, tel string, right string) (token string, err error) {
	token = uuid.NewV4().String()
	t := &authToken{
		Token:      token,
		Uuid:       Uuid,
		UID:        uid,
		Tel:        tel,
		Right:      right,
		ExpireTime: time.Unix(time.Now().Unix()+TOKEN_EXPIRE, 0),
	}

	tx := db.MySQL().Begin()
	if tx.Create(t).RowsAffected == 1 {
		tx.Commit()
		return token, nil
	}
	tx.Rollback()
	return "", errors.New("create token error")
}

func LookUpAuthToken(tokenString string) *authToken {
	token := new(authToken)
	db.MySQL().Where(&authToken{Token: tokenString}).
		Where(" expire_time > ? ", time.Now().Unix()).
		First(&token)
	return token
}

func (token authToken) ExtendExpireTime(){
	db.MySQL().Model(&authToken{}).Where(&token).Update("expire_time",time.Unix(time.Now().Unix()+TOKEN_EXPIRE, 0))
}

func (token authToken) Valid() bool {
	return token.Token != ""
}

func UnauthorizeToken(tokenStr string) {
	token := new(authToken)
	db.MySQL().Where(&authToken{Token: tokenStr}).First(&token)
	if token.Uuid == "" {
		return
	}
	tx := db.MySQL().Begin()
	defer tx.Rollback()
	if tx.Delete(token).RowsAffected == 1 {
		tx.Commit()
		DeleteCache(PREFIX_CACHE_TOKEN+tokenStr, true, true)
		return
	}

}
