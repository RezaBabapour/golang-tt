package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app:=fiber.New(fiber.Config{Prefork: true})
	if !fiber.IsChild(){
		fmt.Println("parent")
	} else {
		fmt.Println("child")
	}
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("this is home")
	})

	app.Listen(":8000")
}