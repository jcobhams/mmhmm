package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	repoMocks "github.com/jcobhams/mmhmm/mocks/repositories"
	"github.com/jcobhams/mmhmm/models"
	"github.com/jcobhams/mmhmm/repositories"
)

func TestNoteService_Create(t *testing.T) {

	ctx := context.Background()
	service := NewNoteService()

	input := CreateNoteInput{
		UserId: "user_id",
		Title:  "title",
		Body:   "body",
	}
	repoMock := new(repoMocks.NoteRepositoryInterface)
	repoMock.On("Create", ctx, mock.MatchedBy(func(i repositories.CreateNoteInput) bool {
		return i.Body == input.Body && i.Title == input.Title && i.UserId == input.UserId
	})).Return(models.Note{}, nil)

	n, err := service.Create(ctx, input, repoMock)
	assert.Nil(t, err)
	assert.NotNil(t, n)

	repoMock.AssertExpectations(t)
}

func TestNoteService_List(t *testing.T) {

	ctx := context.Background()
	service := NewNoteService()

	notes := []models.Note{
		{
			Id:     "id1",
			UserId: "user_id",
			Title:  "title1",
			Body:   "body1",
		},
		{
			Id:     "id2",
			UserId: "user_id",
			Title:  "title2",
			Body:   "body2",
		},
	}

	repoMock := new(repoMocks.NoteRepositoryInterface)
	repoMock.On("List", ctx, "user_id").Return(notes, nil)

	notesResult, err := service.List(ctx, "user_id", repoMock)
	assert.Nil(t, err)
	assert.Equal(t, len(notes), len(notesResult))

	repoMock.AssertExpectations(t)
}

func TestNoteService_Get(t *testing.T) {
	ctx := context.Background()
	service := NewNoteService()

	note := models.Note{
		Id:     "id1",
		UserId: "user_id",
		Title:  "title1",
		Body:   "body1",
	}

	repoMock := new(repoMocks.NoteRepositoryInterface)
	repoMock.On("Get", ctx, "user_id", "id1").Return(note, nil)

	n, err := service.Get(ctx, "user_id", "id1", repoMock)
	assert.Nil(t, err)
	assert.Equal(t, n.Id, note.Id)
	assert.Equal(t, n.Title, note.Title)

	repoMock.AssertExpectations(t)
}
