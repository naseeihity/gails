package routes

import (
	"gails/app/routes/groups"

	"github.com/gin-gonic/gin"
)

//InitRouter :
func InitRouter() *gin.Engine {
	r := gin.Default()

	groups.InitHome(r)
	groups.InitUsers(r)
	groups.InitNews(r)
	groups.InitPasswords(r)

	return r
}
