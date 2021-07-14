package service

import (
	"wall-bbs/service/app"
	"wall-bbs/service/auth"
	"wall-bbs/service/config"
	"wall-bbs/service/db"
	"wall-bbs/service/httpEngine"
)

func Init() {
	config.Init()
	db.Init()
	auth.Init()
	httpEngine.Init()
	app.Init()
}

func Run() {
	httpEngine.Run()
}
