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

// view.Projections = weightProjection.SortByDateAscending(view.Projections)
// 	view.Projections = weightProjection.FilterByGoalID(view.Projections, goalId)

// SortByDateAscending sorts the projections by date in ascending order
func SortByDateAscending(projections []WeightProjection) []WeightProjection {
	if projections == nil {
		return nil
	}
	sorted := make([]WeightProjection, len(projections))
	copy(sorted, projections)
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i].Date.After(sorted[j].Date) {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	return sorted
}
func FilterByGoalID(projections []WeightProjection, goalId int) []WeightProjection {
	if projections == nil {
		return nil
	}
	filtered := make([]WeightProjection, 0)
	for _, projection := range projections {
		if projection.GoalID == goalId {
			filtered = append(filtered, projection)
		}
	}
	return filtered
}
func FilterByUserID(projections []WeightProjection, userId int) []WeightProjection {
	if projections == nil {
		return nil
	}
	filtered := make([]WeightProjection, 0)
	for _, projection := range projections {
		if projection.UserID == userId {
			filtered = append(filtered, projection)
		}
	}
	return filtered
}
