package controller

type Controller struct {
}

func newController() *Controller {
	return &Controller{}
}

func NewUserController() *Controller {
	return newController()
}
