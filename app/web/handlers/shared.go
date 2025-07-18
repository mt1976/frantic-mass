package handlers

import (
	"context"
	"net/http"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/controllers"
)

func Dummy(w http.ResponseWriter, r *http.Request) {
	// This is a dummy router function

	dl, err := controllers.Launcher(context.TODO())
	if err != nil {
		logHandler.ErrorLogger.Println("Error creating DisplayLauncher view:", err)
	} else {
		logHandler.EventLogger.Println("DisplayLauncher view created successfully")
	}

	godump.Dump(dl)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	// Write a simple response
	w.WriteHeader(dl.Common.HttpStatusCode) // Set the HTTP status code
	w.Write([]byte(dl.Common.PageTitle + "\n"))
	w.Write([]byte(dl.Common.PageSummary + "\n"))
	w.Write([]byte("Welcome to the Display Launcher!\n"))
	w.Write([]byte("This is a dummy response for the Display Launcher.\n"))
	w.Write([]byte("You can customize this response as per your requirements.\n"))
	logHandler.InfoLogger.Println("Dummy router executed successfully")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	// This is a not found handler function
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusNotFound) // Set the HTTP status code to 404 Not Found
	w.Write([]byte("404 Not Found\n"))
	logHandler.ErrorLogger.Println("404 Not Found: The requested resource could not be found.")
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	// This is a method not allowed handler function
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusMethodNotAllowed) // Set the HTTP status code to 405 Method Not Allowed
	w.Write([]byte("405 Method Not Allowed\n"))
	logHandler.ErrorLogger.Println("405 Method Not Allowed: The requested method is not allowed for the requested resource.")
}
