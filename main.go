package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Prefork:      true,
	})

	app.Use("/api", func(ctx *fiber.Ctx) error {
		fmt.Println("Middleware before processing request")
		err := ctx.Next()
		fmt.Println("Middleware after processing request")
		return err
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	if fiber.IsChild() {
		fmt.Println("Child process")
	} else {
		fmt.Println("Parent process")
	}

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
