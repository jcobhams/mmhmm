package services

import (
	"context"

	"github.com/jcobhams/mmhmm/models"
	"github.com/jcobhams/mmhmm/repositories"
)

type (
	NoteServiceInterface interface {
		Create(
			ctx context.Context,
			input CreateNoteInput,
			noteRepo repositories.NoteRepositoryInterface,
		) (models.Note, error)
		List(
			ctx context.Context,
			userId string,
			noteRepo repositories.NoteRepositoryInterface,
		) ([]models.Note, error)
		Get(
			ctx context.Context,
			userId, noteId string,
			noteRepo repositories.NoteRepositoryInterface,
		) (models.Note, error)
		Update(
			ctx context.Context,
			input UpdateNoteInput,
			noteRepo repositories.NoteRepositoryInterface,
		) (models.Note, error)
		Delete(
			ctx context.Context,
			userId, noteId string,
			noteRepo repositories.NoteRepositoryInterface,
		) error
	}

	UserIdGenerator interface {
		GenerateUserId(ctx context.Context, email string) (string, error)
	}

	UserServiceInterface interface {
		UserIdGenerator
	}
)
