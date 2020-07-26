package groups

import (
	"gails/app/controllers/home"

	"github.com/gin-gonic/gin"
)

//InitHome :
func InitHome(r *gin.Engine) {

	{
		r.GET("/about", home.About)
		r.GET("/", home.Index)
	}

}
