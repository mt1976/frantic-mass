package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	// This is a dummy router function
	currentUserID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter
	logHandler.InfoLogger.Println("Paths from URL:", r.URL.Path)
	logHandler.InfoLogger.Println("Dashboard URI:", contentProvider.DashboardURI)
	logHandler.InfoLogger.Println("User ID wildcard:", contentProvider.UserWildcard)
	logHandler.ErrorLogger.Println("User ID:", currentUserID)

	if currentUserID == "" {

		logHandler.ErrorLogger.Println("Paths from URL:", r.URL.Path)
		logHandler.ErrorLogger.Println("Dashboard URI:", contentProvider.DashboardURI)
		logHandler.ErrorLogger.Println("User ID wildcard:", contentProvider.UserWildcard)
		logHandler.ErrorLogger.Println("User ID:", currentUserID)

		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	logHandler.InfoLogger.Println("Fetching profile for user ID:", currentUserID)
	// CConvert the ID to an integer if necessary, or handle it as a string
	userNumericID, err := strconv.Atoi(currentUserID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	dl, err := viewProvider.Dashboard(context.TODO(), userNumericID)
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
