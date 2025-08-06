package goal

// Data Access Object goal
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "goal" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the goal_Store struct to match the domain entity
//TODO: Update the FIELD_ constants to match the domain entity

import (
	"time"

	"github.com/mt1976/frantic-core/dao"
	audit "github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-mass/app/types/measures"
)

var domain = "goal"

// Goal represents a Goal entity.
type Goal struct {
	ID    int         `storm:"id,increment=100"` // primary key with auto increment
	Key   string      `storm:"unique"`           // key
	Raw   string      `storm:"unique"`           // raw ID before encoding
	Audit audit.Audit `csv:"-"`                  // audit data
	// Add your fields here
	UserID            int             `storm:"index"`  // Foreign key to User
	Name              string          `storm:"unique"` // Name of the goal
	TargetWeight      measures.Weight // Target weight in kilograms
	TargetBMI         measures.BMI    // Target BMI
	TargetDate        time.Time       // Target date for achieving the goal
	LossPerWeek       measures.Weight // Desired weight loss per week in kilograms
	Note              string
	Description       string
	CompositeID       string        // Composite ID for unique identification of the goal
	NoProjections     int           // Projection Period in weeks, used for calculating the target date based on the current weight and loss per week
	AverageWeightLoss dao.StormBool // Type of goal, e.g., user-defined or average weight loss goal
}

// Define the field set as names
var (
	FIELD_ID    = "ID"
	FIELD_Key   = "Key"
	FIELD_Raw   = "Raw"
	FIELD_Audit = "Audit"
	// Add your fields here
	FIELD_Name              = "Name"
	FIELD_UserID            = "UserID"
	FIELD_TargetWeight      = "TargetWeight"
	FIELD_TargetBMI         = "TargetBMI"
	FIELD_TargetDate        = "TargetDate"
	FIELD_LossPerWeek       = "LossPerWeek"
	FIELD_Note              = "Note"
	FIELD_CompositeID       = "CompositeID"
	FIELD_NoProjections     = "NoProjections"
	FIELD_AverageWeightLoss = "AverageWeightLoss"
)
