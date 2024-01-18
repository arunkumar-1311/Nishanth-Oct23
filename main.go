package main

import (
	routing "blog_post/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	router := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "Blog Post v1.0.1",
	})

	routes := routing.Routes(router)
	if err := routes.Listen(":8000"); err != nil {
		fmt.Println("Can't listen to the server ", err)
		return
	}
}
