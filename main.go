package main

import (
	"gails/app/models"
	"gails/app/pkg/config"
	"gails/app/routes"

	"github.com/foolin/goview/supports/ginview"
)

func init() {
	config.Init()
	models.Init()
}

func main() {
	r := routes.InitRouter()

	//new template engine
	r.HTMLRender = ginview.Default()

	//TODO：
	// 1. sessions:github.com/gin-contrib/sessions
	// 2. csrf:https://github.com/utrack/gin-csrf
	// 3. cache
	// 4. 全局错误处理

	r.Run()
}
