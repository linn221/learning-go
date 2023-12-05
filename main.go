package main

import (
	"os"

	"github.com/linn221/go-blog/models"
	"github.com/linn221/go-blog/routes"
	"github.com/linn221/go-blog/seeders"
)

func main() {
	models.ConnectDB()
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "migrate":
			models.FreshDB()
		case "seed":
			seeders.Run()
		case "migrate:seed":
			models.FreshDB()
			seeders.Run()
		}
		return
	}
	routes.StartServer()
}
