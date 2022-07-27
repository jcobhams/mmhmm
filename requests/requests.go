package requests

type (
	GenerateUserIdRequest struct {
		Email string `json:"email"`
	}

	CreateNoteRequest struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	UpdateNoteRequest struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}
)
