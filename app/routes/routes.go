package routes

import (
	userController "twilux/controllers/users"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController userController.UserController
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	users := c.Group("/user")
	users.POST("/login", controller.UserController.Login)
	users.POST("/register", controller.UserController.Register)
}
