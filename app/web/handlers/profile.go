package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	// This is a dummy router function
	id := chi.URLParam(r, "id") // Get the user ID from the URL parameter

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	logHandler.InfoLogger.Println("Fetching profile for user ID:", id)
	// CConvert the ID to an integer if necessary, or handle it as a string
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	dl, err := viewProvider.Dashboard(context.TODO(), userId)
	if err != nil {
		logHandler.ErrorLogger.Println("Error creating Profile view:", err)
	} else {
		logHandler.EventLogger.Println("Profile view created successfully")
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	// Write a simple response
	w.WriteHeader(dl.Context.HttpStatusCode) // Set the HTTP status code

	render(dl, dl.Context, w)

	logHandler.InfoLogger.Println("Dummy router executed successfully")
}
