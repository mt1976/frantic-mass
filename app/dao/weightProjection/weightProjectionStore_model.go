package weightProjection

// Data Access Object weightProjection
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "weightProjection" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the weightProjection_Store struct to match the domain entity
//TODO: Update the FIELD_ constants to match the domain entity

import (
	"time"

	audit "github.com/mt1976/frantic-core/dao/audit"
	t "github.com/mt1976/frantic-mass/app/types"
)

var domain = "weightProjection"

// weightProjection_Store represents a weightProjection_Store entity.
type weightProjection_Store struct {
	ID    int         `storm:"id,increment=100000"` // primary key with auto increment
	Key   string      `storm:"unique"`              // key
	Raw   string      `storm:"unique"`              // raw ID before encoding
	Audit audit.Audit `csv:"-"`                     // audit data
	// Add your fields here
	UserID       int       `storm:"index"` // Foreign key to User
	GoalID       int       `storm:"index"` // Foreign key to Goal
	ProjectionNo int       `storm:"index"` // Projection number, used for tracking multiple projections for the same goal
	Weight       t.Weight  // Projected weight in kilograms
	Date         time.Time `storm:"index"` // Date of the projection
	Note         string    // Additional notes for the projection
	CompositeID  string    `storm:"index"` // Composite ID for unique identification of the projection
}

// Define the field set as names
var (
	FIELD_ID    = "ID"
	FIELD_Key   = "Key"
	FIELD_Raw   = "Raw"
	FIELD_Audit = "Audit"
	// Add your fields here
	FIELD_UserID       = "UserID"
	FIELD_GoalID       = "GoalID"
	FIELD_ProjectionNo = "ProjectionNo"
	FIELD_Weight       = "Weight"
	FIELD_Date         = "Date"
	FIELD_Note         = "Note"
	FIELD_CompositeID  = "CompositeID"
)
