package main

import "gails/app/routes"

func main() {
	r := routes.InitRouter()
	r.Run()
}
