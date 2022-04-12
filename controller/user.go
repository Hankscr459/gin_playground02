package controller

import (
	"fmt"
	"ginValid/dto/user"
	implement "ginValid/implement"
	valid "ginValid/middleware"
	models "ginValid/models"
	"ginValid/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/copier"
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
	var body models.User
	var read user.Read
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	body.Password, _ = implement.EncriptPassword(body.Password)
	create, err := uc.UserService.CreateUser(&body)
	if dbErr := implement.ValidDbError(err, ctx); dbErr {
		return
	}
	copier.Copy(read, &create)
	tk, err := implement.GenerJWT(read)

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "token": tk})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil && err.Error() == "mongo: no documents in result" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "此會員不存在"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
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
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	var update user.Update
	if err := ctx.ShouldBindBodyWith(&update, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&update, &userId)
	if err != nil && err.Error() == "mongo: no documents in result" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "此會員不存在"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success to update"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", valid.SignupValidator(), uc.CreateUser)
	userroute.PUT("/:userId", uc.UpdateUser)
	userroute.GET("/:userId", uc.GetUserById)
	userroute.GET("/get/:name", uc.GetUser)
}
