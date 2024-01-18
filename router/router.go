package router

import (
	"blog_post/handler"
	"blog_post/handler/admin"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes(router *fiber.App) *fiber.App {
	defer fmt.Println("Starting the server....")

	// Helps to register the user
	router.Post("/register", handler.Register)

	// Helps to login the existing user
	router.Post("/login", handler.Authentication(), admin.ReadAllPosts)

	// Helps to manipulate with list
	router.Post("/category", handler.Authorization(), admin.CreateCategory)
	router.Get("/category", admin.ReadAllCategories)
	router.Patch("/category/:id", handler.Authorization(), admin.UpdateCategory)  // This id indicates category ID
	router.Delete("/category/:id", handler.Authorization(), admin.DeleteCategory) // This id indicates category ID

	// Helps to manipulate with posts
	router.Post("/post", handler.Authorization(), admin.CreatePost)
	router.Get("/posts", admin.ReadAllPosts)
	router.Patch("/post/:id", handler.Authorization(), admin.UpdatePost)  // This id indicates Post ID
	router.Delete("/post/:id", handler.Authorization(), admin.DeletePost) // This id indicates Post ID

	// Helps to add filter the post displaying

	// Helps to manipulate with the comments
	router.Get("/comments", handler.Authorization(), handler.ReadMyComment)
	router.Post("/comment/:id", handler.Authorization(), handler.AddComment)      // This id indicates Post ID
	router.Patch("/comment/:id", handler.Authorization(), handler.UpdateComment)  // This id indicates comment ID
	router.Delete("/comment/:id", handler.Authorization(), handler.DeleteComment) // This id indicates comment ID
	return router
}
