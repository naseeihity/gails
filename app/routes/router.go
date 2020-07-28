package routes

import (
	"gails/app/middlewares"
	"gails/app/routes/groups"

	"github.com/gin-gonic/gin"
)

//InitRouter :初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "./app/assets")
	// global middlewares
	r.Use(middlewares.RedisSession(),
		middlewares.CORS(), middlewares.FlashMessage(), middlewares.AddState())

	//TODO：
	// 1. cache
	// 2. 全局错误处理

	// 路由分组
	groups.InitHome(r)
	groups.InitUsers(r)
	groups.InitNews(r)
	groups.InitPasswords(r)

	return r
}
