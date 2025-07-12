package functions

import (
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

func ExportData() error {

	err := weight.ExportRecordsAsCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error exporting weight records as CSV:", err)
	} else {
		logHandler.InfoLogger.Println("Weight records exported successfully as CSV")
	}

	err = user.ExportRecordsAsCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error exporting user records as CSV:", err)
	} else {
		logHandler.InfoLogger.Println("User records exported successfully as CSV")
	}

	err = baseline.ExportRecordsAsCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error exporting baseline records as CSV:", err)
	} else {
		logHandler.InfoLogger.Println("Baseline records exported successfully as CSV")
	}

	err = goal.ExportRecordsAsCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error exporting goal records as CSV:", err)
	} else {
		logHandler.InfoLogger.Println("Goal records exported successfully as CSV")
	}

	err = tag.ExportRecordsAsCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error exporting tag records as CSV:", err)
	} else {
		logHandler.InfoLogger.Println("Tag records exported successfully as CSV")
	}

	err = weightProjection.ExportRecordsAsCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error exporting weight projection records as CSV:", err)
	} else {
		logHandler.InfoLogger.Println("Weight projection records exported successfully as CSV")
	}

	err = weightTag.ExportRecordsAsCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error exporting weight tag records as CSV:", err)
	} else {
		logHandler.InfoLogger.Println("Weight tag records exported successfully as CSV")
	}
	err = weightProjectionHistory.ExportRecordsAsCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error exporting weight projection history records as CSV:", err)
	} else {
		logHandler.InfoLogger.Println("Weight projection history records exported successfully as CSV")
	}
	err = dateIndex.ExportRecordsAsCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error exporting date index records as CSV:", err)
	} else {
		logHandler.InfoLogger.Println("Date index records exported successfully as CSV")
	}

	logHandler.InfoLogger.Println("Data export completed successfully")
	return nil
}
