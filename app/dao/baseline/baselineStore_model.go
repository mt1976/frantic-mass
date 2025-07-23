package baseline

// Data Access Object baseline
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "baseline" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the baseline_Store struct to match the domain entity
//TODO: Update the FIELD_ constants to match the domain entity

import (
	"time"

	audit "github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-mass/app/types"
)

var domain = "baseline"

// Baseline represents a Baseline entity.
type Baseline struct {
	ID    int         `storm:"id,increment=100"` // primary key with auto increment
	Key   string      `storm:"unique"`           // key
	Raw   string      `storm:"unique"`           // raw ID before encoding
	Audit audit.Audit `csv:"-"`                  // audit data
	// Add your fields here
	UserID           int          `storm:"index"` // Foreign key
	Height           types.Height // Height in centimeters
	ProjectionPeriod int          // Projection period in months
	DateOfBirth      time.Time    // Date of birth
	PivotDate        time.Time    // Pivot date for projections
	Note             string
}

// Define the field set as names
var (
	FIELD_ID    = "ID"
	FIELD_Key   = "Key"
	FIELD_Raw   = "Raw"
	FIELD_Audit = "Audit"
	// Add your fields here
	FIELD_UserID           = "UserID"
	FIELD_Height           = "Height"
	FIELD_ProjectionPeriod = "ProjectionPeriod"
	FIELD_DateOfBirth      = "DateOfBirth"
	FIELD_Note             = "Note"
)
