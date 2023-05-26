package message

type Controller struct {
	Service *Service
}

func NewUserController(service *Service) *Controller {
	return &Controller{
		Service: service,
	}
}
