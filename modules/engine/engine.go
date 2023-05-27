package engine

import (
	"context"
	"messgraph.com/m/database"
	"messgraph.com/m/modules/entities/user"
	"messgraph.com/m/router"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Engine struct {
	router   *router.Router
	handle   *fiber.App
	database *database.Database
}

func New() *Engine {
	app := fiber.New()
	return &Engine{
		router:   router.New(app),
		handle:   app,
		database: database.CreateConnectDatabase(),
	}
}

func (eng *Engine) Run() {
	// define cors middleware
	eng.handle.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language," +
			"Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	defer eng.database.Close(context.Background())
	// logger actions to server
	eng.handle.Use(logger.New())

	//create controllers
	var userController = user.NewUserController(
		user.NewUserService(user.NewUserRepo(eng.database)))
	// routing services application
	eng.router.UserRouter(userController)

	host := os.Getenv("PORT")
	if host != "" {
		err := eng.handle.Listen("localhost:" + host)
		if err != nil {
			panic("Can't run fiber engine")
		}
	} else {
		err := eng.handle.Listen("0.0.0.0:8000")
		if err != nil {
			panic("Can't run fiber engine")
		}
	}
}
