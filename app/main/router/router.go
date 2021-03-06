package router

import (
	"fmt"

	"github.com/eifandevs/main/controllers"
	"github.com/eifandevs/main/interceptor"
	"github.com/eifandevs/main/mock"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func dumpHandler(c echo.Context, reqBody, resBody []byte) {
	output := fmt.Sprintf("%#v", c.Request().Header)

	fmt.Printf("Request Header: %v\n", output)
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.BodyDump(dumpHandler))

	api := e.Group("/api")
	{
		api.POST("/login", controllers.LoginHandler(), interceptor.BasicAuth())
		api.GET("/favorite", controllers.GetFavoriteHandler())
		api.POST("/favorite", controllers.PostFavoriteHandler())
		api.DELETE("/favorite", controllers.DeleteFavoriteHandler())
	}

	mockApi := e.Group("/mock")
	{
		mockApi.GET("/test", mock.GetTest())
		mockApi.POST("/test", mock.PostTest())
	}

	e.Static("/static", "./static")
	e.GET("/digest", interceptor.DigestAuthenticate())

	return e
}
