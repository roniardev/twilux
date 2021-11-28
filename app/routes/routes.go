package routes

import (
	commentController "twilux/controllers/comments"
	savedController "twilux/controllers/saved"
	snippetController "twilux/controllers/snippets"
	userController "twilux/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	JwtConfig         middleware.JWTConfig
	UserController    userController.UserController
	SnippetController snippetController.SnippetController
	SavedController   savedController.SavedController
	CommentController commentController.CommentController
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	c.Use(middleware.CORS())
	c.Use(middleware.Logger())

	users := c.Group("/user")

	users.POST("/login", controller.UserController.Login)
	users.POST("/register", controller.UserController.Register)

	snippet := c.Group("/snippets")
	snippet.GET("/", controller.SnippetController.GetAll)
	snippet.GET("/:id", controller.SnippetController.GetById)
	snippet.POST("/", controller.SnippetController.Create, middleware.JWTWithConfig(controller.JwtConfig))
	snippet.PUT("/:id", controller.SnippetController.Update, middleware.JWTWithConfig(controller.JwtConfig))
	snippet.DELETE("/:id", controller.SnippetController.Delete, middleware.JWTWithConfig(controller.JwtConfig))

	saved := c.Group("/saved")
	saved.GET("/", controller.SavedController.GetAll, middleware.JWTWithConfig(controller.JwtConfig))
	saved.POST("/", controller.SavedController.Create, middleware.JWTWithConfig(controller.JwtConfig))
	saved.DELETE("/", controller.SavedController.Delete, middleware.JWTWithConfig(controller.JwtConfig))

	comment := c.Group("/comments")
	comment.GET("/:id", controller.CommentController.GetAll)
	comment.POST("/:id", controller.CommentController.Create, middleware.JWTWithConfig(controller.JwtConfig))
	comment.PUT("/:id", controller.CommentController.Update, middleware.JWTWithConfig(controller.JwtConfig))
	comment.DELETE("/:id", controller.CommentController.Delete, middleware.JWTWithConfig(controller.JwtConfig))
}
