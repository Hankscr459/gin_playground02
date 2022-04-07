package main

import (
	"ginValid/controller"
	"os"

	db "ginValid/extension"
	vaildService "ginValid/extension"
	valid "ginValid/middleware"

	_ "github.com/joho/godotenv/autoload"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	vaildService.Valid()
	if db.CheckConnection() == 0 {
		log.Fatal("Fiail connect to DB")
		return
	}

	route.GET("/time", valid.DateValidator(), controller.GetTime)
	route.POST("/user/create", valid.SignupValidator(), controller.Create)
	route.GET("/export/excel", controller.Employee)
	route.Run(os.Getenv("PORT"))
}
