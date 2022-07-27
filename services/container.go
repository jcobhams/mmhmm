package services

type (
	Container struct {
		NoteService NoteServiceInterface
		UserService UserServiceInterface
	}
)

func NewContainer() *Container {
	return &Container{
		NoteService: NewNoteService(),
		UserService: NewUserService(),
	}
}
