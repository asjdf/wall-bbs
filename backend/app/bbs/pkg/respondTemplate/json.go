package respondTemplate

import "github.com/gin-gonic/gin"

func RespondJsonWithData(c *gin.Context,code int, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":msg,
		"data":data,
	})
}

func RespondJson(c *gin.Context,code int, msg string) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":msg,
	})
}