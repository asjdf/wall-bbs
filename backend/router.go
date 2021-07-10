package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Service) initRouter() {

	s.Router = gin.Default()
	//enable cors header
	s.CORS()
	//auth middleware request-id middleware


}

// CORS add cors middleware
func (s *Service) CORS() {
	cf := cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "OPTION"},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type",
			"Authorization", "User-Agent",
		},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}
	cf.AllowAllOrigins = true

	s.Router.Use(cors.New(cf))
}
