package routes

import (
	c "day4/tugas-crud-dinamis/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	// Login
	e.POST("v1/login", c.LoginController)

	// User
	e.GET("v1/users", c.GetUserController)
	e.GET("v1/users/:id", c.GetUserControllerById)
	e.POST("v1/users", c.CreateUserController)
	e.DELETE("v1/users/:id", c.DeleteUserController)
	e.PUT("v1/users/:id", c.UpdateUserController)

	// Book
	e.GET("v1/books", c.GetBookController)
	e.GET("v1/books/:id", c.GetBookControllerById)
	e.POST("v1/books", c.CreateBookController)
	e.DELETE("v1/books/:id", c.DeleteBookController)
	e.PUT("v1/books/:id", c.UpdateBookController)

	//JWT Group
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	// User jwt
	// jwtAuth.GET("v1/users", c.GetUserController)
	// jwtAuth.PUT("v1/users/:id", c.UpdateUserController)
	// jwtAuth.DELETE("v1/users/:id", c.DeleteUserController)
	// jwtAuth.GET("v1/users/:id", c.GetUserControllerById)
	// // Book jwt
	// jwtAuth.PUT("v1/books/:id", c.UpdateBookController)
	// jwtAuth.DELETE("v1/books/:id", c.DeleteBookController)
	// jwtAuth.POST("v1/books", c.CreateBookController)

	// implement middleware with group rounting
	// eAuth := e.Group("")
	// eAuth.Use(echoMid.BasicAuth(m.BasicAuthDb))
	// eAuth.DELETE("v1/users/:id", c.GetUserControllerById)

	return e

}
