package weightProjection

// Data Access Object weightProjection
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "weightProjection" TO THE NAME OF THE DOMAIN ENTITY
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
	t "github.com/mt1976/frantic-mass/app/types"
)

func New(ctx context.Context, userID, goalID, projectionNo int, weight, amount t.Weight, date time.Time, note, VsTarget string, toGo t.Weight) (WeightProjection, error) {

	dao.CheckDAOReadyState(domain, audit.CREATE, initialised) // Check the DAO has been initialised, Mandatory.

	//logHandler.InfoLogger.Printf("New %v (%v=%v)", domain, FIELD_ID, field1)
	clock := timing.Start(domain, actions.CREATE.GetCode(), fmt.Sprintf("%v", userID))

	sessionID := idHelpers.GetUUID()

	// Create a new struct
	record := WeightProjection{}
	record.Key = idHelpers.Encode(sessionID)
	record.Raw = sessionID

	record.UserID = userID
	record.GoalID = goalID
	record.ProjectionNo = projectionNo
	record.Weight = weight
	record.Amount = amount
	record.Date = date
	record.Note = note
	record.VsTarget = VsTarget
	record.ToGoal = toGo // Set the total weight loss needed to reach the goal
	// Create a composite ID for unique identification of the projection
	//record.CompositeID = fmt.Sprintf("%d/%d/%d", userID, goalID, projectionNo)
	cid := types.NewCompositeIDFromParts(userID, goalID, projectionNo)
	record.CompositeID = cid.String()

	record.BMI = record.GetBMI() // Calculate BMI based on the projected weight

	// Record the create action in the audit data
	auditErr := record.Audit.Action(ctx, audit.CREATE.WithMessage(fmt.Sprintf("New %v created %v", domain, cid)))
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
