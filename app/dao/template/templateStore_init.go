package template

// Data Access Object Template
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "Template" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Add any initialisation code to the Initialise function

import (
	"context"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var sessionExpiry = 20 // default to 20 mins
var activeDB *database.DB
var initialised bool = false // default to false
var cfg *commonConfig.Settings

func Initialise(ctx context.Context) {
	timing := timing.Start(domain, actions.INITIALISE.GetCode(), "Initialise")
	cfg = commonConfig.Get()
	// For a specific database connection, use NamedConnect, otherwise use Connect
	activeDB = database.ConnectToNamedDB("Template")
	// activeDB = database.Connect()
	initialised = true

	//TODO: Add any initialisation code here

	timing.Stop(1)
	logHandler.EventLogger.Printf("Initialised %v", domain)
}
