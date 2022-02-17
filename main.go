package main

import (
	"fmt"
	"log"
	"worker-transaction/appconfig"
	"worker-transaction/controllers"
	"worker-transaction/services/databaseservice"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	appconfig.LoadInitConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	store := databaseservice.GetDatabase()

	fmt.Println(store.Db.Ping())

	defer func() {
		if databaseservice.GetDatabase() != nil {
			databaseservice.GetDatabase().Db.Close()
		}
	}()

	r.POST("/withdrawal", controllers.Withdrawal)
	r.Run()
}
