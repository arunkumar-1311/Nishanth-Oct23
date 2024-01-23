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
	router.Post("/login", handler.Authentication)

	// Helps to manipulate with list
	router.Post("/admin/category", handler.Authorization(), admin.CreateCategory)
	router.Get("/category", admin.ReadAllCategories)
	router.Patch("/admin/category/:id", handler.Authorization(), admin.UpdateCategory)  // This id indicates category ID
	router.Delete("/admin/category/:id", handler.Authorization(), admin.DeleteCategory) // This id indicates category ID

	// Helps to manipulate with posts
	router.Post("/admin/post", handler.Authorization(), admin.CreatePost)
	router.Get("/posts", admin.ReadAllPosts)
	router.Patch("/admin/post/:id", handler.Authorization(), admin.UpdatePost)  // This id indicates Post ID
	router.Delete("/admin/post/:id", handler.Authorization(), admin.DeletePost) // This id indicates Post ID

	// Helps to manipulate with the comments
	router.Get("/admin/comment/user/:id", handler.Authorization(), handler.ReadCommentByUser) // This id indicates User ID
	router.Get("/comment/post/:id", handler.ReadCommentByPost)                                // This id indicates Post ID
	router.Get("/comment/:id", handler.ReadCommentByID)                                       // This id indicates comment ID
	router.Post("/comment/:id", handler.Authorization(), handler.AddComment)                  // This id indicates Post ID
	router.Patch("/comment/:id", handler.Authorization(), handler.UpdateComment)              // This id indicates comment ID
	router.Delete("/comment/:id", handler.Authorization(), handler.DeleteComment)             // This id indicates comment ID

	// Helps to add filter the post displaying
	router.Post("/posts/date", handler.DateFilter)
	router.Get("/posts/:id", handler.CategoryFilter) // This id indicates category ID

	// Overview of the profile
	router.Get("/admin/overview", handler.Authorization(), admin.Overview)
	return router
}
