package views

import "github.com/mt1976/frantic-mass/app/web/helpers"

type UserChooser struct {
	Users   []User
	Context AppContext
}

type User struct {
	ID      int
	Name    string
	Actions helpers.Actions // Actions available for the user, such as edit or delete
}

func CreateUserChooser() UserChooser {
	view := UserChooser{}
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Users = []User{}
	return view
}
