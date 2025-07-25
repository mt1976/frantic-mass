package baseline

// Data Access Object baseline
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "baseline" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Implement the importProcessor function to process the domain entity

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/logHandler"
)

// tempalteImportProcessor is a helper function to create a new entry instance and save it to the database
// It should be customised to suit the specific requirements of the entryination table/DAO.
func tempalteImportProcessor(inOriginal **Baseline) (string, error) {
	//TODO: Build the import processing functionality for the baseline_Store data here
	//
	importedData := **inOriginal

	//	logHandler.ImportLogger.Printf("Importing %v [%v] [%v]", domain, original.Raw, original.Field1)

	//logger.InfoLogger.Printf("ACT: NEW New %v %v %v", tableName, name, entryination)
	// u := Behaviour_Store{}
	// u.Key = idHelpers.Encode(strings.ToUpper(original.Raw))
	// u.Raw = original.Raw
	// u.Text = original.Text
	// // u.Action = original.Action
	// u.Domain = original.Domain

	// importAction := actions.New(original.Action.Name)
	// bh, err := Declare(importAction, domains.Special(original.Domain), original.Text)
	// if err != nil {
	// 	logHandler.ImportLogger.Panicf("Error importing baseline: %v", err.Error())
	// }

	// Return the created entry and nil error
	//logHandler.ImportLogger.Printf("Imported %v [%+v]", domain, importedData)

	//godump.Dump(importedData)

	stringField1 := strconv.Itoa(importedData.UserID) // Assuming UserID is the field to be returned as a string
	if importedData.Height.LE(0) {
		logHandler.ImportLogger.Panicf("Invalid HeightCm for %v: %v", domain, importedData.Height)
		return stringField1, commonErrors.HandleGoValidatorError(fmt.Errorf("HeightCm must be greater than zero (%v)", importedData.Height))
	}
	_, err := Create(context.TODO(), importedData.UserID, importedData.Height, importedData.ProjectionPeriod, importedData.Note, importedData.DateOfBirth)
	if err != nil {
		logHandler.ImportLogger.Panicf("Error importing %v: %v", domain, err.Error())
		return stringField1, err
	}

	return stringField1, nil
}
