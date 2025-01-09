package controller

type Controller struct {
}

func newController() *Controller {
	return &Controller{}
}

func NewGatewayController() *Controller {
	return newController()
}
