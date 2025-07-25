package dateIndex

// Data Access Object dateIndex
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "dateIndex" TO THE NAME OF THE DOMAIN ENTITY
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
)

// New creates a new DateIndex instance
func New(date time.Time) DateIndex {
	// Set the date to the start of the day to avoid time issues
	di := DateIndex{}
	di.Date = date.Truncate(24 * time.Hour) // Truncate to the start of the day
	// if date = today, then set the current flag to true
	if di.Date.Equal(time.Now().Truncate(24 * time.Hour)) {
		di.Current.Set(true) // Set the current flag to true if the date is today
	} else {
		di.Current.Set(false) // Default to false for other dates
	}
	return di
}

// Create creates a new DateIndex instance in the database
// It takes a date as a parameter and returns the created DateIndex instance or an error if any occurs
// It also checks if the DAO is ready for operations
// It sets the date to the start of the day to avoid time issues
// It records the create action in the audit data and saves the instance to the database
// It returns the created DateIndex instance or an error if any occurs

func Create(ctx context.Context, date time.Time) (DateIndex, error) {

	dao.CheckDAOReadyState(domain, audit.CREATE, initialised) // Check the DAO has been initialised, Mandatory.

	//logHandler.InfoLogger.Printf("New %v (%v=%v)", domain, FIELD_ID, field1)
	clock := timing.Start(domain, actions.CREATE.GetCode(), fmt.Sprintf("%v", date))

	sessionID := idHelpers.GetUUID()

	// Set the date to the start of the day to avoid time issues
	date = date.Truncate(24 * time.Hour) // Truncate to the start of the day

	// Create a new struct
	record := New(date)
	record.Key = idHelpers.Encode(sessionID)
	record.Raw = sessionID
	updatedRecord, _, err := classifyDateIndexRecord(&record)
	if err != nil {
		logHandler.ErrorLogger.Panic(commonErrors.WrapDAOCreateError(domain, updatedRecord.ID, err))
	}
	record = *updatedRecord

	// Record the create action in the audit data
	auditErr := record.Audit.Action(ctx, audit.CREATE.WithMessage(fmt.Sprintf("New %v created %v", domain, date)))
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
