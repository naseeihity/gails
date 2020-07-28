package groups

import (
	"gails/app/controllers/users"

	"github.com/gin-gonic/gin"
)

//InitUsers :
func InitUsers(r *gin.Engine) {
	u := r.Group("/user")

	{
		u.GET("/sign_in", users.SignInPage)
		u.POST("/sign_in", users.SignIn)
		u.GET("/log_out", users.LogOut)
		u.GET("/", users.Index)
	}

}
