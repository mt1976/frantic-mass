package baseline

// Data Access Object baseline
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "baseline" TO THE NAME OF THE DOMAIN ENTITY
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
	"github.com/mt1976/frantic-mass/app/types"
)

// New creates a new Baseline instance
func New() Baseline {
	return Baseline{}
}

// Create creates a new Baseline instance in the database
// It takes userID, heightCm, projectionPeriod, and note as parameters and returns the created Baseline instance or an error if any occurs
// It also checks if the DAO is ready for operations
// It sets the CompositeID and Audit fields based on the provided parameters
// It returns the created Baseline instance or an error if any occurs
// It also records the create action in the audit data and saves the instance to the database
func Create(ctx context.Context, userID int, heightCm types.Height, projectionPeriod int, note string, dateOfBirth time.Time) (Baseline, error) {

	dao.CheckDAOReadyState(domain, audit.CREATE, initialised) // Check the DAO has been initialised, Mandatory.

	//logHandler.InfoLogger.Printf("New %v (%v=%v)", domain, FIELD_ID, field1)
	clock := timing.Start(domain, actions.CREATE.GetCode(), fmt.Sprintf("%v", userID))

	sessionID := idHelpers.GetUUID()

	// Create a new struct
	record := Baseline{}
	record.Key = idHelpers.Encode(sessionID)
	record.Raw = sessionID
	record.Audit = audit.Audit{}
	record.UserID = userID
	record.Height = heightCm
	record.ProjectionPeriod = projectionPeriod
	record.Note = note
	record.DateOfBirth = dateOfBirth

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
