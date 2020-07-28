package news

import (
	"gails/app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 展示抓取消息
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "news", helpers.CommonHTMLRes("News", c))
}
