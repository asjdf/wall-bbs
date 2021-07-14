package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strings"
)

const (
	AUTH_FLAG          = "GatewayAuth"
	PREFIX_CACHE_TOKEN = "auth_token_"
	TOKEN_EXPIRE       = 86400 * 30 //60天
)

func Middleware(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	tokenStrFromQuery := c.Query("auth")

	var (
		tokenStr string
		authType string
	)

	authArr := strings.Split(authorization, " ")
	_, err := uuid.FromString(tokenStrFromQuery)

	switch {
	case len(authArr) == 2:
		tokenStr = authArr[1]
		authType = authArr[0]
	case err == nil:
		tokenStr = tokenStrFromQuery
		authType = "token"
	default:
		c.Set(AUTH_FLAG, nil)
		return
	}

	switch authType {
	case "token":
		//token auth
		token := new(authToken)
		//get token from cache
		err = GetCacheT(PREFIX_CACHE_TOKEN+tokenStr, token, 3600)
		//get token from mysql
		if err != nil {
			token = LookUpAuthToken(tokenStr)
			if token.Valid() {
				SetCache(PREFIX_CACHE_TOKEN+tokenStr, token, 3600)
				token.ExtendExpireTime()
			}
		}
		switch {
		//token 不存在
		case !token.Valid():
			c.Set(AUTH_FLAG, MakeAuthNotFound())
			return
		default:
			r := MakeSuccessAuth(token)
			c.Set(AUTH_FLAG, r)
			return
		}
	default:
		//default
		c.Set(AUTH_FLAG, MakeAuthNotFound())
		return
	}
}

func MakeSuccessAuth(token *authToken) *authResp {
	return &authResp{
		Found:    true,
		Token:    token,
		Err:      nil,
		ErrCode:  0,
		HTTPCode: 200,
	}
}

func MakeAuthNotFound() *authResp {
	return &authResp{
		Found:    false,
		Token:    nil,
		Err:      errors.New("unauthorized"),
		ErrCode:  40100,
		HTTPCode: 401,
	}
}

func GetAuthStatus(c *gin.Context) *authResp {
	auth, b := c.Get(AUTH_FLAG)
	if !b || auth == nil {
		return MakeAuthNotFound()
	}
	return auth.(*authResp)
}
