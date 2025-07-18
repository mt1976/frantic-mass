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
		view.Common.Errors = append(view.Common.Errors, "Error fetching users")
		view.Common.Status = 500 // Internal Server Error
		view.Common.Success = false
		view.Common.Title = "Choose User"
		view.Common.Description = "An error occurred while fetching users."
		view.Common.Keywords = "error, users"

		return view, err
	}

	/// Assign the list of users to the view
	/// Rangle over the users and assign them to the view
	if len(users) == 0 {
		// If no users are found, set the view's Common fields accordingly
		view.Common.Title = "Choose User"
		view.Common.Description = "No users found."
		view.Common.Keywords = "no users"
		view.Common.Status = 404 // Not Found
		view.Common.Success = false
		logHandler.InfoLogger.Println("No users found")
		// Append an error message to the view's Common errors
		view.Common.Errors = append(view.Common.Errors, "No users found")
		view.Common.Messages = append(view.Common.Messages, "Please create a user first.")
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
		view.Common.Title = "Choose User"
		view.Common.Description = "No valid users found."
		view.Common.Keywords = "no valid users"
		view.Common.Status = 404 // Not Found
		view.Common.Success = false
		view.Common.AddError("No valid users found")
		view.Common.AddMessage("Please create a user first.")
		// Return an error indicating no valid users were found
		return view, fmt.Errorf("no valid users found")
	}

	// Set the common fields for the view
	view.Common.Title = "Choose User"
	view.Common.Description = "Select a user to proceed."
	view.Common.Keywords = "choose user, select user"
	view.Common.Status = 200 // OK
	view.Common.Success = true
	view.Common.Messages = append(view.Common.Messages, "Users loaded successfully")

	logHandler.InfoLogger.Println("ChooseUser view created successfully with", len(view.Users), "users")
	// Return the populated view

	return view, nil
}
