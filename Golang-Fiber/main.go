package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	if fiber.IsChild() {
		fmt.Println("I'm child")
	} else {
		fmt.Println("I'm parents")
	}

	app.Use("/api", func(ctx *fiber.Ctx) error {
		fmt.Println("I'm middleware before")
		err := ctx.Next()
		fmt.Println("I'm middleware after")
		return err
	})

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
