package main

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/middlewares"
	"acp9-redy-gigih/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":6969"))
}
