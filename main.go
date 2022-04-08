package main

import (
	"ginValid/controller"
	db "ginValid/extension"
	vaildService "ginValid/extension"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	server         *gin.Engine
	usercontroller controller.UserController
)

func init() {
	usercontroller = db.ConnectDb()
	vaildService.Valid()
	server = gin.Default()
}

func main() {
	defer db.Disconnect()
	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(os.Getenv("PORT")))
}

// route.GET("/time", valid.DateValidator(), controller.GetTime)
// route.POST("/user/create", valid.SignupValidator(), controller.Create)
// route.GET("/export/excel", controller.Employee)
