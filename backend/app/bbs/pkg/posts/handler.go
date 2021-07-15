package posts

import (
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
	boards := httpEngine.Group("/boards")
	boards.POST("/new", handleNewBoard)
	boards.GET("/list", handleListBoards)
}

func handleNewPost(c *gin.Context) {
	authResp := auth.GetAuthStatus(c)
	if !authResp.Found {
		respondTemplate.RespondJson(c, 40000, "尚未登陆")
	}
	if authResp.Token.Right == "-1" {
		respondTemplate.RespondJson(c, 40000, "无发布权限")
	}

	postData := struct {
		ParentID string `json:"parent_id"`
		Content    string `json:"content"`
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
	result := db.MySQL().Model(&post{}).Where("pid = ?",postData.ParentID).First(&parentPost)
	if result.RowsAffected != 1{
		respondTemplate.RespondJson(c, 40000, "父节点无效")
		return
	}
	//todo：下面存content之前要加上防xss的过滤器 以及Statue的默认值应该为可配置的
	post := post{
		PostUuid:   uuid.NewV4().String(),
		Uuid:       authResp.Token.Uuid,
		ParentUuid: parentPost.PostUuid,
		Content:    postData.Content,
		Statue:     0,
	}
	if err := db.MySQL().Create(&post).Error; err != nil {
		respondTemplate.RespondJson(c, 50000, "无法存入数据库")
		return
	}
	respondTemplate.RespondJsonWithData(c, 20000, "发布成功", gin.H{
		"puuid": post.PostUuid,
		"pid":   post.Pid,
	})
	return
}

func handleNewBoard(c *gin.Context) {
	authResp := auth.GetAuthStatus(c)
	if !authResp.Found {
		respondTemplate.RespondJson(c, 40000, "尚未登陆")
	}
	if authResp.Token.Right != "1" {
		respondTemplate.RespondJson(c, 40000, "无新建权限")
	}

	postData := struct {
		ParentID string `json:"parent_id"`
		Content    string `json:"content"`
	}{}

	if err := c.ShouldBindJSON(&postData); err != nil {
		respondTemplate.RespondJson(c, 40000, "无法解析JSON")
		return
	}

	if postData.ParentID == "" {
		respondTemplate.RespondJson(c, 40000, "参数不全")
		return
	}

	//todo：下面存content之前要加上防xss的过滤器 以及Statue的默认值应该为可配置的
	post := post{
		PostUuid:   uuid.NewV4().String(),
		Uuid:       authResp.Token.Uuid,
		ParentUuid: "",
		Content:    postData.Content,
		Statue:     0,
	}
	if err := db.MySQL().Create(&post).Error; err != nil {
		respondTemplate.RespondJson(c, 50000, "无法存入数据库")
		return
	}
	respondTemplate.RespondJsonWithData(c, 20000, "发布成功", gin.H{
		"puuid": post.PostUuid,
		"pid":   post.Pid,
	})
	return
}

func handlelListPosts(c *gin.Context) {
	postData := struct {
		ParentID string `json:"parent_id"`
	}{}

	if err := c.ShouldBindJSON(&postData); err != nil {
		respondTemplate.RespondJson(c, 40000, "无法解析JSON")
		return
	}

	parentPost := post{}
	result := db.MySQL().Model(&post{}).Where("pid = ?",postData.ParentID).First(&parentPost)
	if result.RowsAffected != 1{
		respondTemplate.RespondJson(c, 40000, "父节点无效")
		return
	}

	var posts []post
	db.MySQL().Select("name", "age").Where("parent_uuid = ?",parentPost.PostUuid).Find(&posts)

	respondTemplate.RespondJsonWithData(c, 20000, "success", posts)
}

func handleListBoards(c *gin.Context) {
	var posts []post
	db.MySQL().Select("name", "age").Where("parent_uuid = ?","").Find(&posts)

}