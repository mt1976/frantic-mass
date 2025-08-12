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

func UserChooser(w http.ResponseWriter, r *http.Request) {

	// This is the handler for the choose user page
	uc, err := viewProvider.Users(r.Context())
	if err != nil {
		http.Error(w, "Error creating UserChooser view", http.StatusInternalServerError)
		return
	}

	render(uc, uc.Context, w)

	logHandler.InfoLogger.Println("Dummy router executed successfully")

}

func ViewOrEditUser(w http.ResponseWriter, r *http.Request) {

	userID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter

	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	logHandler.InfoLogger.Printf("User ID from URL: %s, Method: %s", userID, r.Method)
	// Validate the request method
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Method %s not valid", r.Method), http.StatusMethodNotAllowed)
		return
	}

	logHandler.EventLogger.Printf("User request received for userID: %s with method: %s", userID, r.Method)

	// Handle the GET request to view or edit user details
	logHandler.InfoLogger.Printf("User ID: %s", userID)
	uc := contentProvider.UserView{}
	var err error
	if userID == actionHelpers.NEW {
		uc, err = viewProvider.CreateUser(r.Context(), userID) // This is the handler for creating a new user
	} else {
		uc, err = viewProvider.GetUser(r.Context(), userID) // Assuming userID is passed in the context or URL
	}
	if err != nil {
		http.Error(w, "Error creating UserEdit view", http.StatusInternalServerError)
		return
	}
	render(uc, uc.Context, w)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {

	userID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter

	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	logHandler.InfoLogger.Printf("User ID from URL: %s, Method: %s", userID, r.Method)
	// Validate the request method
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s not valid", r.Method), http.StatusMethodNotAllowed)
		return
	}

	logHandler.EventLogger.Printf("User request received for userID: %s with method: %s", userID, r.Method)

	if userID != actionHelpers.NEW {
		http.Error(w, "Method not allowed for existing user", http.StatusMethodNotAllowed)
		return
	}
	logHandler.InfoLogger.Printf("Creating new user! pseudo: %s", userID)

	// Handle the PUT request to update user details
	uc, err := contentActioner.NewUser(w, r) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}
	uc.Context.AddMessage("User create screen generated successfully")
	uc.Context.HttpStatusCode = http.StatusOK // Set the HTTP status code to 200 OK
	uc.Context.WasSuccessful = true
	render(uc, uc.Context, w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	logHandler.InfoLogger.Println("UserUpdate called with method:", r.Method)
	userID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter

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

	if userID == actionHelpers.NEW {
		http.Error(w, "Method not allowed for new user", http.StatusMethodNotAllowed)
		return
	}
	logHandler.InfoLogger.Printf("Update existing user! ID: %s", userID)

	// Convert userID to int for processing
	// This assumes userID is a string representation of an integer
	// TODO: Add Error handling to htmlHelpers.ValueToInt
	int_userID := htmlHelpers.ValueToInt(userID)
	if int_userID <= 0 || int_userID >= 999999999 {
		logHandler.ErrorLogger.Printf("Invalid user ID: %s", userID)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Handle the PUT request to update user details
	uc, err := contentActioner.UpdateUser(w, r, int_userID) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	//uc.Context.HttpStatusCode = http.StatusOK // Set the HTTP status code to 200 OK
	//uc.Context.WasSuccessful = true
	render(uc, uc.Context, w)
}
