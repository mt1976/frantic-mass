package handlers

import (
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func Weight(w http.ResponseWriter, r *http.Request) {

	weightID := getURLParamValue(r, contentProvider.WeightWildcard) // Get the weight ID from the URL parameter
	if weightID == "" {
		http.Error(w, "Weight ID is required", http.StatusBadRequest)
		return
	}

	// This is the handler for the edit user page
	uc, err := viewProvider.Weight(r.Context(), weightID) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error creating Weight view", http.StatusInternalServerError)
		return
	}

	render(uc, uc.Context, w)

	logHandler.InfoLogger.Println("Weight router executed successfully")
}
