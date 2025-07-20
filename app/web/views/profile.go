package views

import (
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/user"
)

type Profile struct {
	User    User
	Context AppContext
}

func BuildProfile(userId int) Profile {
	view := Profile{}
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "profile"
	view.User = User{}

	// Here you would typically fetch the user data based on userId
	UserRecord, err := user.GetBy(user.FIELD_ID, userId)
	if err != nil {
		view.Context.WasSuccessful = false
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.AddError("User not found")
		view.Context.AddMessage("Please check the user ID and try again.")
		logHandler.ErrorLogger.Println("Error fetching user:", err)
		return view
	}
	if UserRecord.ID == 0 {
		view.Context.WasSuccessful = false
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.AddError("User not found")
		view.Context.AddMessage("Please check the user ID and try again.")
		logHandler.ErrorLogger.Println("Error fetching user: user record not found")
		return view
	}

	view.User.ID = UserRecord.ID
	view.User.Name = UserRecord.Username
	// Add more fields as necessary

	view.Context.PageTitle = "User Profile"
	view.Context.PageKeywords = "user, profile"
	view.Context.PageSummary = "Display the user information for the selected user."
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true

	logHandler.InfoLogger.Println("Profile view created for user ID:", userId)

	return view
}
