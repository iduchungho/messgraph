package message

type MessageController struct {
	Service *MessageService
}

func NewUserController(service *MessageService) *MessageController {
	return &MessageController{
		Service: service,
	}
}