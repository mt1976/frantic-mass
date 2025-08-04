package contentProvider

import (
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/web/actionHelpers"
	methods "github.com/mt1976/frantic-mass/app/web/actionHelpers"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
)

var UserChooserWildcard = ""           // Wildcard for the user ID in the URI
var UserChooserURI = "/users"          // Define the URI for the user chooser
var UserChooserName = "Users"          // Name for the user chooser
var UserChooserIcon = glyphs.Users     // Icon for the user chooser
var UserChooserHover = "Select a user" // Description for the user chooser

type UserChooser struct {
	Users   []User
	Context AppContext
}

type User struct {
	ID      int
	Name    string
	Actions methods.Actions // Actions available for the user, such as edit or delete
}

func CreateUserChooser(view UserChooser) (UserChooser, error) {
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "users"
	view.Context.SetIsViewWorkflow() // Set the request type to GET for viewing users

	view.Users = []User{}

	users, err := user.GetAll()
	if err != nil {
		// Log the error and return an empty view with the error
		logHandler.ErrorLogger.Println("Error fetching users:", err)
		view.Context.AddError("Error fetching users")
		view.Context.AddMessage("An error occurred while fetching users. Please try again later.")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		return view, nil
	}

	/// Assign the list of users to the view
	/// Rangle over the users and assign them to the view
	if len(users) == 0 {
		// If no users are found, set the view's Common fields accordingly

		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.WasSuccessful = false
		// Append an error message to the view's Common errors
		view.Context.AddError("No users found")
		view.Context.AddMessage("No users are available. Please create a user first.")
		logHandler.ErrorLogger.Println("No users found, returning empty UserChooser view")
		return view, nil // No users found, return empty view
	}

	// Iterate over the users and append them to the view
	for _, u := range users {
		// Ensure that the user is not nil before appending
		if u.Audit.DeletedBy != "" {
			continue // Skip deleted users
		}
		// Append the user to the view's Users slice
		addview := User{
			ID:   u.ID,
			Name: u.Username,
		}
		uri := DashboardURI // Use the defined URI for the user dashboard
		if uri == "" {
			uri = "/dash/" + UserWildcard // Default URI for no user ID
		}
		// Replace the placeholder with the actual user ID
		uri = ReplacePathParam(uri, UserWildcard, fmt.Sprintf("%d", u.ID))
		logHandler.InfoLogger.Println("Adding user:", u.Username, "with URI:", uri)

		// Add the user action to the view
		addview.Actions.AddSubmitButton(u.Username, fmt.Sprintf(UserHover, u.Username), UserIcon, uri, methods.READ, "", style.DEFAULT(), css.NONE())
		view.Users = append(view.Users, addview)
	}

	// If no users were added, return an empty view
	if len(view.Users) == 0 {
		logHandler.InfoLogger.Println("No valid users found")

		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.WasSuccessful = false
		view.Context.AddError("No valid users found")
		view.Context.AddMessage("Please create a user first.")
		// Return an error indicating no valid users were found
		return view, nil
	}
	// Set the common fields for the view

	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	// Log the successful creation of the view
	//view.Context.AddMessage("Users loaded successfully")
	view.Context.AddMessage(fmt.Sprintf("Found %d users", len(view.Users)))
	logHandler.InfoLogger.Println("ChooseUser view created successfully with", len(view.Users), "users")

	view.Context.PageActions.Clear() // Clear any existing page actions
	view.Context.PageActions.AddSubmitButton("Add User", "Create a new user", glyphs.Add, ReplacePathParam(UserURI, UserWildcard, actionHelpers.NEW), methods.READ, "", style.NONE(), css.NONE())
	// Return the populated view

	view.Context.AddBreadcrumb(LauncherName, fmt.Sprintf(LauncherHover, view.Context.AppName), LauncherURI, LauncherIcon)
	view.Context.AddBreadcrumb(UserChooserName, UserChooserHover, "", UserChooserIcon)

	return view, nil
}
