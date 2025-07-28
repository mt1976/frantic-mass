package contentProvider

import (
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

var UserURI = "/user/{id}" // Define the URI for the user chooser

type UserView struct {
	ID       int
	User     user.User
	Baseline baseline.Baseline
	Context  AppContext
}

func UserEdit(view UserView, userID string) (UserView, error) {
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "users"

	userIdInt, err := StringToInt(userID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error converting userID to int:", err)
		view.Context.AddError("Invalid user ID format")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.User = user.User{}

	UserRecord, err := user.GetBy(user.FIELD_ID, userIdInt)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching user:", err)
		view.Context.AddError("Error fetching user")
		view.Context.AddMessage("An error occurred while fetching user details. Please try again later.")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.User = UserRecord

	// Fetch the user's baseline data
	baselineRecord, err := baseline.GetBy(baseline.FIELD_UserID, userIdInt)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching baseline data:", err)
		view.Context.AddError("Error fetching baseline data")
		view.Context.AddMessage("An error occurred while fetching baseline data. Please try again later.")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.Baseline = baselineRecord

	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	// Log the successful creation of the view
	//view.Context.AddMessage("Users loaded successfully")
	view.Context.AddMessage(fmt.Sprintf("Found user %s", view.User.Username))
	uri := DashboardURI // Use the defined URI for the dashboard
	uri = ReplacePathParam(uri, "id", fmt.Sprintf("%d", view.User.ID))
	view.Context.PageActions.Add(helpers.NewAction("Back", "Back to User Chooser", glyphs.Back, uri, helpers.GET, ""))
	logHandler.InfoLogger.Println("UserEdit view created successfully with user", view.User.Username)
	// Return the populated view

	return view, nil
}
