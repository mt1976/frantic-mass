package user

import (
	"fmt"

	w "github.com/mt1976/frantic-mass/app/dao/weight"
	"github.com/mt1976/frantic-mass/app/types"
)

var Nil = User{}

func (record *User) StartingWeight() (types.Weight, error) {
	// This function returns the starting weight of the user as an appropriate string.

	userID := record.ID
	if userID <= 0 {
		return types.Weight{}, fmt.Errorf("invalid user ID: %d", userID)
	}
	weightRecords, err := w.GetAllWhere(FIELD_ID, userID)
	if err != nil {
		return types.Weight{}, err
	}
	if len(weightRecords) == 0 {
		return types.Weight{}, fmt.Errorf("No weight records found for user ID %d", userID)
	}
	// Find earliest weight record
	var earliestRecord *w.Weight
	for _, wr := range weightRecords {
		if earliestRecord == nil || wr.RecordedAt.Before(earliestRecord.RecordedAt) {
			earliestRecord = &wr
		}
	}
	if earliestRecord == nil {
		return types.Weight{}, fmt.Errorf("No valid weight records found for user ID %d", userID)
	}

	return earliestRecord.Weight, nil
}
