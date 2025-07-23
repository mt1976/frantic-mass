package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/controllers"
)

func Projection(w http.ResponseWriter, r *http.Request) {
	// This is a dummy router function
	userID := chi.URLParam(r, "id")   // Get the user ID from the URL parameter
	goalID := chi.URLParam(r, "goal") // Get the goal ID from the URL parameter

	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	if goalID == "" {
		http.Error(w, "Goal ID is required", http.StatusBadRequest)
		return
	}
	logHandler.InfoLogger.Println("Fetching profile for user ID:", userID, " and goal ID:", goalID)
	// CConvert the ID to an integer if necessary, or handle it as a string
	userIdInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	goalIDInt, err := strconv.Atoi(goalID)
	if err != nil {
		http.Error(w, "Invalid goal ID", http.StatusBadRequest)
		return
	}

	dl, err := controllers.Projection(context.TODO(), userIdInt, goalIDInt)
	if err != nil {
		logHandler.ErrorLogger.Println("Error creating Projection view:", err)
	} else {
		logHandler.EventLogger.Println("Projection view created successfully, user ID:", userIdInt, "goal ID:", goalIDInt)
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
