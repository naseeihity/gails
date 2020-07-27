package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Home",
	})
}

//About 关于页面
func About(c *gin.Context) {

}
