package main

import (
	"gails/app/models"
	"gails/app/pkg/config"
	"gails/app/routes"
)

func init() {
	config.Init()
	models.Init()
}

func main() {
	r := routes.InitRouter()
	r.Run()
}
