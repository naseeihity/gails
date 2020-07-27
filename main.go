package main

import (
	"gails/app/models"
	"gails/app/pkg/config"
	"gails/app/pkg/redis"
	"gails/app/routes"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
)

func init() {
	config.Init()
	redis.Init()
	models.Init()
}

func main() {
	r := routes.InitRouter()

	//new template engine
	//TODO:only use gin as API server, user React+NextJS to do router and render
	r.HTMLRender = ginview.New(goview.Config{
		Root:      "app/views",
		Extension: ".html",
		Master:    "layout/master",
	})

	r.Run(":" + config.ServerCfg.Port)
}
