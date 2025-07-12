package functions

import (
	"context"

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

func ClearDown(thisContext context.Context) error {

	err := user.ClearDown(thisContext)
	if err != nil {
		logHandler.ErrorLogger.Println("Error clearing down user records:", err)
	} else {
		logHandler.InfoLogger.Println("User records cleared successfully")
	}

	err = baseline.ClearDown(thisContext)
	if err != nil {
		logHandler.ErrorLogger.Println("Error clearing down baseline records:", err)
	} else {
		logHandler.InfoLogger.Println("Baseline records cleared successfully")
	}

	err = tag.ClearDown(thisContext)
	if err != nil {
		logHandler.ErrorLogger.Println("Error clearing down tag records:", err)
	} else {
		logHandler.InfoLogger.Println("Tag records cleared successfully")
	}

	err = goal.ClearDown(thisContext)
	if err != nil {
		logHandler.ErrorLogger.Println("Error clearing down goal records:", err)
	} else {
		logHandler.InfoLogger.Println("Goal records cleared successfully")
	}

	err = weight.ClearDown(thisContext)
	if err != nil {
		logHandler.ErrorLogger.Println("Error clearing down weight records:", err)
	} else {
		logHandler.InfoLogger.Println("Weight records cleared successfully")
	}

	err = weightProjection.ClearDown(thisContext)
	if err != nil {
		logHandler.ErrorLogger.Println("Error clearing down weight projection records:", err)
	} else {
		logHandler.InfoLogger.Println("Weight projection records cleared successfully")
	}

	err = weightTag.ClearDown(thisContext)
	if err != nil {
		logHandler.ErrorLogger.Println("Error clearing down weight tag records:", err)
	} else {
		logHandler.InfoLogger.Println("Weight tag records cleared successfully")
	}

	err = weightProjectionHistory.ClearDown(thisContext)
	if err != nil {
		logHandler.ErrorLogger.Println("Error clearing down weight projection history records:", err)
	} else {
		logHandler.InfoLogger.Println("Weight projection history records cleared successfully")
	}

	err = dateIndex.ClearDown(thisContext)
	if err != nil {
		logHandler.ErrorLogger.Println("Error clearing down date index records:", err)
	} else {
		logHandler.InfoLogger.Println("Date index records cleared successfully")
	}

	logHandler.InfoLogger.Println("All records cleared successfully")
	logHandler.InfoLogger.Println("End of cleardown function")
	return nil
}
