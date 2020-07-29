package home

import (
	"gails/app/helpers"
	"gails/app/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Index 首页
func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page == 0 {
		page = 1
	}

	articles, err := models.GetArticlesByPage(page)
	if err != nil {
		log.Println("models.GetArticlesByPage ==> ", err)
	}
	count, err := models.GetArticlesCount()
	if err != nil {
		log.Println("models.GetArticlesCount ==> ", err)
	}

	data := gin.H{
		"articles": articles,
		"count":    count,
	}

	c.HTML(http.StatusOK, "index", helpers.CommonHTMLRes("Home", c, data))
}

//About 关于页面
func About(c *gin.Context) {

	c.HTML(http.StatusOK, "about", helpers.CommonHTMLRes("About", c))
}
