package handlers

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/viewProvider"
)

func Dummy(w http.ResponseWriter, r *http.Request) {
	// This is a dummy router function

	dl, err := viewProvider.Launcher(context.TODO())
	if err != nil {
		logHandler.ErrorLogger.Println("Error creating DisplayLauncher view:", err)
	} else {
		logHandler.EventLogger.Println("DisplayLauncher view created successfully")
	}

	//godump.Dump(dl)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	// Write a simple response
	w.WriteHeader(dl.Context.HttpStatusCode) // Set the HTTP status code
	_, _ = w.Write([]byte(dl.Context.PageTitle + "\n"))
	_, _ = w.Write([]byte(dl.Context.PageSummary + "\n"))
	_, _ = w.Write([]byte("Welcome to the Display Launcher!\n"))
	_, _ = w.Write([]byte("This is a dummy response for the Display Launcher.\n"))
	_, _ = w.Write([]byte("You can customize this response as per your requirements.\n"))
	logHandler.InfoLogger.Println("Dummy router executed successfully")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	// This is a not found handler function
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusNotFound) // Set the HTTP status code to 404 Not Found
	_, _ = w.Write([]byte("404 Not Found\n"))
	logHandler.ErrorLogger.Println("404 Not Found: The requested resource could not be found.")
	logHandler.ErrorLogger.Printf("Requested URL: %s, Method: %s", r.URL.Path, r.Method)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	// This is a method not allowed handler function
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusMethodNotAllowed) // Set the HTTP status code to 405 Method Not Allowed
	_, _ = w.Write([]byte("405 Method Not Allowed\n"))
	logHandler.ErrorLogger.Println("405 Method Not Allowed: The requested method is not allowed for the requested resource.")
}

// isRequestFromLocalhost checks if the request came from localhost
func isRequestFromLocalhost(r *http.Request) bool {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// fallback to basic IP check
		host = r.RemoteAddr
	}
	ip := net.ParseIP(host)
	return ip != nil && (ip.IsLoopback() || isLocalhostString(host))
}

// covers cases where RemoteAddr might be just "127.0.0.1"
func isLocalhostString(host string) bool {
	return host == "127.0.0.1" || host == "::1" || strings.HasPrefix(host, "localhost")
}
