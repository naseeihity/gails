package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 首页
func Index(c *gin.Context) {
	isSignIn := c.GetBool("state_isUserSignIn")

	c.HTML(http.StatusOK, "index", gin.H{
		"Title":    "Home",
		"Active":   "home",
		"IsSginIn": isSignIn,
	})
}

//About 关于页面
func About(c *gin.Context) {
	isSignIn := c.GetBool("state_isUserSignIn")

	c.HTML(http.StatusOK, "about", gin.H{
		"Title":    "About",
		"Active":   "about",
		"IsSginIn": isSignIn,
	})
}
