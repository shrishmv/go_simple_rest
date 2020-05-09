package server

import (
	"go_simple_rest/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter for routing requests
func NewRouter() *echo.Echo {

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{echo.OPTIONS, echo.POST, echo.GET, echo.PUT, echo.DELETE, echo.PATCH},
	}))

	// Endpoints for healthcheck
	router.GET("/status", controllers.GetStatus)

	restService := router.Group("/")
	var version *echo.Group

	version = restService.Group("v1")

	//this group is for test APIS.
	readApis := version.Group("/test")
	readApis.GET("/myget/", controllers.TestGet)

	return router
}
