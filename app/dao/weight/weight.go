package weight

import (
	"fmt"

	"github.com/asdine/storm/index"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/types"
)

func (record *Weight) String() string {
	if record == nil {
		return "Weight_Store is nil"
	}
	return "Weight_Store{" +
		"ID: " + fmt.Sprintf("%d", record.ID) +
		", Key: " + record.Key +
		", Raw: " + record.Raw +
		", UserID: " + fmt.Sprintf("%d", record.UserID) +
		", RecordedAt: " + record.RecordedAt.String() +
		", Weight: " + record.Weight.String() +
		", BMI: " + record.BMI.String() +
		", Note: " + record.Note +
		"}"
}

func (record *Weight) GetBMI() types.BMI {
	logHandler.EventLogger.Printf("GetBMI called for Weight_Store ID %d", record.ID)
	if record == nil {
		return types.BMI{}
	}
	// Recalculate BMI if it is not set
	if record.BMI.IsZero() {
		// Get the baseline height for the user
		bl, err := baseline.GetByUserID(record.UserID)
		if err != nil {
			logHandler.ErrorLogger.Printf("Error retrieving baseline height for user %d: %v", record.UserID, err)
			return types.BMI{}
		}
		if bl.Height.LE(0) {
			logHandler.ErrorLogger.Printf("No height found for user ID %d, cannot calculate BMI", record.UserID)
			return types.BMI{}
		}
		logHandler.ErrorLogger.Printf("BMI is not set for weight record ID %d, recalculating...", record.ID)
		// Assuming we have a method to calculate BMI based on weight and height
		// This is a placeholder, actual implementation may vary
		record.BMI.SetBMIFromWeightAndHeight(record.Weight, bl.Height) // Example height in cm
	}
	return record.BMI
}

func GetComplex(...func(*index.Options)) ([]Weight, error) {

	dao.CheckDAOReadyState(domain, audit.GET, initialised) // Check the DAO has been initialised, Mandatory.

	recordList := []Weight{}

	clock := timing.Start(domain, actions.GETALL.GetCode(), "ALL")

	if errG := activeDB.GetAll(&recordList); errG != nil {
		clock.Stop(0)
		return []Weight{}, commonErrors.WrapNotFoundError(domain, errG)
	}

	var errPost error
	if recordList, errPost = postGetList(&recordList); errPost != nil {
		clock.Stop(0)
		return nil, errPost
	}

	clock.Stop(len(recordList))

	return recordList, nil
}
