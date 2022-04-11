package controller

import (
	"fmt"
	implement "ginValid/implement"
	valid "ginValid/middleware"
	models "ginValid/models"
	"ginValid/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct {
	UserService service.UserService
}

func New(userservice service.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user.Password, _ = implement.EncriptPassword(user.Password)
	err := uc.UserService.CreateUser(&user)
	if dbErr := implement.ValidDbError(err, ctx); dbErr {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil && err.Error() == "mongo: no documents in result" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "此會員不存在"})
		fmt.Println("runnning3")
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		fmt.Println("runnning2")
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user, err := uc.UserService.GetUserById(&userId)
	fmt.Println(user)
	if err != nil && err.Error() == "mongo: no documents in result" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "此會員不存在"})
		fmt.Println("runnning3")
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		fmt.Println("runnning2")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", valid.SignupValidator(), uc.CreateUser)
	userroute.GET("/:userId", uc.GetUserById)
	userroute.GET("/get/:name", uc.GetUser)
}
