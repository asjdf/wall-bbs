package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wall-bbs/config"
)

type Service struct {
	Router *gin.Engine
}

func (s *Service) init() {
	s.initRouter()
	s.run()
}

func (s *Service) run() {
	fmt.Println("start listening port " + config.Get().Conf.Server.Port)
	err := s.Router.Run(config.Get().Conf.Server.Port)
	panic(err)
}