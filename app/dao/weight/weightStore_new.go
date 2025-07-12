package weight

// Data Access Object weight
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "weight" TO THE NAME OF THE DOMAIN ENTITY
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
	"github.com/mt1976/frantic-mass/app/types"
)

func New(ctx context.Context, userID int, weightKg types.Weight, note string, recordedAt time.Time) (Weight_Store, error) {

	dao.CheckDAOReadyState(domain, audit.CREATE, initialised) // Check the DAO has been initialised, Mandatory.

	//logHandler.InfoLogger.Printf("New %v (%v=%v)", domain, FIELD_ID, userID)
	clock := timing.Start(domain, actions.CREATE.GetCode(), fmt.Sprintf("%v", userID))

	sessionID := idHelpers.GetUUID()

	// Create a new struct
	record := Weight_Store{}
	record.Key = idHelpers.Encode(sessionID)
	record.Raw = sessionID
	record.Audit = audit.Audit{}
	record.UserID = userID
	// Set the recordedAt field to the provided time
	if recordedAt.IsZero() {
		// If the recordedAt is zero, use the current time
		recordedAt = time.Now()
	}
	record.RecordedAt = recordedAt // Convert float64 to time.Time
	record.Weight = weightKg
	record.Note = note

	// Calculate BMI based on the weight and user ID
	// Get Users Baseline Height
	bl, err := baseline.GetByUserID(record.UserID)
	if err != nil {
		// Log and panic if there is an error retrieving the height
		logHandler.ErrorLogger.Panic(commonErrors.WrapDAOReadError(domain, FIELD_UserID, userID, err))
	}
	if bl.Height.LE(0) {
		// If no height is found, set a default value or handle the error as needed
		logHandler.ErrorLogger.Panic(fmt.Errorf("no height found for user ID %d", userID))
	}

	BMI := types.BMI{}
	xx, err := BMI.SetBMIFromWeightAndHeight(weightKg, bl.Height) // Calculate BMI based on the weight and user ID
	if err != nil {
		// Log and panic if there is an error calculating the BMI
		logHandler.ErrorLogger.Panic(commonErrors.WrapDAOCreateError(domain, userID, err))
	}
	record.BMI = *xx // Set the calculated BMI

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
