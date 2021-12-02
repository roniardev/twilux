package controllers

import (
	"net/http"
	"strings"
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
	comment, err := controller.usecase.GetAll(id, ctx)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}

	return controllers.SuccessResponse(c, []string{"Get all comment succed."}, response.ToListDomain(comment))
}

func (controller *CommentController) GetAllUser(c echo.Context) error {
	ctx := c.Request().Context()
	userId := middlewares.GetUser(c)
	comment, err := controller.usecase.GetAllUser(userId.Username, ctx)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}

	return controllers.SuccessResponse(c, []string{"Get all comment succed."}, response.ToListDomain(comment))
}

func (controller *CommentController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	id = strings.Replace(id, "/", "", -1)

	var commentCreate request.CommentCreate
	userId := middlewares.GetUser(c)
	commentCreate.Username = userId.Username
	commentCreate.SnippetId = id

	errs := c.Bind(&commentCreate)
	if errs != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", errs)
	}

	comment, err := controller.usecase.Create(*commentCreate.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, []string{"Create comment succed."}, response.FromDomain(comment))
}

func (controller *CommentController) Update(c echo.Context) error {
	comReq := request.CommentUpdate{}
	ctx := c.Request().Context()
	id := c.Param("id")
	id = strings.Replace(id, "/", "", -1)

	userId := middlewares.GetUser(c)

	if err := c.Bind(&comReq); err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	commentDomain := comReq.ToUpdateDomain()
	commentDomain.SnippetId = id
	commentDomain.Username = userId.Username

	res, err := controller.usecase.Update(*commentDomain, ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, []string{"Update comment succed."}, response.FromUpdateDomain(res))
}

// Delete comment controller

func (controller *CommentController) Delete(c echo.Context) error {
	comReq := request.CommentDelete{}
	ctx := c.Request().Context()
	id := c.Param("id")
	id = strings.Replace(id, "/", "", -1)

	comReq.SnippetId = id

	userId := middlewares.GetUser(c)

	if err := c.Bind(&comReq); err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	commentDomain := comReq.ToDeleteDomain()
	commentDomain.Username = userId.Username

	res, err := controller.usecase.Delete(*commentDomain, ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, []string{"Delete comment succed."}, response.FromDeleteDomain(res))
}
