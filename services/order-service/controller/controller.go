package controller

type Controller struct {
}

func newController() *Controller {
	return &Controller{}
}

func NewOrderController() *Controller {
	return newController()
}
