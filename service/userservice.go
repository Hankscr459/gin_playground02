package service

import (
	"ginValid/dto/user"
	"ginValid/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(*models.User) (*mongo.InsertOneResult, error)
	GetUser(*string) (*user.Read, error)
	GetUserByEmail(*string) (*user.Read, error)
	GetUserById(*string) (*user.Read, error)
	UpdateUser(*user.Update, *string) error
}
