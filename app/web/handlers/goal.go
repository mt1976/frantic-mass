package handlers

import (
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func Goal(w http.ResponseWriter, r *http.Request) {

	goalID := getURLParamValue(r, contentProvider.GoalWildcard) // Get the goal ID from the URL parameter
	if goalID == "" {
		http.Error(w, "Goal ID is required", http.StatusBadRequest)
		return
	}
	// This is the handler for the edit user page
	uc, err := viewProvider.Goal(r.Context(), goalID) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error creating Goal view", http.StatusInternalServerError)
		return
	}

	render(uc, uc.Context, w)

	logHandler.InfoLogger.Println("Goal router executed successfully")
}
