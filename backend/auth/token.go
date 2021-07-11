package auth

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)
func CreateAuthToken(UID string, right string) (token string, err error) {
	token = uuid.NewV4().String()
	t := &accessToken{
		AccessToken: token,
		UID: UID,
		Right: right,
	}

	tx := s.AuthDB.Begin()
	if tx.Create(t).RowsAffected == 1 {
		tx.Commit()
		return token, nil
	}
	tx.Rollback()
	return "", errors.New("create token error")
}
