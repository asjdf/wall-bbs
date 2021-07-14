package ping

import (
	"github.com/gin-gonic/gin"
	"wall-bbs/service/auth"
	"wall-bbs/service/httpEngine"
)

func Init() {
	ping := httpEngine.Group("/ping")
	ping.GET("",handlePingPong)
}

func handlePingPong(c *gin.Context) {
	a,_ := c.Get(auth.AUTH_FLAG)
	c.JSON(200, gin.H{
		"msg": "pong",
		"auth": a,
	})
}
