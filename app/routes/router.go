package routes

import (
	"gails/app/routes/groups"

	"github.com/gin-gonic/gin"
)

//InitRouter :
func InitRouter() *gin.Engine {
	r := gin.Default()

	groups.InitUsers(r)

	return r
}
