package models

type (
	// User is a user of the system.
	User struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"-"`
	}

	Note struct {
		Id     string `json:"id"`
		UserId string `json:"userId"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}
)
