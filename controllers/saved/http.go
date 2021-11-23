package controllers

import (
	"net/http"
	"twilux/business/saved"
	"twilux/controllers"
	"twilux/controllers/saved/request"
	"twilux/controllers/saved/response"
	"twilux/middlewares"

	"github.com/labstack/echo/v4"
)

type SavedController struct {
	usecase saved.SavedUsecaseInterface
}

func NewSavedController(uc saved.SavedUsecaseInterface) *SavedController {
	return &SavedController{
		usecase: uc,
	}
}

func (controller *SavedController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	saved, err := controller.usecase.GetAll(ctx)
	if err != nil {
		return controllers.SuccessResponse(c, response.ToListDomain(saved))
	}

	return controllers.SuccessResponse(c, response.ToListDomain(saved))
}

func (controller *SavedController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var savedCreate request.SavedCreate
	userId := middlewares.GetUser(c)
	savedCreate.Username = userId.Username

	errs := c.Bind(&savedCreate)
	if errs != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", errs)
	}

	saved, err := controller.usecase.Create(*savedCreate.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(saved))
}

// Delete saved controller

func (controller *SavedController) Delete(c echo.Context) error {
	snippReq := request.SavedDelete{}
	ctx := c.Request().Context()

	userId := middlewares.GetUser(c)

	if err := c.Bind(&snippReq); err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	savedDomain := snippReq.ToDeleteDomain()
	savedDomain.Username = userId.Username

	_, err := controller.usecase.Delete(*savedDomain, ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(*savedDomain))
}
