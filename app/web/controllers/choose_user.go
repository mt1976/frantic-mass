package controllers

import (
	"context"
	"fmt"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/web/views"
)

// Users build the page content for the choose user page.
func Users(ctx context.Context) (views.UserChooser, error) {
	view := views.CreateUserChooser()
	godump.Dump(view)
	// Get the list of users from the user package
	users, err := user.GetAll()
	if err != nil {
		// Log the error and return an empty view with the error
		logHandler.ErrorLogger.Println("Error fetching users:", err)
		view.Context.AddError("Error fetching users")
		view.Context.AddMessage("Please try again later.")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		view.Context.PageTitle = "Choose User"
		view.Context.PageSummary = "An error occurred while fetching users."
		view.Context.PageKeywords = "error, users"

		return view, err
	}

	/// Assign the list of users to the view
	/// Rangle over the users and assign them to the view
	if len(users) == 0 {
		// If no users are found, set the view's Common fields accordingly
		view.Context.PageTitle = "Choose User"
		view.Context.PageSummary = "No users found."
		view.Context.PageKeywords = "no users"
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.WasSuccessful = false
		// Append an error message to the view's Common errors
		view.Context.AddError("No users found")
		view.Context.AddMessage("Please create a user first.")
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
		view.Context.PageTitle = "Choose User"
		view.Context.PageSummary = "No valid users found."
		view.Context.PageKeywords = "no valid users"
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.WasSuccessful = false
		view.Context.AddError("No valid users found")
		view.Context.AddMessage("Please create a user first.")
		// Return an error indicating no valid users were found
		return view, fmt.Errorf("no valid users found")
	}

	// Set the common fields for the view
	view.Context.PageTitle = "Choose User"
	view.Context.PageSummary = "Select a user to proceed."
	view.Context.PageKeywords = "choose user, select user"
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	// Log the successful creation of the view
	view.Context.AddMessage("Users loaded successfully")
	view.Context.AddMessage(fmt.Sprintf("Found %d users", len(view.Users)))
	logHandler.InfoLogger.Println("ChooseUser view created successfully with", len(view.Users), "users")
	// Return the populated view

	return view, nil
}
