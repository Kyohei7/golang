package main

import (
	"latihan-simple-rest-api-middleware/database"
	"latihan-simple-rest-api-middleware/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
