package handlers

import (
	"fmt"
	"net/http"

	"github.com/mt1976/frantic-core/htmlHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/actionHelpers"
	"github.com/mt1976/frantic-mass/app/web/contentActioner"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func ViewOrEditGoal(w http.ResponseWriter, r *http.Request) {
	logHandler.InfoLogger.Println("Goal View/Edit called with method:", r.Method)
	userID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter
	goalID := getURLParamValue(r, contentProvider.GoalWildcard) // Get the goal ID from the URL parameter
	if goalID == "" {
		http.Error(w, "Goal ID is required", http.StatusBadRequest)
		return
	}
	if userID == "" || userID == actionHelpers.NEW {
		http.Error(w, "User ID is required and cannot be new", http.StatusMethodNotAllowed)
		return
	}

	userID_int, err := contentProvider.StringToInt(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	uc := contentProvider.GoalView{}

	if goalID == actionHelpers.NEW {
		uc, err = viewProvider.NewGoal(r.Context(), userID_int) // This is the handler for creating a new user
	} else {
		uc, err = viewProvider.ViewGoal(r.Context(), goalID) // Assuming userID is passed in the context or URL

	}
	// This is the handler for the edit user page
	if err != nil {
		http.Error(w, "Error creating Goal view", http.StatusInternalServerError)
		return
	}

	render(uc, uc.Context, w)

	logHandler.InfoLogger.Println("Goal router executed successfully")
}

func CreateNewGoal(w http.ResponseWriter, r *http.Request) {
	logHandler.InfoLogger.Println("CreateNewGoal called with method:", r.Method)
	goalID := getURLParamValue(r, contentProvider.GoalWildcard) // Get the goal ID from the URL parameter
	userID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter

	logHandler.InfoLogger.Printf("Goal ID from URL: %s, User ID: %s, Method: %s", goalID, userID, r.Method)

	int_userID := htmlHelpers.ValueToInt(userID)

	if goalID == "" || goalID != actionHelpers.NEW {
		http.Error(w, "Not a valid goal ID, must be 'new'", http.StatusBadRequest)
		return
	}

	if int_userID <= 0 || int_userID >= 999999999 {
		logHandler.ErrorLogger.Printf("Invalid user ID: %s", userID)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	logHandler.InfoLogger.Printf("Goal ID from URL: %s, Method: %s", goalID, r.Method)
	// Validate the request method
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s not valid", r.Method), http.StatusMethodNotAllowed)
		return
	}

	logHandler.EventLogger.Printf("User request received for userID: %s with method: %s", userID, r.Method)

	if goalID != actionHelpers.NEW {
		http.Error(w, "Method not allowed for existing goal", http.StatusMethodNotAllowed)
		return
	}
	logHandler.InfoLogger.Printf("Creating new goal log! ID: %s", goalID)

	// Handle the PUT request to update user details
	uc, err := contentActioner.NewGoal(w, r, int_userID) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error updating goal", http.StatusInternalServerError)
		return
	}
	uc.Context.AddMessage("Goal create screen generated successfully")
	uc.Context.HttpStatusCode = http.StatusOK // Set the HTTP status code to 200 OK
	uc.Context.WasSuccessful = true
	render(uc, uc.Context, w)
}

func UpdateGoal(w http.ResponseWriter, r *http.Request) {

	logHandler.InfoLogger.Println("Goal Update called with method:", r.Method)
	userID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter
	goalID := getURLParamValue(r, contentProvider.GoalWildcard) // Get the goal ID from the URL parameter

	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	logHandler.InfoLogger.Printf("User ID from URL: %s, Method: %s", userID, r.Method)
	// Validate the request method
	if r.Method != http.MethodPut {
		http.Error(w, fmt.Sprintf("Method %s not valid", r.Method), http.StatusMethodNotAllowed)
		return
	}

	logHandler.EventLogger.Printf("User request received for userID: %s with method: %s", userID, r.Method)

	if goalID == actionHelpers.NEW {
		http.Error(w, "Method not allowed for new goal", http.StatusMethodNotAllowed)
		return
	}
	logHandler.InfoLogger.Printf("Update existing goal! ID: %s", goalID)

	// Convert userID to int for processing
	// This assumes userID is a string representation of an integer
	// TODO: Add Error handling to htmlHelpers.ValueToInt
	int_userID := htmlHelpers.ValueToInt(userID)
	int_goalID := htmlHelpers.ValueToInt(goalID)
	if int_userID <= 0 || int_userID >= 999999999 {
		logHandler.ErrorLogger.Printf("Invalid user ID: %s", userID)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if int_goalID <= 0 || int_goalID >= 999999999 {
		logHandler.ErrorLogger.Printf("Invalid goal ID: %s", goalID)
		http.Error(w, "Invalid goal ID", http.StatusBadRequest)
		return
	}

	// Handle the PUT request to update user details
	uc, err := contentActioner.UpdateGoal(w, r, int_userID, int_goalID) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	//uc.Context.HttpStatusCode = http.StatusOK // Set the HTTP status code to 200 OK
	//uc.Context.WasSuccessful = true
	render(uc, uc.Context, w)
}
