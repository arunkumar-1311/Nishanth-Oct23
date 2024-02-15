package router

import (
	"to-do/adaptor"
	"to-do/handler"
	"to-do/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	MW "to-do/middleware"
)

func Router(db adaptor.Database, service service.ServiceMethods) error {

	var handlers handler.EndPoint
	handlers.DB, handlers.Service, handlers.MW = db, service, MW.AcquireMiddleware(db)

	api := handler.AcqurieAPI(handlers)

	routes := echo.New()
	routes.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		},
	))

	routes.POST("/signup", api.SignIn)
	routes.POST("/login", api.Login)

	profile := routes.Group("/profile", handlers.MW.Authentication)
	profile.GET("", api.GetProfile)
	profile.PATCH("", api.UpdateProfile)
	profile.DELETE("", api.DeleteProfile)

	task := routes.Group("/task", handlers.MW.Authentication)
	task.POST("", api.AddTask)
	task.PATCH("/:id", api.UpdateTask)
	task.DELETE("/:id", api.DeleteTask)
	task.PATCH("/status", api.UpdateAllTaskStatus)

	if err := routes.Start(":8000"); err != nil {
		return err
	}
	return nil
}
