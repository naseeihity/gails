package main

import (
	"gails/app/models"
	"gails/app/pkg/config"
	"gails/app/pkg/redis"
	"gails/app/routes"

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
	r.HTMLRender = ginview.Default()

	r.Run(":" + config.ServerCfg.Port)
}
