package handlers

import (
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
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

func User(w http.ResponseWriter, r *http.Request) {

	userID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter

	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	// This is the handler for the edit user page
	uc, err := viewProvider.GetUserView(r.Context(), userID) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error creating UserEdit view", http.StatusInternalServerError)
		return
	}

	render(uc, uc.Context, w)

	logHandler.InfoLogger.Println("Dummy router executed successfully")
}
