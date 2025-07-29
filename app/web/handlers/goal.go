package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func Goal(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id") // Get the user ID from the URL parameter
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	// This is the handler for the edit user page
	uc, err := viewProvider.Goal(r.Context(), id) // Assuming userID is passed in the context or URL
	if err != nil {
		http.Error(w, "Error creating Goal view", http.StatusInternalServerError)
		return
	}

	render(uc, uc.Context, w)

	logHandler.InfoLogger.Println("Dummy router executed successfully")
}
