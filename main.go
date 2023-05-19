package main

import (
	"DTSGolang/FinalProject/database"
	"DTSGolang/FinalProject/routers"
)

func main() {
	database.StartDB()
	defer database.CloseDB()

	r := routers.StartApp()
	var PORT = ":8000"

	r.Run(PORT)
}
