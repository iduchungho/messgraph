package user

type Controller struct {
	Service IService
}

func NewUserController(service IService) *Controller {
	return &Controller{
		Service: service,
	}
}
