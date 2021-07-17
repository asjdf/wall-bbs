package httpEngine

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"wall-bbs/service/auth"
	"wall-bbs/service/config"
)

var httpEngineService *service

type service struct {
	HttpEngine *gin.Engine
	config config.HttpEngine
}

func Init() {
	httpEngineService = new(service)
	httpEngineService.config = config.GetHttpEngineConfig()
	httpEngineService.HttpEngine = gin.Default()
	httpEngineService.CORS()
	httpEngineService.HttpEngine.Use(auth.Middleware)

}

// CORS add cors middleware
func (s *service) CORS() {
	cf := cors.Config{
		AllowOrigins: []string{
			"https://cs.ptlzwnq.cn",
			"http://cs.ptlzwnq.cn",
			"http://localhost:8081",
			"https://localhost:8081",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "OPTION"},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type",
			"Authorization", "User-Agent",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	//cf.AllowAllOrigins = true

	s.HttpEngine.Use(cors.New(cf))
}
