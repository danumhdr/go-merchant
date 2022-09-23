package main

import (
	"go-merchant/database"
	"go-merchant/router"
)

func main() {
	database.Connect()
	setup := router.SetupRouter()
	setup.Run(":9000")
}
