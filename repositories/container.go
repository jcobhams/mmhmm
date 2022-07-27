package repositories

import "github.com/jcobhams/mmhmm/database"

type (
	Container struct {
		NoteRepository NoteRepositoryInterface
	}
)

func NewContainer(db database.Database) *Container {
	return &Container{
		NoteRepository: NewNoteRepository(db),
	}
}
