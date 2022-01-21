package main

import "github.com/gofiber/fiber/v2"

func main() {
	app:=fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return fiber.NewError(500,"Don't panic this is a test error from server!!")
	})

	app.Listen(":8000")
}
