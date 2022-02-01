package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app:= fiber.New()

	app.Get("/",home)
	app.Listen(":8000")

}

func home(c *fiber.Ctx) error {
	return c.SendString("salam")
}