package posts

import (
	"github.com/feiin/go-xss"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"wall-bbs/app/bbs/pkg/respondTemplate"
	"wall-bbs/service/auth"
	"wall-bbs/service/db"
	"wall-bbs/service/httpEngine"
)

func Init() {
	if err := db.MySQL().AutoMigrate(&post{}); err != nil {
		panic(err)
	}
	posts := httpEngine.Group("/posts")
	posts.POST("/new", handleNewPost)
	posts.POST("/list", handlelListPosts)
	//posts.POST("/detail", handleDetailPost)
	boards := httpEngine.Group("/boards")
	boards.POST("/new", handleNewBoard)
	boards.GET("/list", handleListBoards)
}

func handleNewPost(c *gin.Context) {
	authResp := auth.GetAuthStatus(c)
	if !authResp.Found {
		respondTemplate.RespondJson(c, 40001, "尚未登陆")
		return
	}
	if authResp.Token.Right == "-1" {
		respondTemplate.RespondJson(c, 40000, "无发布权限")
		return
	}

	postData := struct {
		ParentID  string `json:"parent_id"`
		Content   string `json:"content"`
		Anonymous bool   `json:"anonymous"`
	}{}

	if err := c.ShouldBindJSON(&postData); err != nil {
		respondTemplate.RespondJson(c, 40000, "无法解析JSON")
		return
	}

	if postData.ParentID == "" {
		respondTemplate.RespondJson(c, 40000, "参数不全")
		return
	}

	parentPost := post{}
	result := db.MySQL().Model(&post{}).Where("pid = ?", postData.ParentID).First(&parentPost)
	if result.RowsAffected != 1 {
		respondTemplate.RespondJson(c, 40000, "父节点无效")
		return
	}
	//todo：下面存content之前要加上防xss的过滤器 以及Statue的默认值应该为可配置的
	safeContent := xss.FilterXSS(postData.Content,xss.XssOption{
		WhiteList: xss.GetDefaultWhiteList(),
	})
	post := post{
		PostUuid:   uuid.NewV4().String(),
		Uuid:       authResp.Token.Uuid,
		ParentUuid: parentPost.PostUuid,
		Content:    safeContent,
		Statue:     0,
		Anonymous:  postData.Anonymous,
	}
	if err := db.MySQL().Create(&post).Error; err != nil {
		respondTemplate.RespondJson(c, 50000, "无法存入数据库")
		return
	}
	respondTemplate.RespondJson(c, 20000, "发布成功")
	return
}

