package main

import (
	"Grocery-Shopping-User-Module/src/app/api"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)
func main() {
	fmt.Print("Starting User Module.... ")
	config := api.GetConfig()

	fmt.Print("Configuration Values :" +config.DB.Host)
	app := &api.App{}
	app.Initialize(config)
	fmt.Print("Port No. is :"+os.Getenv("PORT"))
	app.Run(":3000")

	fmt.Print("Started User Module.... ")
}