package router

import (
	"blog_post/adaptor"
	handlers "blog_post/handler"
	"blog_post/middleware"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes(router *fiber.App, db adaptor.Database) *fiber.App {
	defer fmt.Println("Starting the server....")

	var handler handlers.Handler
	handler.Method = db
	// Helps to register the user
	router.Post("/register", handler.Register)

	// Helps to login the existing user
	router.Post("/login", handler.Authentication)

	// Helps to manipulate with list
	router.Post("/admin/category", middleware.Authorization(), handler.CreateCategory)
	router.Get("/category", handler.ReadAllCategories)
	router.Patch("/admin/category/:id", middleware.Authorization(), handler.UpdateCategory)  // This id indicates category ID
	router.Delete("/admin/category/:id", middleware.Authorization(), handler.DeleteCategory) // This id indicates category ID

	// Helps to manipulate with posts
	router.Post("/admin/post", middleware.Authorization(), handler.CreatePost)
	router.Get("/posts", handler.ReadAllPosts)
	router.Patch("/admin/post/:id", middleware.Authorization(), handler.UpdatePost)  // This id indicates Post ID
	router.Delete("/admin/post/:id", middleware.Authorization(), handler.DeletePost) // This id indicates Post ID

	// Helps to manipulate with the comments
	router.Get("/admin/comment/user/:id", middleware.Authorization(), handler.ReadCommentByUser) // This id indicates User ID
	router.Get("/comment/post/:id", handler.ReadCommentByPost)                                   // This id indicates Post ID
	router.Get("/comment/:id", handler.ReadCommentByID)                                          // This id indicates comment ID
	router.Post("/comment/:id", middleware.Authorization(), handler.AddComment)                  // This id indicates Post ID
	router.Patch("/comment/:id", middleware.Authorization(), handler.UpdateComment)              // This id indicates comment ID
	router.Delete("/comment/:id", middleware.Authorization(), handler.DeleteComment)             // This id indicates comment ID

	// Helps to add filter the post displaying
	router.Post("/posts/date", handler.DateFilter)
	router.Get("/posts/:id", handler.CategoryFilter) // This id indicates category ID

	// Overview of the profile
	router.Get("/admin/overview", middleware.Authorization(), handler.Overview)
	return router
}
