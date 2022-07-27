package services

import (
	"context"
	"fmt"

	"github.com/jcobhams/mmhmm/models"
	"github.com/jcobhams/mmhmm/repositories"
)

type (
	NoteService struct{}

	CreateNoteInput struct {
		UserId string
		Title  string
		Body   string
	}

	UpdateNoteInput struct {
		Title  string
		Body   string
		UserId string
		NoteId string
	}
)

var _ NoteServiceInterface = (*NoteService)(nil)

func NewNoteService() *NoteService {
	return &NoteService{}
}

// Create creates a new note.
func (n *NoteService) Create(
	ctx context.Context,
	input CreateNoteInput,
	noteRepo repositories.NoteRepositoryInterface,
) (models.Note, error) {
	if input.UserId == "" {
		return models.Note{}, fmt.Errorf("user_id is required")
	}

	if input.Title == "" {
		return models.Note{}, fmt.Errorf("title is required")
	}

	if input.Body == "" {
		return models.Note{}, fmt.Errorf("body is required")
	}

	return noteRepo.Create(ctx, repositories.CreateNoteInput(input))
}

// List returns all notes.
func (n *NoteService) List(
	ctx context.Context,
	userId string,
	noteRepo repositories.NoteRepositoryInterface,
) ([]models.Note, error) {
	return noteRepo.List(ctx, userId)
}

// Get returns a note by id.
func (n *NoteService) Get(ctx context.Context,
	userId, noteId string,
	noteRepo repositories.NoteRepositoryInterface,
) (models.Note, error) {
	return noteRepo.Get(ctx, userId, noteId)
}

// Update updates a note.
func (n *NoteService) Update(
	ctx context.Context,
	input UpdateNoteInput,
	noteRepo repositories.NoteRepositoryInterface,
) (models.Note, error) {
	if input.NoteId == "" {
		return models.Note{}, fmt.Errorf("note id is required")
	}

	if input.UserId == "" {
		return models.Note{}, fmt.Errorf("user id is required")
	}

	note := models.Note{
		Id:     input.NoteId,
		UserId: input.UserId,
		Title:  input.Title,
		Body:   input.Body,
	}
	return noteRepo.Update(ctx, note)
}

// Delete deletes a note.
func (n *NoteService) Delete(
	ctx context.Context,
	userId, noteId string,
	noteRepo repositories.NoteRepositoryInterface,
) error {
	return noteRepo.Delete(ctx, userId, noteId)
}
