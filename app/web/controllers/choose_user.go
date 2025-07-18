package controllers

import (
	"context"
	"fmt"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/web/views"
)

// UserChooser build the page content for the choose user page.
func UserChooser(ctx context.Context) (views.UserChooser, error) {
	view := views.CreateUserChooser()
	godump.Dump(view)
	// Get the list of users from the user package
	users, err := user.GetAll()
	if err != nil {
		// Log the error and return an empty view with the error
		logHandler.ErrorLogger.Println("Error fetching users:", err)
		view.SessionData.AddError("Error fetching users")
		view.SessionData.AddMessage("Please try again later.")
		view.SessionData.HttpStatusCode = 500 // Internal Server Error
		view.SessionData.WasSuccessful = false
		view.SessionData.PageTitle = "Choose User"
		view.SessionData.PageSummary = "An error occurred while fetching users."
		view.SessionData.PageKeywords = "error, users"

		return view, err
	}

	/// Assign the list of users to the view
	/// Rangle over the users and assign them to the view
	if len(users) == 0 {
		// If no users are found, set the view's Common fields accordingly
		view.SessionData.PageTitle = "Choose User"
		view.SessionData.PageSummary = "No users found."
		view.SessionData.PageKeywords = "no users"
		view.SessionData.HttpStatusCode = 404 // Not Found
		view.SessionData.WasSuccessful = false
		// Append an error message to the view's Common errors
		view.SessionData.AddError("No users found")
		view.SessionData.AddMessage("Please create a user first.")
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
		view.Users = append(view.Users, views.User{
			ID:   u.ID,
			Name: u.Username,
		})
	}

	// If no users were added, return an empty view
	if len(view.Users) == 0 {
		logHandler.InfoLogger.Println("No valid users found")
		view.SessionData.PageTitle = "Choose User"
		view.SessionData.PageSummary = "No valid users found."
		view.SessionData.PageKeywords = "no valid users"
		view.SessionData.HttpStatusCode = 404 // Not Found
		view.SessionData.WasSuccessful = false
		view.SessionData.AddError("No valid users found")
		view.SessionData.AddMessage("Please create a user first.")
		// Return an error indicating no valid users were found
		return view, fmt.Errorf("no valid users found")
	}

	// Set the common fields for the view
	view.SessionData.PageTitle = "Choose User"
	view.SessionData.PageSummary = "Select a user to proceed."
	view.SessionData.PageKeywords = "choose user, select user"
	view.SessionData.HttpStatusCode = 200 // OK
	view.SessionData.WasSuccessful = true
	// Log the successful creation of the view
	view.SessionData.AddMessage("Users loaded successfully")
	view.SessionData.AddMessage(fmt.Sprintf("Found %d users", len(view.Users)))
	logHandler.InfoLogger.Println("ChooseUser view created successfully with", len(view.Users), "users")
	// Return the populated view

	return view, nil
}
