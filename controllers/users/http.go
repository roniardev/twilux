package controllers

import (
	"fmt"
	"net/http"
	"twilux/business/users"
	"twilux/controllers"
	"twilux/controllers/users/request"
	"twilux/controllers/users/response"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UserUsecaseInterface
}

func NewUserController(uc users.UserUsecaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var userLogin request.UserLogin
	errs := c.Bind(&userLogin)
	if errs != nil {
		fmt.Println(errs)
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", errs)
	}
	user, err := controller.usecase.Login(*userLogin.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Email or password is incorrect.", err)
	}
	return controllers.SuccessResponse(c, []string{"Logged in."}, response.FromLogDomain(user))
}

// SignUp controller
func (controller *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()
	var userRegister request.UserRegister
	errs := c.Bind(&userRegister)
	if errs != nil {
		fmt.Println(errs)
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", errs)
	}
	user, err := controller.usecase.Register(*userRegister.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "Username or email has been registered.", err)
	}
	return controllers.SuccessResponse(c, []string{"Register succed."}, response.FromRegDomain(user))
}
