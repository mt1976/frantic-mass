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

		err := baseline.ExportRecordsAsCSV()
		if err != nil {
			logHandler.ErrorLogger.Printf("Error exporting baseline records: %v", err)
		}
		if err := dateIndex.ExportRecordsAsCSV(); err != nil {
			logHandler.ErrorLogger.Printf("Error exporting date index records: %v", err)
		}
		if err := goal.ExportRecordsAsCSV(); err != nil {
			logHandler.ErrorLogger.Printf("Error exporting goal records: %v", err)
		}
		if err := tag.ExportRecordsAsCSV(); err != nil {
			logHandler.ErrorLogger.Printf("Error exporting tag records: %v", err)
		}
		if err := user.ExportRecordsAsCSV(); err != nil {
			logHandler.ErrorLogger.Printf("Error exporting user records: %v", err)
		}
		if err := weight.ExportRecordsAsCSV(); err != nil {
			logHandler.ErrorLogger.Printf("Error exporting weight records: %v", err)
		}
		if err := weightProjection.ExportRecordsAsCSV(); err != nil {
			logHandler.ErrorLogger.Printf("Error exporting weight projection records: %v", err)
		}
		if err := weightProjectionHistory.ExportRecordsAsCSV(); err != nil {
			logHandler.ErrorLogger.Printf("Error exporting weight projection history records: %v", err)
		}
		if err := weightTag.ExportRecordsAsCSV(); err != nil {
			logHandler.ErrorLogger.Printf("Error exporting weight tag records: %v", err)
		}

		baseline.Close()
		dateIndex.Close()
		goal.Close()
		tag.Close()
		user.Close()
		weight.Close()
		weightProjection.Close()
		weightProjectionHistory.Close()
		weightTag.Close()

		// Execute shutdown logic in background
		go shutdownFunc()
	}
}
