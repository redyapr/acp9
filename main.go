package main

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/routes"
)


func main() {
	config.InitDB()
	config.InitialMigration()
	e := routes.New()

	e.Start(":6969")
}
