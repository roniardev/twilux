package controllers

import (
	"net/http"
	"twilux/business/comments"
	"twilux/controllers"
	"twilux/controllers/comments/request"
	"twilux/controllers/comments/response"
	"twilux/middlewares"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	usecase comments.CommentUsecaseInterface
}

func NewCommentController(uc comments.CommentUsecaseInterface) *CommentController {
	return &CommentController{
		usecase: uc,
	}
}

func (controller *CommentController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	saved, err := controller.usecase.GetAll(id, ctx)

	if err != nil {
		return controllers.SuccessResponse(c, response.ToListDomain(saved))
	}

	return controllers.SuccessResponse(c, response.ToListDomain(saved))
}

func (controller *CommentController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	var commentCreate request.CommentCreate
	userId := middlewares.GetUser(c)
	commentCreate.Username = userId.Username
	commentCreate.SnippetId = id

	errs := c.Bind(&commentCreate)
	if errs != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", errs)
	}

	saved, err := controller.usecase.Create(*commentCreate.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(saved))
}

func (controller *CommentController) Update(c echo.Context) error {
	comReq := request.CommentUpdate{}
	ctx := c.Request().Context()
	id := c.Param("id")

	userId := middlewares.GetUser(c)

	if err := c.Bind(&comReq); err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	commentDomain := comReq.ToUpdateDomain()
	commentDomain.SnippetId = id
	commentDomain.Username = userId.Username

	_, err := controller.usecase.Update(*commentDomain, ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(*commentDomain))
}

// Delete saved controller

func (controller *CommentController) Delete(c echo.Context) error {
	comReq := request.CommentDelete{}
	ctx := c.Request().Context()
	id := c.Param("id")
	comReq.SnippetId = id

	userId := middlewares.GetUser(c)

	if err := c.Bind(&comReq); err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	commentDomain := comReq.ToDeleteDomain()
	commentDomain.Username = userId.Username

	_, err := controller.usecase.Delete(*commentDomain, ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(*commentDomain))
}
