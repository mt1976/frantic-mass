package weightProjectionHistory

// Data Access Object weightProjectionHistory
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "weightProjectionHistory" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the weightProjectionHistory_Store struct to match the domain entity
//TODO: Update the FIELD_ constants to match the domain entity

import (
	audit "github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-mass/app/dao/dateIndex"
	"github.com/mt1976/frantic-mass/app/dao/weightProjection"
)

var domain = "weightProjectionHistory"

// WeightProjectionHistory represents a WeightProjectionHistory entity.
type WeightProjectionHistory struct {
	ID    int         `storm:"id,increment=100000"` // primary key with auto increment
	Key   string      `storm:"unique"`              // key
	Raw   string      `storm:"unique"`              // raw ID before encoding
	Audit audit.Audit `csv:"-"`                     // audit data
	// Add your fields here
	DateIndex        dateIndex.DateIndex               `` // Foreign key to DateIndex
	WeightProjection weightProjection.WeightProjection `` // Foreign key to WeightProjection
}

// Define the field set as names
var (
	FIELD_ID    = "ID"
	FIELD_Key   = "Key"
	FIELD_Raw   = "Raw"
	FIELD_Audit = "Audit"
	// Add your fields here
	FIELD_Field1 = "Field1"
	FIELD_Field2 = "Field2"
	FIELD_Field3 = "Field3"
)
