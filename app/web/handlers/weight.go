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

func ViewOrEditWeightLogEntry(w http.ResponseWriter, r *http.Request) {
	logHandler.InfoLogger.Println("ViewOrEditWeightLogEntry called with method:", r.Method)
	weightID := getURLParamValue(r, contentProvider.WeightWildcard) // Get the weight ID from the URL parameter
	if weightID == "" {
		http.Error(w, "Weight ID is required", http.StatusBadRequest)
		return
	}

	userID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	userIDInt, err := contentProvider.StringToInt(userID)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	uc := contentProvider.WeightView{}

	// This is the handler for the edit user page
	if weightID == actionHelpers.NEW {
		logHandler.EventLogger.Println("Creating new Weight entry for user ID:", userIDInt)
		// If the weight ID is "new", redirect to the create weight handler
		uc, err = viewProvider.NewWeight(r.Context(), userIDInt) // Assuming userID is passed in the context or URL

		if err != nil {
			http.Error(w, "Error creating Weight view", http.StatusInternalServerError)
			return
		}
	} else {
		logHandler.EventLogger.Println("Editing Weight entry for user ID:", userIDInt, "and weight ID:", weightID)
		uc, err = viewProvider.ViewWeight(r.Context(), weightID) // Assuming userID is passed in the context or URL
		if err != nil {
			http.Error(w, "Error creating Weight view", http.StatusInternalServerError)
			return
		}
	}

	render(uc, uc.Context, w)

	logHandler.InfoLogger.Println("Weight router executed successfully")
}

func CreateNewWeightLogEntry(w http.ResponseWriter, r *http.Request) {
	logHandler.InfoLogger.Println("CreateNewWeightLogEntry called with method:", r.Method)
	weightLogID := getURLParamValue(r, contentProvider.WeightWildcard) // Get the weight log ID from the URL parameter
	userID := getURLParamValue(r, contentProvider.UserWildcard)        // Get the user ID from the URL parameter
	int_userID := htmlHelpers.ValueToInt(userID)

	if weightLogID == "" {
		http.Error(w, "Weight Log ID is required", http.StatusBadRequest)
		return
	}

	if int_userID <= 0 || int_userID >= 999999999 {
		logHandler.ErrorLogger.Printf("Invalid user ID: %s", userID)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	logHandler.InfoLogger.Printf("Weight Log ID from URL: %s, Method: %s", weightLogID, r.Method)
	// Validate the request method
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s not valid", r.Method), http.StatusMethodNotAllowed)
		return
	}

	logHandler.EventLogger.Printf("User request received for userID: %s with method: %s", weightLogID, r.Method)

	if weightLogID != actionHelpers.NEW {
		http.Error(w, "Method not allowed for existing weight log", http.StatusMethodNotAllowed)
		return
	}
	logHandler.InfoLogger.Printf("Creating new weight log! ID: %s", weightLogID)

	// Handle the PUT request to update user details
	uc, err := contentActioner.NewWeightLogEntry(w, r, int_userID) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error updating weight log", http.StatusInternalServerError)
		return
	}
	uc.Context.AddMessage("Weight log create screen generated successfully")
	uc.Context.HttpStatusCode = http.StatusOK // Set the HTTP status code to 200 OK
	uc.Context.WasSuccessful = true
	render(uc, uc.Context, w)
}

func UpdateWeightLogEntry(w http.ResponseWriter, r *http.Request) {

	logHandler.InfoLogger.Println("WeightUpdate called with method:", r.Method)
	userID := getURLParamValue(r, contentProvider.UserWildcard)     // Get the user ID from the URL parameter
	weightID := getURLParamValue(r, contentProvider.WeightWildcard) // Get the weight ID from the URL parameter

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
	int_weightID := htmlHelpers.ValueToInt(weightID)
	if int_userID <= 0 || int_userID >= 999999999 {
		logHandler.ErrorLogger.Printf("Invalid user ID: %s", userID)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if int_weightID <= 0 || int_weightID >= 999999999 {
		logHandler.ErrorLogger.Printf("Invalid weight ID: %s", weightID)
		http.Error(w, "Invalid weight ID", http.StatusBadRequest)
		return
	}

	// Handle the PUT request to update user details
	uc, err := contentActioner.UpdateWeightLogEntry(w, r, int_userID, int_weightID) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	//uc.Context.HttpStatusCode = http.StatusOK // Set the HTTP status code to 200 OK
	//uc.Context.WasSuccessful = true
	render(uc, uc.Context, w)
}
