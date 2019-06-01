package main

import (
	"Grocery-Shopping-User-Module/src/app/api"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)
func main() {
	config := api.GetConfig()

	app := &api.App{}
	app.Initialize(config)
	app.Run(os.Getenv("PORT"))
}