package bbs

import (
	"wall-bbs/app/bbs/pkg/ping"
	"wall-bbs/app/bbs/pkg/posts"
	"wall-bbs/app/bbs/pkg/user"
)

func Init() {
	ping.Init()
	user.Init()
	posts.Init()
}
