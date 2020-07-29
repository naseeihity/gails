package articles

import (
	"gails/app/helpers"
	"gails/app/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//NewArticle 新建文章页面
func NewArticle(c *gin.Context) {

}

//Show 文章详情
func Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := models.GetArticleByID(id)
	if err != nil || article == nil {
		log.Println("models.GetArticleByID ==> ", err)
		c.Redirect(http.StatusFound, "/")
		return
	}
	author, err := models.GetUserByID(article.UserID)
	if err != nil {
		log.Println("models.GetUserByID ==> ", err)
	}

	data := gin.H{
		"article": article,
		"author":  author.Name,
	}

	c.HTML(http.StatusOK, "article", helpers.CommonHTMLRes(article.Title, c, data))

}

//Update 更新文章
func Update(c *gin.Context) {

}

//Edit 编辑页面
func Edit(c *gin.Context) {

}

//Create 新建文件
func Create(c *gin.Context) {

}

// ====== Middlewares ======

//CheckArticleOwner 检查文章权限
func CheckArticleOwner() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//CheckParamsBody 检查参数
func CheckParamsBody() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
