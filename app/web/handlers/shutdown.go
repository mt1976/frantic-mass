package handlers

import (
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/dateIndex"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/tag"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weight"
	"github.com/mt1976/frantic-mass/app/dao/weightProjection"
	"github.com/mt1976/frantic-mass/app/dao/weightProjectionHistory"
	"github.com/mt1976/frantic-mass/app/dao/weightTag"
)

// ShutdownHandler returns a handler function with injected shutdown logic
func ShutdownHandler(shutdownFunc func()) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isRequestFromLocalhost(r) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		w.Write([]byte("Shutting down server...\n"))
		Shutdown()

		// Execute shutdown logic in background
		go shutdownFunc()
	}
}

func Shutdown() {
	logHandler.WarningLogger.Println("Shutting down application...")
	message := "shutdown"
	baseline.ExportRecordsAsJSON(message)
	dateIndex.ExportRecordsAsJSON(message)
	goal.ExportRecordsAsJSON(message)
	tag.ExportRecordsAsJSON(message)
	user.ExportRecordsAsJSON(message)
	weight.ExportRecordsAsJSON(message)
	weightProjection.ExportRecordsAsJSON(message)
	weightProjectionHistory.ExportRecordsAsJSON(message)
	weightTag.ExportRecordsAsJSON(message)

	baseline.Close()
	dateIndex.Close()
	goal.Close()
	tag.Close()
	user.Close()
	weight.Close()
	weightProjection.Close()
	weightProjectionHistory.Close()
	weightTag.Close()
	logHandler.EventLogger.Println("Application shutdown complete.")
}
