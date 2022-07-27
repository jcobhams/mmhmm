package controllers

import (
	"github.com/labstack/echo/v4"

	"github.com/jcobhams/mmhmm/repositories"
	"github.com/jcobhams/mmhmm/services"
)

type (
	Controllers struct {
		NoteController *NoteController
		UserController *UserController
	}

	ResponseBody struct {
		Message string                 `json:"message"`
		Payload map[string]interface{} `json:"payload,omitempty"`
	}
)

func New() *Controllers {
	return &Controllers{
		NoteController: NewNoteController(),
		UserController: NewUserController(),
	}
}

func BindRoutes(e *echo.Echo, sc *services.Container, rc *repositories.Container) {
	controllers := New()

	apiVersion := e.Group("/v1")

	// Notes Endpoints
	{
		notes := apiVersion.Group("/notes")
		notes.POST("/:userId", controllers.NoteController.Create(sc.NoteService, rc.NoteRepository))
		notes.GET("/:userId", controllers.NoteController.List(sc.NoteService, rc.NoteRepository))
		notes.GET("/:userId/:noteId", controllers.NoteController.Get(sc.NoteService, rc.NoteRepository))
		notes.PUT("/:userId/:noteId", controllers.NoteController.Update(sc.NoteService, rc.NoteRepository))

	}

	// Users Endpoints
	{
		users := apiVersion.Group("/users")
		users.POST("", controllers.UserController.GenerateUserId(sc.UserService))
	}
}

func ErrorResponseHandler(c echo.Context, statusCode int, err error) error {
	return ResponseHandler(c, statusCode, ResponseBody{
		Message: err.Error(),
	})
}

func ResponseHandler(c echo.Context, statusCode int, body ResponseBody) error {
	return c.JSON(statusCode, body)
}
