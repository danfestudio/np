package server

import "github.com/gofiber/fiber/v2"

func HTTPServer() {

	app := fiber.New()
	app.Static("/", "./public")
	app.Listen(":3000")

}
