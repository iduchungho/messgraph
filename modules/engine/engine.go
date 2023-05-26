package engine

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

type Engine struct {
	router *fiber.App
}

func NewEngine() *Engine {
	return &Engine{
		router: fiber.New(),
	}
}

func (eng *Engine) Prepare() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (eng *Engine) Run() {
	// define cors middleware
	eng.router.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language," +
			"Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// logger actions to server
	eng.router.Use(logger.New())
	//app.r.Use(middleware.Redirect)

	// routing services application

	host := os.Getenv("PORT")
	if host != "" {
		err := eng.router.Listen("localhost:" + host)
		if err != nil {
			panic("Can't run fiber engine")
		}
	} else {
		err := eng.router.Listen("0.0.0.0:8000")
		if err != nil {
			panic("Can't run fiber engine")
		}
	}
}
