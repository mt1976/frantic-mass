package goal

// Data Access Object goal
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "goal" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the New function to implement the creation of a new domain entity
//TODO: Create any new functions required to support the domain entity

import (
	"context"
	"fmt"
	"time"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/types/measures"
)

// New creates a new Goal instance
func New() Goal {
	return Goal{}
}

// Create creates a new Goal instance in the database
// It takes userID, name, targetWeight, targetDate, lossPerWeekKg, note, and isAverageType as parameters
// It returns the created Goal instance or an error if any occurs
// It also checks if the DAO is ready for operations
// It sets the CompositeID and AverageWeightLoss fields based on the provided parameters
// It calculates the BMI based on the target weight and returns the Goal instance

func Create(ctx context.Context, userID int, name string, targetWeight measures.Weight, targetDate time.Time, lossPerWeekKg measures.Weight, note string, isAverageType bool) (Goal, error) {

	dao.CheckDAOReadyState(domain, audit.CREATE, initialised) // Check the DAO has been initialised, Mandatory.

	//logHandler.InfoLogger.Printf("New %v (%v=%v)", domain, FIELD_ID, field1)
	clock := timing.Start(domain, actions.CREATE.GetCode(), fmt.Sprintf("%v", userID))

	sessionID := idHelpers.GetUUID()

	// Create a new struct
	record := Goal{}
	// Create a composite ID for unique identification of the goal
	record.CompositeID = fmt.Sprintf("%d/%s", userID, sessionID)
	record.Key = idHelpers.Encode(record.CompositeID)
	record.Raw = record.CompositeID
	record.UserID = userID
	record.Name = name
	record.TargetWeight = targetWeight
	record.TargetDate = targetDate
	record.LossPerWeek = lossPerWeekKg
	record.Note = note

	// Get the current basline for the user
	baseline, err := baseline.GetByUserID(userID)
	if err != nil {
		// Log and panic if there is an error retrieving the baseline
		logHandler.ErrorLogger.Panic(commonErrors.WrapDAOReadError(domain, FIELD_UserID, userID, err))
	}
	if baseline.Height.LE(0) {
		logHandler.ErrorLogger.Panicf("No height found for user ID %d, cannot calculate BMI", userID)
	}

	bmiPtr, err := record.TargetBMI.SetBMIFromWeightAndHeight(record.TargetWeight, baseline.Height)
	if err != nil {
		logHandler.ErrorLogger.Panicf("Error calculating BMI for user ID %d: %v", userID, err)
	}
	if bmiPtr != nil {
		record.TargetBMI = *bmiPtr
	}

	record.AverageWeightLoss.Set(isAverageType)

	// Record the create action in the audit data
	auditErr := record.Audit.Action(ctx, audit.CREATE.WithMessage(fmt.Sprintf("New %v created %v", domain, userID)))
	if auditErr != nil {
		// Log and panic if there is an error creating the status instance
		logHandler.ErrorLogger.Panic(commonErrors.WrapDAOUpdateAuditError(domain, record.ID, auditErr))
	}

	// Save the status instance to the database
	writeErr := activeDB.Create(&record)
	if writeErr != nil {
		// Log and panic if there is an error creating the status instance
		logHandler.ErrorLogger.Panic(commonErrors.WrapDAOCreateError(domain, record.ID, writeErr))
		//	panic(writeErr)
	}

	//logHandler.AuditLogger.Printf("[%v] [%v] ID=[%v] Notes[%v]", audit.CREATE, domain, record.ID, fmt.Sprintf("New %v: %v", domain, field1))
	clock.Stop(1)
	return record, nil
}
