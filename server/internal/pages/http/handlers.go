package http

import (
	"net/http"
	"notik/internal/pages"
	"notik/internal/pages/pages_repo"
	"notik/pkg/httpErrors"
	"notik/pkg/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type pagesHandler struct {
	uc pages.Usecase
}

func New(uc pages.Usecase) pages.Handler {
	return &pagesHandler{uc: uc}
}

// func (c *pagesHandler) GetByUserId() echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		userId := ctx.Param("id")
// 		pages, err := c.uc.GetByUserId(ctx.Request().Context(), userId)
// 		if err != nil {
// 			return ctx.JSON(httpErrors.RequestError(err))
// 		}

// 		return ctx.JSON(http.StatusOK, pages)
// 	}
// }

func (c *pagesHandler) Create() echo.HandlerFunc {
	type createInput struct {
		pages_repo.CreateParams
		Title string `validate:"required,min=3,max=80"`
	}
	return func(ctx echo.Context) error {
		var input createInput
		if err := utils.SanitizeRequest(ctx, &input); err != nil {
			return ctx.JSON(httpErrors.RequestError(err))
		}

		cook, err := ctx.Cookie("id")
		if err != nil {
			return ctx.NoContent(http.StatusUnauthorized)
		}
		userId, err := strconv.Atoi(cook.Value)
		if err != nil {
			return ctx.JSON(httpErrors.RequestError(err))
		}

		inputRet := input.CreateParams
		inputRet.Title = input.Title
		inputRet.UserID = int64(userId)

		page, err := c.uc.Create(ctx.Request().Context(), inputRet)
		if err != nil {
			return ctx.JSON(httpErrors.RequestError(err))
		}

		return ctx.JSON(http.StatusOK, page)
	}
}

// func (c *pagesHandler) Delete() echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		id := ctx.Param("id")

// 		if err := c.uc.Delete(ctx.Request().Context(), id); err != nil {
// 			return ctx.JSON(httpErrors.RequestError(err))
// 		}

// 		return ctx.JSON(http.StatusOK, nil)
// 	}
// }

// func (c *pagesHandler) Update() echo.HandlerFunc {
// 	type request struct {
// 		Title string
// 	}
// 	return func(ctx echo.Context) error {
// 		var req request
// 		if err := utils.SanitizeRequest[request](ctx, &req); err != nil {
// 			return ctx.JSON(httpErrors.RequestError(err))
// 		}

// 		data, err := c.uc.Update(ctx.Request().Context(), ctx.Param("id"), req.Title)
// 		if err != nil {
// 			return ctx.JSON(httpErrors.RequestError(err))
// 		}

// 		return ctx.JSON(http.StatusOK, data)
// 	}
// }
