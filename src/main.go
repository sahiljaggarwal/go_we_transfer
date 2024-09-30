package main

import (
	"log"
	"super_crud/src/common"
	"super_crud/src/config"
	"super_crud/src/cron"
	"super_crud/src/database"
	"super_crud/src/routes"

	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	app := fiber.New()
	database.ConnectDB()

	// middlewares
	// Rate Limiter
	app.Use(limiter.New(limiter.Config{
    Max:            20,
    Expiration:     30 * time.Second,
    LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Use(cors.New(cors.Config{
	AllowOrigins:     "*",
	AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	AllowHeaders:     "Origin,Content-Type,X-Auth-Token",
	ExposeHeaders:    "Content-Length",
	MaxAge:           300,
	}))
	app.Use(helmet.New())
	app.Use(cache.New(cache.Config{
	Next: func(c *fiber.Ctx) bool {
		return c.Method() != fiber.MethodGet
	},
	Expiration: 30 * time.Second, 
	}))
	app.Use(logger.New())
	
	// Routes
	routes.Routes(app, database.DB)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "WeTransfer Metrics Page"}))
	app.Get("/", func(c *fiber.Ctx) error {
		response := common.Response{
			Status: "SUCCESS",
			Message: "Well Come To We Transfer Platform",
			Data: nil,
		}
		return c.Status(200).JSON(response)
	})

	//cron job
	cron.StartCronJob(database.DB)

	// server
	PORT := config.PORT
	if PORT == ""{
		PORT = "3000"
	}
	log.Println("Server is running on port 3000")
	log.Fatal(app.Listen(":"+PORT))

}