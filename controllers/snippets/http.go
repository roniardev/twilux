package controllers

import (
	"net/http"
	"twilux/business/snippets"
	"twilux/controllers"
	"twilux/controllers/snippets/request"
	"twilux/controllers/snippets/response"
	"twilux/middlewares"

	"github.com/labstack/echo/v4"
)

type SnippetController struct {
	usecase snippets.SnippetUsecaseInterface
}

func NewSnippetController(uc snippets.SnippetUsecaseInterface) *SnippetController {
	return &SnippetController{
		usecase: uc,
	}
}

func (controller *SnippetController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	snippet, err := controller.usecase.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}

	return controllers.SuccessResponse(c, []string{"Get all snippets succed."}, response.ToListDomain(snippet))
}

func (controller *SnippetController) GetById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	snippet, err := controller.usecase.GetById(id, ctx)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Snippet with this id is not found", err)
	}

	return controllers.SuccessResponse(c, []string{"Get snippet succed."}, response.FromDomain(snippet))
}

func (controller *SnippetController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var snippetCreate request.SnippetCreate
	userId := middlewares.GetUser(c)
	snippetCreate.Username = userId.Username

	errs := c.Bind(&snippetCreate)
	if errs != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", errs)
	}
	snippet, err := controller.usecase.Create(*snippetCreate.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, []string{"A new snippet created."}, response.FromCreateDomain(snippet))
}

// Update snippet controller
func (controller *SnippetController) Update(c echo.Context) error {
	var snippetUpdate request.SnippetUpdate
	ctx := c.Request().Context()
	id := c.Param("id")
	userId := middlewares.GetUser(c)
	snippetUpdate.Id = id
	snippetUpdate.Username = userId.Username

	if err := c.Bind(&snippetUpdate); err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	res, err := controller.usecase.Update(*snippetUpdate.ToUpdateDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, []string{"Data updated."}, response.FromUpdateDomain(res))
}

// Delete snippet controller

func (controller *SnippetController) Delete(c echo.Context) error {
	var snippetDelete request.SnippetDelete
	ctx := c.Request().Context()
	id := c.Param("id")
	userId := middlewares.GetUser(c)
	snippetDelete.Id = id
	snippetDelete.Username = userId.Username

	if err := c.Bind(&snippetDelete); err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	res, err := controller.usecase.Delete(*snippetDelete.ToDeleteDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, []string{"Snippet deleted."}, response.FromDeleteDomain(res))
}
