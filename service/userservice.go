package service

import (
	"ginValid/models"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
}
