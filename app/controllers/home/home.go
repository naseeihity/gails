package home

import (
	"gails/app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 首页
func Index(c *gin.Context) {
	var username string
	isSignIn := c.GetBool("state_isUserSignIn")
	if user, _ := c.Get("state_curUser"); user != nil {
		if curUser, ok := user.(*models.Users); ok {
			username = curUser.Name
		}
	}

	log.Println(isSignIn, "Index Redirect User ==> ", username)

	c.HTML(http.StatusOK, "index", gin.H{
		"Title":    "Home",
		"Active":   "home",
		"IsSginIn": isSignIn,
		"curUser":  username,
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
