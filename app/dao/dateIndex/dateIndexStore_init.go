package dateIndex

// Data Access Object dateIndex
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "dateIndex" TO THE NAME OF THE DOMAIN ENTITY
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
	//activeDB = database.ConnectToNamedDB("dateIndex")
	activeDB = database.Connect()
	initialised = true

	//TODO: Add any initialisation code here

	timing.Stop(1)
	logHandler.EventLogger.Printf("Initialised %v", domain)
}

func IsInitialised() bool {
	return initialised
}

func Close() {
	logHandler.EventLogger.Printf("Closing connection to %v", domain)
	if activeDB != nil {
		activeDB.Disconnect()
	}
	logHandler.EventLogger.Printf("Closed connection to %v", domain)
}
