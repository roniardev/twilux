package controllers

import (
	"fmt"
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

func (controller *SnippetController) GetAllSnippets(c echo.Context) error {
	return controllers.SuccessResponse(c, response.SnippetResponse{})
}

func (controller *SnippetController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var snippetCreate request.SnippetCreate
	errs := c.Bind(&snippetCreate)
	if errs != nil {
		fmt.Println(errs)
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", errs)
	}
	snippet, err := controller.usecase.Create(*snippetCreate.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(snippet))
}
