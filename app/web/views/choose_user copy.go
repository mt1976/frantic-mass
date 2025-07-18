package views

type UserChooser struct {
	Users  []User
	Common Common
}

type User struct {
	ID   int
	Name string
}

func CreateUserChooser() UserChooser {
	view := UserChooser{}
	view.Common.SetDefaults() // Initialize the Common view with defaults
	view.Users = []User{}
	return view
}
