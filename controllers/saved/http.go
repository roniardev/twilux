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
	userId := middlewares.GetUser(c)
	saved, err := controller.usecase.GetAll(userId.Username, ctx)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusNonAuthoritativeInfo, "You are not authorized", err)
	}

	return controllers.SuccessResponse(c, []string{"Get all saved data succed."}, response.ToListDomain(saved))
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
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Please provide valid snippet_id", err)
	}
	return controllers.SuccessResponse(c, []string{"Create saved data succed."}, response.FromDomain(saved))
}

// Delete saved controller

func (controller *SavedController) Delete(c echo.Context) error {
	var savedDelete request.SavedDelete
	ctx := c.Request().Context()
	userId := middlewares.GetUser(c)
	savedDelete.Username = userId.Username

	if err := c.Bind(&savedDelete); err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	res, err := controller.usecase.Delete(*savedDelete.ToDeleteDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, []string{"Saved deleted."}, response.FromDeleteDomain(res))
}
