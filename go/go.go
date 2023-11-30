package main

import (
	"os"

	"github.com/linn221/go-blog/models"
	"github.com/linn221/go-blog/seeders"
)

func main() {
	models.ConnectDB()
	command := os.Args[1]
	switch command {
	case "seed":
		seeders.SeedCategories()
	}
}
