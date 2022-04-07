package controller

import (
	"fmt"
	models "ginValid/models"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Create(c *gin.Context) {
	var user models.User
	c.ShouldBindBodyWith(&user, binding.JSON)
	c.JSON(http.StatusBadRequest, user)
}

func GetTime(c *gin.Context) {
	value, _ := c.Get("date")
	date := reflect.ValueOf(value)
	createTime := date.FieldByName("CreateTime")
	fmt.Println(createTime)
	c.JSON(http.StatusOK, gin.H{"message": date})
}
