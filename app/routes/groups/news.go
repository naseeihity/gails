package groups

import (
	"gails/app/controllers/news"

	"github.com/gin-gonic/gin"
)

//InitNews :
func InitNews(r *gin.Engine) {
	n := r.Group("/news")

	{
		n.GET("/", news.Index)
	}

}
