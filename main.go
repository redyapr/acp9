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
	port := config.Env("APP_PORT")
	if port == "" {
		e.Logger.Fatal("APP_PORT must be set")
	}
	e.Logger.Fatal(e.Start(":" + port))
}
