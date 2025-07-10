package main

import (
	"badminton-app/database"
	"badminton-app/router"
)

func main() {
	database.InitDB()
	r := router.SetupRouter()
	r.Run(":8080")
}
