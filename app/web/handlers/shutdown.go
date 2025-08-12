package handlers

import (
	"net/http"

	"github.com/mt1976/frantic-mass/app/dao/tag"
)

// ShutdownHandler returns a handler function with injected shutdown logic
func ShutdownHandler(shutdownFunc func()) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isRequestFromLocalhost(r) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		w.Write([]byte("Shutting down server...\n"))

		tag.Close()

		// Execute shutdown logic in background
		go shutdownFunc()
	}
}
