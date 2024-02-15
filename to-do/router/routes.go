package router

import (
	"to-do/adaptor"
	"to-do/handler"
	"to-do/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router(db adaptor.Database, service service.ServiceMethods) error {

	var handlers handler.EndPoint
	handlers.DB, handlers.Service = db, service

	api := handler.AcqurieAPI(handlers)

	routes := echo.New()
	routes.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		},
	))
	
	routes.POST("/signup", api.SignIn)

	if err := routes.Start(":8000"); err != nil {
		return err
	}
	return nil
}
