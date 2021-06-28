package main

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":6969"))
}
