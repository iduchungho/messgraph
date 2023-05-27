package router

import (
	"github.com/gofiber/fiber/v2"
	"messgraph.com/m/modules/entities/user"
)

type Router struct {
	handle *fiber.App
}

func New(app *fiber.App) *Router {
	return &Router{
		handle: app,
	}
}

func (r *Router) UserRouter(controller user.IController) {
	r.handle.Post("/api/signup", controller.Register)
	r.handle.Post("/api/login", controller.Login)
}
