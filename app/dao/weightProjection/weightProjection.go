package weightProjection

import (
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	t "github.com/mt1976/frantic-mass/app/types"
)

func (record *WeightProjection) GetBMI() t.BMI {
	logHandler.EventLogger.Printf("GetBMI called for weightProjection_Store ID %d", record.ID)
	if record == nil {
		return t.BMI{}
	}
	// Recalculate BMI if it is not set
	if record.BMI.IsZero() {
		// Get the baseline height for the user
		bl, err := baseline.GetByUserID(record.UserID)
		if err != nil {
			logHandler.ErrorLogger.Printf("Error retrieving baseline height for user %d: %v", record.UserID, err)
			return t.BMI{}
		}
		var errBMI error
		bmiPtr, errBMI := record.BMI.SetBMIFromWeightAndHeight(record.Weight, bl.Height) // Calculate BMI based on the weight and user ID
		if errBMI != nil || bmiPtr == nil || bmiPtr.IsZero() {
			// If BMI calculation failed or is still zero, log an error and return an empty BMI
			logHandler.ErrorLogger.Printf("Error calculating BMI for user %d with weight %v and height %v: %v", record.UserID, record.Weight, bl.Height, errBMI)
			return t.BMI{}
		}
		record.BMI = *bmiPtr
		logHandler.EventLogger.Printf("Calculated BMI for user %d: %v", record.UserID, record.BMI)
	}
	return record.BMI
}
