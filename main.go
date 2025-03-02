package main

import (
	"fralla/config"
	"fralla/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.LoadEnv()

	app := fiber.New(fiber.Config{
		Network: fiber.NetworkTCP,
	})

	app.Use(logger.New())
	app.Static("/assets", os.Getenv("APP_PATH")+"assets")

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Set("Surrogate-Control", "no-store")
		return c.Next()
	})

	routes.SetRoutes(app)

	go func() {
		if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
			log.Fatal(err)
		}
	}()

	if err := app.ListenTLS("0.0.0.0:"+os.Getenv("APP_PORT_SSL"),
		os.Getenv("APP_PATH")+"ssl/server.crt",
		os.Getenv("APP_PATH")+"ssl/server.key"); err != nil {
		log.Fatal(err)
	}

}
