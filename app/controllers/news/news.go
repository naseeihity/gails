package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 展示抓取消息
func Index(c *gin.Context) {
	isSignIn := c.GetBool("state_isUserSignIn")

	c.HTML(http.StatusOK, "news", gin.H{
		"Title":    "News",
		"Active":   "news",
		"IsSginIn": isSignIn,
	})
}
