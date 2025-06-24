package main

import (
	"log"

	"github.com/viniciuscg/gin-api/config"
	"github.com/viniciuscg/gin-api/db"
	"github.com/viniciuscg/gin-api/routes"
)

func main() {
	config.LoadEnv()
	dbConn := db.Connect()
	router := routes.SetupRouter(dbConn)

	log.Fatal(router.Run(":8080"))
}
