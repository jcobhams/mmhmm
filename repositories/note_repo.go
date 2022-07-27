package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/jcobhams/mmhmm/database"
	"github.com/jcobhams/mmhmm/models"
)

type (
	NoteRepository struct {
		keyPrefix string
		db        database.Database
	}

	CreateNoteInput struct {
		UserId string
		Title  string
		Body   string
	}
)

func NewNoteRepository(db database.Database) *NoteRepository {
	return &NoteRepository{
		keyPrefix: "note",
		db:        db,
	}
}

// Create creates a new note.
func (n *NoteRepository) Create(
	ctx context.Context,
	input CreateNoteInput,
) (models.Note, error) {

	note := models.Note{
		Id:     uuid.New().String(),
		UserId: input.UserId,
		Title:  input.Title,
		Body:   input.Body,
	}

	return note, n.db.Write(input.UserId, note)
}

// List returns all notes belonging to a user.
func (n *NoteRepository) List(ctx context.Context, userId string) ([]models.Note, error) {
	return n.db.Read(userId)
}

// Get returns a note by id.
func (n *NoteRepository) Get(ctx context.Context, userId, noteId string) (models.Note, error) {
	notes, err := n.List(ctx, userId)
	if err != nil {
		return models.Note{}, err
	}

	for _, note := range notes {
		if note.Id == noteId {
			return note, nil
		}
	}
	return models.Note{}, fmt.Errorf("note not found")
}

// Update updates a note by id.
func (n *NoteRepository) Update(ctx context.Context, note models.Note) (models.Note, error) {
	if err := n.Delete(ctx, note.UserId, note.Id); err != nil {
		return models.Note{}, err
	}
	return note, n.db.Write(note.UserId, note)
}

// Delete deletes a note by id.
func (n *NoteRepository) Delete(ctx context.Context, userId, noteId string) error {

	notes, err := n.List(ctx, userId)
	if err != nil {
		return err
	}

	n.db.Delete(userId)

	for _, note := range notes {
		if note.Id == noteId {
			continue
		}
		n.db.Write(userId, note)
	}

	return nil
}
