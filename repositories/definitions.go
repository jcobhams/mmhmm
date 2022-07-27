package repositories

import (
	"context"

	"github.com/jcobhams/mmhmm/models"
)

type (

	// NoteRepositoryInterface is an interface for a note repository.
	NoteRepositoryInterface interface {
		Create(ctx context.Context, input CreateNoteInput) (models.Note, error)
		List(ctx context.Context, userId string) ([]models.Note, error)
		Get(ctx context.Context, userId, noteId string) (models.Note, error)
		Update(ctx context.Context, note models.Note) (models.Note, error)
		Delete(ctx context.Context, userId, noteId string) error
	}
)
