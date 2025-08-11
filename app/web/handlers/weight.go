package handlers

import (
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/actionHelpers"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func ViewWeight(w http.ResponseWriter, r *http.Request) {

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

func CreateWeight(w http.ResponseWriter, r *http.Request) {

	weightID := getURLParamValue(r, contentProvider.WeightWildcard) // Get the weight ID from the URL parameter
	if weightID == "" {
		http.Error(w, "Weight ID is required", http.StatusBadRequest)
		return
	}

	// This is the handler for the edit user page
	uc, err := viewProvider.ViewWeight(r.Context(), weightID) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error creating Weight view", http.StatusInternalServerError)
		return
	}

	render(uc, uc.Context, w)

	logHandler.InfoLogger.Println("Weight router executed successfully")
}
