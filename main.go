package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/tag"
	user "github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weight"
	"github.com/mt1976/frantic-mass/app/dao/weightProjection"
	"github.com/mt1976/frantic-mass/app/dao/weightTag"
	"github.com/mt1976/frantic-mass/app/jobs"
	"github.com/mt1976/frantic-mass/app/types"
)

func main() {
	// This is the main function

	logHandler.InfoLogger.Println("massStore", "Initialise", "Start")
	// Initialise the database connections
	logHandler.InfoLogger.Println("massStore", "Initialise", "Initialising Database Connections")

	logHandler.InfoLogger.Println("massStore", "Initialise", "Initialising User DAO")
	user.Initialise(context.TODO())

	logHandler.InfoLogger.Println("massStore", "Initialise", "Initialising Baseline DAO")
	baseline.Initialise(context.TODO())

	logHandler.InfoLogger.Println("massStore", "Initialise", "Initialising Tag DAO")
	tag.Initialise(context.TODO())

	logHandler.InfoLogger.Println("massStore", "Initialise", "Initialising Goal DAO")
	goal.Initialise(context.TODO())

	logHandler.InfoLogger.Println("massStore", "Initialise", "Initialising Weight DAO")
	weight.Initialise(context.TODO())

	logHandler.InfoLogger.Println("massStore", "Initialise", "Initialising WeightTag DAO")
	weightTag.Initialise(context.TODO())
	logHandler.InfoLogger.Println("massStore", "Initialise", "Initialising WeightProjection DAO")
	weightProjection.Initialise(context.TODO())

	logHandler.InfoLogger.Println("massStore", "Initialise", "Initialised Database Connections")

	// Lets clear down the session db
	initErr := user.ClearDown(context.TODO())
	if initErr != nil {
		logHandler.ErrorLogger.Println(initErr)
	}

	logHandler.InfoLogger.Println("massStore", "Initialise", "Done")

	iniErr := user.ClearDown(context.TODO())
	if iniErr != nil {
		logHandler.ErrorLogger.Println(iniErr)
	}
	randLast := 0
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(10000-1000) + 1000
		logHandler.InfoLogger.Println("randNum:", randNum)

		newRec, newRedErr := user.New(context.TODO(), fmt.Sprintf("massUser%v", randNum), fmt.Sprintf("Test %v", randNum), fmt.Sprintf("massUser%v@example.com", randNum))
		if newRedErr != nil {
			logHandler.ErrorLogger.Println(newRedErr)
		} else {
			logHandler.InfoLogger.Printf("newRec:[%+v]", newRec)
		}
		newRec.ExportRecordAsJSON("test")
		randLast = randNum
	}

	userID := randLast

	user.ExportRecordsAsJSON(time.Now().Format("20060102"))
	user.ExportRecordsAsJSON("")

	lk, lkErr := user.GetDefaultLookup()
	if lkErr != nil {
		logHandler.ErrorLogger.Printf("two:[%+v]", lkErr)
	} else {
		logHandler.InfoLogger.Printf("lk:[%+v]", lk)
	}

	count, cerr := user.Count()
	if cerr != nil {
		logHandler.ErrorLogger.Println(cerr)
	} else {
		logHandler.InfoLogger.Printf("count:[%+v]", count)
	}

	count, cerr = user.CountWhere(user.FIELD_ID, randLast)
	if cerr != nil {
		logHandler.ErrorLogger.Println(cerr)
	} else {
		logHandler.InfoLogger.Printf("count:[%+v]", count)
	}

	count2, cerr2 := user.CountWhere(user.FIELD_Email, "test")
	if cerr2 != nil {
		logHandler.ErrorLogger.Println(cerr2)
	} else {
		logHandler.InfoLogger.Printf("count2:[%+v]", count2)
	}

	count3, cerr3 := user.CountWhere(user.FIELD_ID, "poopoolala")
	if cerr3 != nil {
		logHandler.ErrorLogger.Println(cerr3)
	} else {
		logHandler.InfoLogger.Printf("count3:[%+v]", count3)
	}

	count4, cerr4 := user.CountWhere(user.FIELD_ID, 123)
	if cerr4 != nil {
		logHandler.ErrorLogger.Println(cerr4)
	} else {
		logHandler.InfoLogger.Printf("count4:[%+v]", count4)
	}
	/// FInal tests
	dropErr := user.Drop()
	if dropErr != nil {
		logHandler.ErrorLogger.Println(dropErr)
	} else {
		logHandler.InfoLogger.Printf("drop:[templ]")
	}

	cdrop, cdroperr := user.Count()
	if cdroperr != nil {
		logHandler.ErrorLogger.Println(cdroperr)
	} else {
		logHandler.InfoLogger.Printf("cdrop:[%+v]", cdrop)
	}

	for i := 0; i < 10; i++ {
		randNum := rand.Intn(10000-1000) + 1000
		//logHandler.InfoLogger.Println("randNum:", randNum)

		_, newRedErr := user.New(context.TODO(), fmt.Sprintf("massUser%v", randNum), fmt.Sprintf("Test %v", randNum), fmt.Sprintf("massUser%v@example.com", randNum))
		if newRedErr != nil {
			logHandler.ErrorLogger.Println(newRedErr)
		}
	}

	user.ExportRecordsAsCSV()
	// Drop Again
	dropErr = user.ClearDown(context.TODO())
	if dropErr != nil {
		logHandler.ErrorLogger.Println(dropErr)
	} else {
		logHandler.InfoLogger.Printf("Drop Data Post Export:[templ]")
	}

	user.ImportRecordsFromCSV()

	cdrop, cdroperr = user.Count()
	if cdroperr != nil {
		logHandler.ErrorLogger.Println(cdroperr)
	} else {
		logHandler.InfoLogger.Printf("Post Import:[%+v]", cdrop)
	}

	logHandler.InfoLogger.Println("massStore", "Initialise", "Creating Baseline for UserID:", userID)
	thisBaseline, baselineErr := baseline.New(context.TODO(), userID, types.Height{Value: 187.96}, fmt.Sprintf("BaselineFor%v", userID))
	if baselineErr != nil {
		logHandler.ErrorLogger.Println(baselineErr)
	} else {
		logHandler.InfoLogger.Printf("Baseline Created:[%+v]", thisBaseline)
	}

	logHandler.InfoLogger.Println("massStore", "Initialise", "Creating Goal for UserID:", userID)
	thisGoal, goalErr := goal.New(context.TODO(), userID, fmt.Sprintf("GoalFor%v", userID), types.Weight{Value: 90.00}, time.Now().AddDate(0, 0, 30), types.Weight{Value: 2.0}, "This is a test goal to check the goal creation process")

	if goalErr != nil {
		logHandler.ErrorLogger.Println(goalErr)
	} else {
		logHandler.InfoLogger.Printf("Goal Created:[%+v]", thisGoal)
	}

	logHandler.InfoLogger.Println("massStore", "Initialise", "Creating Tag")
	thisTag, tagErr := tag.New(context.TODO(), fmt.Sprintf("TagFor%v", userID))
	if tagErr != nil {
		logHandler.ErrorLogger.Println(tagErr)
	} else {
		logHandler.InfoLogger.Printf("Tag Created:[%+v]", thisTag)
	}

	logHandler.InfoLogger.Println("massStore", "Initialise", "Creating Weight for UserID:", userID)
	thisWeight, weightErr := weight.New(context.TODO(), userID, types.Weight{Value: 120.00}, fmt.Sprintf("WeightFor%v", userID), time.Now())
	if weightErr != nil {
		logHandler.ErrorLogger.Println(weightErr)
	} else {
		logHandler.InfoLogger.Printf("Weight Created:[%+v]", thisWeight)
	}

	for i := 0; i < 10; i++ {
		randNum := (120.00 - i) - rand.Intn(10-1) + 1
		logHandler.InfoLogger.Println("randNum:", randNum)

		_, newRedErr := weight.New(context.TODO(), userID, types.Weight{Value: float64(randNum)}, fmt.Sprintf("WeightFor%v", randNum), time.Now().Add(time.Duration(randNum)*(time.Hour*24)))
		if newRedErr != nil {
			logHandler.ErrorLogger.Println(newRedErr)
		}

		for j := 1; j < 11; j++ {

			np, newProjectionErr := weightProjection.New(context.TODO(), userID, thisGoal.ID, j, types.Weight{Value: float64(randNum + j)}, time.Now().Add(time.Duration(randNum+j)*(time.Hour*24)), fmt.Sprintf("ProjectionFor%v-%v", randNum, j))
			if newProjectionErr != nil {
				logHandler.ErrorLogger.Println(newProjectionErr)
			} else {
				logHandler.InfoLogger.Printf("Projection Created:[%v]", np.CompositeID)
			}

		}
	}
	logHandler.InfoLogger.Println("massStore", "Initialise", "Creating WeightTag")
	thisWeightTag, weightTagErr := weightTag.New(context.TODO(), thisWeight.ID, thisTag.ID)
	if weightTagErr != nil {
		logHandler.ErrorLogger.Println(weightTagErr)
	} else {
		logHandler.InfoLogger.Printf("WeightTag Created:[%+v]", thisWeightTag)
	}

	var userJobInstance jobs.UserJob = jobs.UserJob{}

	userJobInstance.AddDatabaseAccessFunctions(user.GetDatabaseConnections())
	userJobInstance.AddDatabaseAccessFunctions(user.GetDatabaseConnections())
	// Lets check the job processing
	err := userJobInstance.Run()
	if err != nil {
		logHandler.ErrorLogger.Println(err)
	}

	logHandler.InfoLogger.Println("massStore", "HeightCM", thisBaseline.HeightCm.CmAsString())
	logHandler.InfoLogger.Println("massStore", "HeightM", thisBaseline.HeightCm.MetresAsString())

	logHandler.InfoLogger.Println("massStore", "HeightIN", thisBaseline.HeightCm.InchesAsString())
	logHandler.InfoLogger.Println("massStore", "HeightFT", thisBaseline.HeightCm.FeetAsString())

	logHandler.InfoLogger.Println("massStore", "HeightString", thisBaseline.HeightCm.String())

	logHandler.InfoLogger.Println("massStore", "WeightKg", thisWeight.Weight.KgAsString())
	logHandler.InfoLogger.Println("massStore", "WeightLbs", thisWeight.Weight.LbsAsString())
	stones, _ := thisWeight.Weight.StonesAsString()
	logHandler.InfoLogger.Println("massStore", "WeightStone", stones)
	logHandler.InfoLogger.Println("massStore", "WeightString", thisWeight.Weight.String())

	loss, lossErr := weight.AverageWeightLoss(userID)
	if lossErr != nil {
		logHandler.ErrorLogger.Println(lossErr)
	} else {
		logHandler.InfoLogger.Printf("Average Weight Loss for User %d: %s", userID, loss.KgAsString())
	}
	logHandler.InfoLogger.Println("massStore", "Initialise", "End")

}
