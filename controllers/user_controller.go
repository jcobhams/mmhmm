package controllers

import (
	"github.com/labstack/echo/v4"

	"github.com/jcobhams/mmhmm/logging"
	"github.com/jcobhams/mmhmm/requests"
	"github.com/jcobhams/mmhmm/services"
)

type (
	UserController struct {
	}
)

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GenerateUserId(
	userService services.UserIdGenerator,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var requestBody requests.GenerateUserIdRequest
		err := c.Bind(&requestBody)
		if err != nil {
			return ErrorResponseHandler(c, 400, err)
		}

		ctx := logging.WrapContext(c.Request().Context(), "controllers.CreateUser")

		userId, err := userService.GenerateUserId(ctx, requestBody.Email)
		if err != nil {
			return ErrorResponseHandler(c, 400, err)
		}

		body := ResponseBody{
			Message: "OK",
			Payload: map[string]interface{}{
				"userId": userId,
			},
		}
		return ResponseHandler(c, 201, body)
	}
}
