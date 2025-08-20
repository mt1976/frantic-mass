package handlers

import (
	"context"
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func Error(w http.ResponseWriter, r *http.Request) {
	// This is a dummy router function
	userID := getURLParamValue(r, contentProvider.UserWildcard) // Get the user ID from the URL parameter

	message := getURLParamValue(r, contentProvider.ErrorMessageWildcard)
	description := getURLParamValue(r, contentProvider.ErrorDescriptionWildcard)
	code := getURLParamValue(r, contentProvider.ErrorCodeWildcard)

	if userID == "" {
		userID = "0"
	}
	userID_int, err := contentProvider.StringToInt(userID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error converting userID to int:", err)
		return
	}

	if message == "" {
		logHandler.ErrorLogger.Println("Error message is missing")
		message = "An unexpected error occurred"
	}

	if description == "" {
		logHandler.WarningLogger.Println("Error description is missing")
		description = "No further details are available"
	}

	if code == "" {
		logHandler.WarningLogger.Println("Error code is missing")
		code = "UNKNOWN_ERROR"
	}

	logHandler.InfoLogger.Println("Paths from URL:", r.URL.Path)
	dl, err := viewProvider.Error(context.TODO(), userID_int, message, description, code)
	if err != nil {
		logHandler.ErrorLogger.Println("Error creating DisplayError view:", err)
	} else {
		logHandler.EventLogger.Println("DisplayError view created successfully")
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
