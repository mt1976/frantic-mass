package handlers

import (
	"context"
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/controllers"
)

func Launcher(w http.ResponseWriter, r *http.Request) {
	// This is a dummy router function

	dl, err := controllers.Launcher(context.TODO())
	if err != nil {
		logHandler.ErrorLogger.Println("Error creating DisplayLauncher view:", err)
	} else {
		logHandler.EventLogger.Println("DisplayLauncher view created successfully")
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	// Write a simple response
	w.WriteHeader(dl.Context.HttpStatusCode) // Set the HTTP status code

	executeTemplateResponse(dl, dl.Context, w)

	logHandler.InfoLogger.Println("Dummy router executed successfully")
}
