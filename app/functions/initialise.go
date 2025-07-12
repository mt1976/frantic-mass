package functions

import (
	"context"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/dateIndex"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/tag"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weight"
	"github.com/mt1976/frantic-mass/app/dao/weightProjection"
	"github.com/mt1976/frantic-mass/app/dao/weightProjectionHistory"
	"github.com/mt1976/frantic-mass/app/dao/weightTag"
)

func Initialise(thisContext context.Context) error {
	logHandler.InfoLogger.Println("Initialising User DAO")

	user.Initialise(thisContext)

	logHandler.InfoLogger.Println("Initialising Baseline DAO")
	baseline.Initialise(thisContext)

	logHandler.InfoLogger.Println("Initialising Tag DAO")
	tag.Initialise(thisContext)

	logHandler.InfoLogger.Println("Initialising Goal DAO")
	goal.Initialise(thisContext)

	logHandler.InfoLogger.Println("Initialising Weight DAO")
	weight.Initialise(thisContext)

	logHandler.InfoLogger.Println("Initialising WeightTag DAO")
	weightTag.Initialise(thisContext)

	logHandler.InfoLogger.Println("Initialising WeightProjection DAO")
	weightProjection.Initialise(thisContext)

	logHandler.InfoLogger.Println("Initialising WeightProjectionHistory DAO")
	weightProjectionHistory.Initialise(thisContext)

	logHandler.InfoLogger.Println("Initialising DateIndex DAO")
	dateIndex.Initialise(thisContext)

	logHandler.InfoLogger.Println("Initialised all DAOs")
	logHandler.InfoLogger.Println("Initialised Functions")

	logHandler.InfoLogger.Println("Initialised Database Connections")
	return nil
}
