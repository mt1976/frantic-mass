package goal

// Data Access Object goal
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "goal" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the goal_Store struct to match the domain entity
//TODO: Update the FIELD_ constants to match the domain entity

import (
	"time"

	audit "github.com/mt1976/frantic-core/dao/audit"
	t "github.com/mt1976/frantic-mass/app/types"
)

var domain = "goal"

// goal_Store represents a goal_Store entity.
type goal_Store struct {
	ID    int         `storm:"id,increment=100000"` // primary key with auto increment
	Key   string      `storm:"unique"`              // key
	Raw   string      `storm:"unique"`              // raw ID before encoding
	Audit audit.Audit `csv:"-"`                     // audit data
	// Add your fields here
	UserID        int      `storm:"index"`  // Foreign key to User
	Name          string   `storm:"unique"` // Name of the goal
	TargetWeight  t.Weight // Target weight in kilograms
	TargetDate    time.Time
	LossPerWeek   t.Weight // Expected weight loss per week in kilograms
	Note          string
	CompositeID   string // Composite ID for unique identification of the goal
	NoProjections int    // Projection Period in weeks, used for calculating the target date based on the current weight and loss per week
}

// Define the field set as names
var (
	FIELD_ID    = "ID"
	FIELD_Key   = "Key"
	FIELD_Raw   = "Raw"
	FIELD_Audit = "Audit"
	// Add your fields here
	FIELD_Name          = "Name"
	FIELD_UserID        = "UserID"
	FIELD_TargetWeight  = "TargetWeight"
	FIELD_TargetDate    = "TargetDate"
	FIELD_LossPerWeek   = "LossPerWeek"
	FIELD_Note          = "Note"
	FIELD_CompositeID   = "CompositeID"
	FIELD_NoProjections = "NoProjections"
)
