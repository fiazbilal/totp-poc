package main

import (
	"totp-poc/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/register", handler.RegisterHandler)
	app.Post("/verify", handler.VerifyHandler)

	app.Listen(":3000")
}
