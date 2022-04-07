package main

import (
	controller "ginValid/controller"
	"os"

	db "ginValid/extension"
	vaildService "ginValid/extension"
	valid "ginValid/middleware"
	models "ginValid/models"

	_ "github.com/joho/godotenv/autoload"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	route := gin.Default()
	vaildService.Valid()
	if db.CheckConnection() == 0 {
		log.Fatal("Sin connect a la DB")
		return
	}

	route.GET("/time", valid.DateValidator(), getTime)
	route.POST("/user/create", valid.SignupValidator(), create)
	route.GET("/export/excel", controller.Employee)
	route.Run(os.Getenv("PORT"))
}

func create(c *gin.Context) {
	var user models.User
	c.ShouldBindBodyWith(&user, binding.JSON)
	c.JSON(http.StatusBadRequest, user)
}

func getTime(c *gin.Context) {
	date, _ := c.Get("date")
	c.JSON(http.StatusOK, gin.H{"message": date})
}
