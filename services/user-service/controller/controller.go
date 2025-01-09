package controller

import (
	"user-service/repository"
	"user-service/service"
)

type Controller struct {
	UserService service.UserServiceInterface
	Repository  *repository.UserServiceRepository
}

func newController(
	userService service.UserServiceInterface,
	repository *repository.UserServiceRepository) *Controller {
	return &Controller{
		UserService: userService,
		Repository:  repository,
	}
}

func NewUserController(
	userService service.UserServiceInterface,
	repository *repository.UserServiceRepository,
) *Controller {
	return newController(userService, repository)
}
