package views

type UserChooser struct {
	Users       []User
	SessionData AppContext
}

type User struct {
	ID   int
	Name string
}

func CreateUserChooser() UserChooser {
	view := UserChooser{}
	view.SessionData.SetDefaults() // Initialize the Common view with defaults
	view.Users = []User{}
	return view
}
