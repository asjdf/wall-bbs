package httpEngine

import "github.com/gin-gonic/gin"

func Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return httpEngineService.HttpEngine.Group(relativePath, handlers...)
}

func Run() {
	err := httpEngineService.HttpEngine.Run(httpEngineService.config.Port)
	if err != nil {
		panic(err)
	}
}
