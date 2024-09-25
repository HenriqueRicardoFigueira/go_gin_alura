package main

import (
	"go_gin_alura/database"
	"go_gin_alura/routes"
)

func main() {
	database.ConnectDb()
	routes.HandleRequests()
}
