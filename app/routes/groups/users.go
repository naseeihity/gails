package groups

import (
	"github.com/gin-gonic/gin"
)

//InitUsers :
func InitUsers(r *gin.Engine) {
	user := r.Group("/user")

	{
		user.POST("/sign_in", SignIn)
	}

}

func SignIn(c *gin.Context) {

}
