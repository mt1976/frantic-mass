package weight

// Data Access Object weight
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "weight" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the weight_Store struct to match the domain entity
//TODO: Update the FIELD_ constants to match the domain entity

import (
	"time"

	audit "github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-mass/app/types/measures"
)

var domain = "weight"

// Weight represents a Weight entity.
type Weight struct {
	ID    int         `storm:"id,increment=100"` // primary key with auto increment
	Key   string      `storm:"unique"`           // key
	Raw   string      `storm:"unique"`           // raw ID before encoding
	Audit audit.Audit `csv:"-"`                  // audit data
	// Add your fields here
	UserID     int             `storm:"index"` // User ID of the person who recorded the weight
	RecordedAt time.Time       `storm:"index"`
	Weight     measures.Weight // Weight in kilograms
	BMI        measures.BMI    // Body Mass Index
	Note       string
}

// Define the field set as names
var (
	FIELD_ID    = "ID"
	FIELD_Key   = "Key"
	FIELD_Raw   = "Raw"
	FIELD_Audit = "Audit"
	// Add your fields here
	FIELD_UserID     = "UserID"
	FIELD_RecordedAt = "RecordedAt"
	FIELD_Weight     = "Weight"
	FIELD_Note       = "Note"
)
