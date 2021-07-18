package user

import (
	"github.com/gin-gonic/gin"
	"wall-bbs/app/bbs/pkg/respondTemplate"
	"wall-bbs/service/auth"
	"wall-bbs/service/httpEngine"
)

func Init() {
	if err := initDB(); err != nil {
		panic(err)
	}
	user := httpEngine.Group("/user")
	user.POST("/register", handleUserRegister)
	user.POST("/login", handleUserLogin)
	user.POST("/verify-card-info", handleVerifyCardInfo)
	user.GET("/update-card-info", handleUpdateCardInfo)
	user.GET("/logout", handleLogout)
}

func handleUserRegister(c *gin.Context) {
	var formData RegisterForm

	if err := c.ShouldBindJSON(&formData); err != nil {
		respondTemplate.RespondJson(c, 40000, "无法解析JSON")
		return
	}

	if formData.Name == "" || formData.Password == "" || formData.CardCode == "" || formData.Tel == "" || formData.RealName == "" {
		respondTemplate.RespondJson(c, 40000, "参数不全")
		return
	}
	if VerifyCardCode(formData.RealName, formData.CardCode) == true {
		if uid, err := NewUser(formData.Name, formData.Tel, formData.Password, "0", formData.RealName, formData.CardCode); err == nil {
			respondTemplate.RespondJsonWithData(c, 20000, "success", gin.H{"uid": uid})
			return
		} else if err.Error() == "用户已存在" {
			respondTemplate.RespondJson(c, 40300, "注册用户失败，用户已存在")
			return
		} else {
			respondTemplate.RespondJson(c, 50000, "注册用户出现未知错误")
			return
		}
	} else {
		respondTemplate.RespondJson(c, 40000, "卡号验证错误，请重新核对")
		return
	}
	//respondTemplate.RespondJson(c, 40300, "服务流程出现错误")
	//return
}

func handleUserLogin(c *gin.Context) {
	var formData LoginForm

	if err := c.ShouldBindJSON(&formData); err != nil {
		respondTemplate.RespondJson(c, 40000, "无法解析JSON")
		return
	}

	if formData.Tel == "" || formData.Password == "" {
		respondTemplate.RespondJson(c, 40000, "参数不全")
		return
	}

	userData, err := GetUserInfoWithTel(formData.Tel)
	if err != nil {
		if err.Error() == "record not found" {
			respondTemplate.RespondJson(c, 40000, "用户不存在")
			return
		}
		respondTemplate.RespondJson(c, 50000, "发生非预期错误，请稍后重试")
		return
	}
	//respondTemplate.RespondJsonWithData(c, 20000,"",userData)
	if formData.Password == userData.Password {
		token, err := auth.CreateAuthToken(userData.Uuid, userData.Uid, userData.Tel, userData.Right)
		if err != nil {
			respondTemplate.RespondJson(c, 50000, "登录出错")
			return
		}
		respondTemplate.RespondJsonWithData(c, 20000, "登录成功", gin.H{
			"token": token,
			"uid":   userData.Uid,
			"right": userData.Right,
		})
		return
	} else {
		respondTemplate.RespondJson(c, 40300, "密码错误")
		return
	}
}

func handleUpdateCardInfo(c *gin.Context) {
	errs := UpdateCardInfo()
	if errs != nil {
		respondTemplate.RespondJson(c, 50000, "fail")
		return
	}
	respondTemplate.RespondJson(c, 20000, "success")
	return
}

func handleVerifyCardInfo(c *gin.Context) {
	formData := struct {
		RealName string `json:"real_name"`
		CardId   string `json:"card_id"`
	}{}
	if err := c.ShouldBindJSON(&formData); err != nil {
		respondTemplate.RespondJson(c, 40000, "无法解析JSON")
		return
	}
	if formData.RealName == "" || formData.CardId == "" {
		respondTemplate.RespondJson(c, 40000, "参数不全")
		return
	}

	verifyResult := VerifyCardCode(formData.RealName, formData.CardId)
	respondTemplate.RespondJsonWithData(c, 20000, "success", gin.H{"result": verifyResult})
	return
}

func handleLogout(c *gin.Context) {
	authResp := auth.GetAuthStatus(c)
	if !authResp.Found {
		respondTemplate.RespondJson(c, 40000, "没登录，退出个鬼")
	}

	auth.UnauthorizeToken(authResp.Token.Token)
	respondTemplate.RespondJson(c, 20000, "success")
	return
}
