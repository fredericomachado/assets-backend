package main

import (
	"assets-backend/database"
	"assets-backend/routes"
)

func main() {
	database.Connect()
	routes.HandleRequest()
}
