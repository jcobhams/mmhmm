package controllers

import (
	"github.com/labstack/echo/v4"

	"github.com/jcobhams/mmhmm/logging"
	"github.com/jcobhams/mmhmm/repositories"
	"github.com/jcobhams/mmhmm/requests"
	"github.com/jcobhams/mmhmm/services"
)

type (
	NoteController struct{}
)

func NewNoteController() *NoteController {
	return &NoteController{}
}

// Create creates a new note
func (n *NoteController) Create(
	noteService services.NoteServiceInterface,
	noteRepo repositories.NoteRepositoryInterface,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var requestBody requests.CreateNoteRequest
		err := c.Bind(&requestBody)
		if err != nil {
			return ErrorResponseHandler(c, 400, err)
		}

		ctx := logging.WrapContext(c.Request().Context(), "controllers.CreateNote")

		input := services.CreateNoteInput{
			UserId: c.Param("userId"),
			Title:  requestBody.Title,
			Body:   requestBody.Body,
		}
		note, err := noteService.Create(ctx, input, noteRepo)
		if err != nil {
			return ErrorResponseHandler(c, 400, err)
		}

		body := ResponseBody{
			Message: "Note created",
			Payload: map[string]interface{}{
				"note": note,
			},
		}
		return ResponseHandler(c, 201, body)
	}
}

//List lists all notes for a user
func (n *NoteController) List(
	noteService services.NoteServiceInterface,
	noteRepo repositories.NoteRepositoryInterface,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := logging.WrapContext(c.Request().Context(), "controllers.ListNotes")

		userId := c.Param("userId")
		notes, err := noteService.List(ctx, userId, noteRepo)
		if err != nil {
			return ErrorResponseHandler(c, 400, err)
		}

		body := ResponseBody{
			Message: "OK",
			Payload: map[string]interface{}{
				"notes": notes,
			},
		}
		return ResponseHandler(c, 200, body)
	}
}

// Get gets a note
func (n *NoteController) Get(
	noteService services.NoteServiceInterface,
	noteRepo repositories.NoteRepositoryInterface,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := logging.WrapContext(c.Request().Context(), "controllers.GetNote")

		userId := c.Param("userId")
		noteId := c.Param("noteId")
		note, err := noteService.Get(ctx, userId, noteId, noteRepo)
		if err != nil {
			return ErrorResponseHandler(c, 400, err)
		}

		body := ResponseBody{
			Message: "OK",
			Payload: map[string]interface{}{
				"note": note,
			},
		}
		return ResponseHandler(c, 200, body)
	}

}

// Update updates a note
func (n *NoteController) Update(
	noteService services.NoteServiceInterface,
	noteRepo repositories.NoteRepositoryInterface,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var requestBody requests.UpdateNoteRequest
		err := c.Bind(&requestBody)
		if err != nil {
			return ErrorResponseHandler(c, 400, err)
		}

		ctx := logging.WrapContext(c.Request().Context(), "controllers.UpdateNote")

		input := services.UpdateNoteInput{
			Title:  requestBody.Title,
			Body:   requestBody.Body,
			UserId: c.Param("userId"),
			NoteId: c.Param("noteId"),
		}
		note, err := noteService.Update(ctx, input, noteRepo)
		if err != nil {
			return ErrorResponseHandler(c, 400, err)
		}

		body := ResponseBody{
			Message: "Note updated",
			Payload: map[string]interface{}{
				"note": note,
			},
		}
		return ResponseHandler(c, 200, body)
	}
}

// Delete deletes a note
func (n *NoteController) Delete(
	noteService services.NoteServiceInterface,
	noteRepo repositories.NoteRepositoryInterface,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := logging.WrapContext(c.Request().Context(), "controllers.DeleteNote")

		userId := c.Param("userId")
		noteId := c.Param("noteId")
		err := noteService.Delete(ctx, userId, noteId, noteRepo)
		if err != nil {
			return ErrorResponseHandler(c, 400, err)
		}

		body := ResponseBody{
			Message: "Note deleted",
		}
		return ResponseHandler(c, 200, body)
	}
}
