package main

import (
	"Grocery-Shopping-User-Module/src/app/api"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)
func main() {
	fmt.Println("Starting User Module.... ")
	config := api.GetConfig()

	fmt.Println("Configuration Values :" +os.Getenv(os.Getenv("DATABASE_URL"))+" : "+ os.Getenv("DATABASE_HOST")+" :"+os.Getenv("DATABASE_USERNAME")+" : "+os.Getenv("DATABASE_PASSWORD"))
	app := &api.App{}
	app.Initialize(config)

	port := os.Getenv("PORT")
	fmt.Println("Port No. is :"+port)
	app.Run(":"+port)

	fmt.Println("Started User Module.... ")
}