func handlelListPosts(c *gin.Context) {
	postData := struct {
		ParentID string `json:"parent_id"`
		PageNum  uint   `json:"page_num"`
		Desc     bool   `json:"desc"`
	}{}

	if err := c.ShouldBindJSON(&postData); err != nil {
		respondTemplate.RespondJson(c, 40000, "无法解析JSON")
		return
	}
	if postData.PageNum == 0 {
		postData.PageNum = 1
	}

	orderStr := "created_at"
	if postData.Desc {
		orderStr += " desc"
	}
	var posts []post

	type postInfo struct {
		Content  string `json:"content"`
		PID      uint   `json:"pid"`
		UserName string `json:"user_name"`
		UID      uint   `json:"uid"`
		PostTime string `json:"post_time"`
	}
	type postsList struct {
		TotalPage  int64      `json:"total_page"`
		ParentPost postInfo   `json:"parent_post"`
		ChildList  []postInfo `json:"child_list"`
	}
	boardListOutput := postsList{}

	if postData.ParentID == "" {
		db.MySQL().Model(&post{}).Where("parent_uuid != ''").
			Where("statue = 1").
			Count(&boardListOutput.TotalPage)
		boardListOutput.TotalPage = boardListOutput.TotalPage/10+1
		db.MySQL().Where("parent_uuid != ''").
			Where("statue = 1").
			Joins("User").
			Order(orderStr).
			Limit(10).
			Offset(int((postData.PageNum - 1) * 10)).
			Find(&posts)
	} else {
		parentPost := post{}
		result := db.MySQL().Model(&post{}).
			Joins("User").
			Where("pid = ?", postData.ParentID).First(&parentPost)
		if result.RowsAffected != 1 {
			respondTemplate.RespondJson(c, 40000, "父节点无效")
			return
		}

		boardListOutput.ParentPost = postInfo{
			Content:  parentPost.Content,
			PID:      parentPost.Pid,
			UserName: parentPost.User.Name,
			UID:      parentPost.User.Uid,
			PostTime: parentPost.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		if parentPost.Anonymous == true {
			boardListOutput.ParentPost.UserName = "匿名"
			boardListOutput.ParentPost.UID = 0
		}

		db.MySQL().Model(&post{}).
			Where("parent_uuid = ?", parentPost.PostUuid).
			Where("statue = 1").
			Count(&boardListOutput.TotalPage)
		boardListOutput.TotalPage = boardListOutput.TotalPage/10+1
		db.MySQL().Where("parent_uuid = ?", parentPost.PostUuid).
			Where("statue = 1").
			Joins("User").
			Order(orderStr).
			Limit(10).
			Offset(int((postData.PageNum - 1) * 10)).
			Find(&posts)
	}

	for _, v := range posts {
		bInfo := postInfo{
			Content:  v.Content,
			PID:      v.Pid,
			UserName: v.User.Name,
			UID:      v.User.Uid,
			PostTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		if v.Anonymous == true {
			bInfo.UserName = "匿名"
			bInfo.UID = 0
		}
		boardListOutput.ChildList = append(boardListOutput.ChildList, bInfo)
	}

	respondTemplate.RespondJsonWithData(c, 20000, "success", boardListOutput)
}

//func handleDetailPost(c *gin.Context)  {
//	postData := struct {
//		Pid uint `json:"pid"`
//	}{}
//
//	if err := c.ShouldBindJSON(&postData); err != nil {
//		respondTemplate.RespondJson(c, 40000, "无法解析JSON")
//		return
//	}
//
//	post := post{}
//	db.MySQL().
//		Where("pid = ?", postData.Pid).
//		Joins("User").
//		First(&post)
//	output := struct {
//
//	}{}
//}

func handleNewBoard(c *gin.Context) {
	authResp := auth.GetAuthStatus(c)
	if !authResp.Found {
		respondTemplate.RespondJson(c, 40001, "尚未登陆")
		return
	}
	if authResp.Token.Right != "1" {
		respondTemplate.RespondJson(c, 40000, "无新建权限")
		return
	}

	postData := struct {
		Content string `json:"board_name"`
	}{}

	if err := c.ShouldBindJSON(&postData); err != nil {
		respondTemplate.RespondJson(c, 40000, "无法解析JSON")
		return
	}

	if postData.Content == "" {
		respondTemplate.RespondJson(c, 40000, "参数不全")
		return
	}

	//todo：下面存content之前要加上防xss的过滤器 以及Statue的默认值应该为可配置的
	safeContent := xss.FilterXSS(postData.Content,xss.XssOption{
		WhiteList: xss.GetDefaultWhiteList(),
	})
	post := post{
		PostUuid:   uuid.NewV4().String(),
		Uuid:       authResp.Token.Uuid,
		ParentUuid: "",
		Content:    safeContent,
		Statue:     0,
	}
	if err := db.MySQL().Create(&post).Error; err != nil {
		respondTemplate.RespondJson(c, 50000, "无法存入数据库")
		return
	}
	respondTemplate.RespondJson(c, 20000, "发布成功")
	return
}

func handleListBoards(c *gin.Context) {
	var posts []post
	db.MySQL().Where("parent_uuid = ?", "").
		Where("statue = 1").
		Find(&posts)

	type boardsInfo struct {
		BoardName string `json:"board_name"`
		BoardID   uint   `json:"board_id"`
	}
	type boardList struct {
		List []boardsInfo `json:"list"`
	}
	boardListOutput := boardList{}
	for _, v := range posts {
		bInfo := boardsInfo{
			BoardName: v.Content,
			BoardID:   v.Pid,
		}
		boardListOutput.List = append(boardListOutput.List, bInfo)
	}

	respondTemplate.RespondJsonWithData(c, 20000, "success", boardListOutput)
}
