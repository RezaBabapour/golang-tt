package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	//app:=fiber.New(fiber.Config{Prefork: true})
	app:=fiber.New()
	if !fiber.IsChild(){
		fmt.Println("parent")
	} else {
		fmt.Println("child")
	}
	app.Get("/", func(ctx *fiber.Ctx) error {
		log.Println("home page called")
		return ctx.SendString("this is home")
	})

	app.Listen(":8000")
}