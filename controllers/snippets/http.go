package controllers

import (
	"net/http"
	"twilux/business/snippets"
	"twilux/controllers"
	"twilux/controllers/snippets/request"
	"twilux/controllers/snippets/response"

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
		return controllers.SuccessResponse(c, response.ToListDomain(snippet))
	}

	return controllers.SuccessResponse(c, response.ToListDomain(snippet))
}

func (controller *SnippetController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var snippetCreate request.SnippetCreate
	errs := c.Bind(&snippetCreate)
	if errs != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", errs)
	}
	snippet, err := controller.usecase.Create(*snippetCreate.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(snippet))
}

// Update snippet controller
func (controller *SnippetController) Update(c echo.Context) error {
	snippReq := request.SnippetUpdate{}
	ctx := c.Request().Context()
	id := c.Param("id")

	if err := c.Bind(&snippReq); err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	snippetDomain := snippReq.ToUpdateDomain()
	snippetDomain.Id = id

	_, err := controller.usecase.Update(*snippetDomain, ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(*snippetDomain))
}
