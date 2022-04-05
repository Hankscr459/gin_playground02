package main

import (
	validService "ginValid/extension"
	valid "ginValid/middleware"
	models "ginValid/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	route := gin.Default()
	validService.Valid()
	route.GET("/time", valid.DateValidator(), getTime)
	route.POST("/user/create", valid.SignupValidator(), create)
	route.Run(":8080")
}

func create(c *gin.Context) {
	var user models.User
	c.ShouldBindBodyWith(&user, binding.JSON)
	user.Age = 22
	c.JSON(http.StatusBadRequest, user)
}

func getTime(c *gin.Context) {
	date, _ := c.Get("date")
	c.JSON(http.StatusOK, gin.H{"message": date})
}
