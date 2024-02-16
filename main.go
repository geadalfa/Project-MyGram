package main

import (
	"Project-MyGram/database"
	"Project-MyGram/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
