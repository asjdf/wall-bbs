package posts

import (
	"github.com/gin-gonic/gin"
	"wall-bbs/app/bbs/pkg/respondTemplate"
	"wall-bbs/service/auth"
	"wall-bbs/service/db"
	"wall-bbs/service/httpEngine"
)

func Init() {
	if err := db.MySQL().AutoMigrate(&post{}); err != nil{
		panic(err)
	}
	posts := httpEngine.Group("/posts")
	posts.GET("/post", handlePost)
}

func handlePost(c *gin.Context) {
	authResp := auth.GetAuthStatus(c)
	if !authResp.Found {
		respondTemplate.RespondJson(c, 40000, "尚未登陆")
	}
	if authResp.Token.Right == "-1" {
		respondTemplate.RespondJson(c, 40000, "无发布权限")
	}


}