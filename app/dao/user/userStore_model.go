package user

// Data Access Object User
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "User" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the User_Store struct to match the domain entity
//TODO: Update the FIELD_ constants to match the domain entity

import (
	audit "github.com/mt1976/frantic-core/dao/audit"
)

var domain = "User"

// User represents a User entity.
type User struct {
	ID    int         `storm:"id,increment=100"` // primary key with auto increment
	Key   string      `storm:"unique"`           // key
	Raw   string      `storm:"unique"`           // raw ID before encoding
	Audit audit.Audit `csv:"-"`                  // audit data
	// Add your fields here
	Username          string `storm:"unique"`
	Email             string `storm:"unique"`
	PasswordHash      string
	Name              string `storm:"index"` // name for indexing
	MeasurementSystem int
}

// Define the field set as names
var (
	FIELD_ID    = "ID"
	FIELD_Key   = "Key"
	FIELD_Raw   = "Raw"
	FIELD_Audit = "Audit"
	// Add your fields here
	FIELD_Username          = "Username"
	FIELD_Email             = "Email"
	FIELD_PasswordHash      = "PasswordHash"
	FIELD_Name              = "Name"
	FIELD_MeasurementSystem = "MeasurementSystem"
)
