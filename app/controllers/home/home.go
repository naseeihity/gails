package home

import (
	"gails/app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", helpers.CommonHTMLRes("Home", c))
}

//About 关于页面
func About(c *gin.Context) {

	c.HTML(http.StatusOK, "about", helpers.CommonHTMLRes("About", c))
}
