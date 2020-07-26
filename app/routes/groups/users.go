package groups

import (
	"gails/app/controllers/users"

	"github.com/gin-gonic/gin"
)

//InitUsers :
func InitUsers(r *gin.Engine) {
	user := r.Group("/user")

	{
		user.GET("/sign_in", users.SignInPage)
		user.POST("/sign_in", users.SignIn)
		user.GET("/log_out", users.LogOut)
		user.GET("/", users.Index)
	}

}
