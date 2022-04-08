package service

import (
	"ginValid/dto/user"
	"ginValid/models"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*user.Read, error)
	GetUserById(*string) (*user.Read, error)
}
