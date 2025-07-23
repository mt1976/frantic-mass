package functions

import (
	"fmt"
	"time"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/weight"
	"github.com/mt1976/frantic-mass/app/types"
)

func AverageWeightLoss(userID int) (avg *types.Weight, tot *types.Weight, err error) {
	// This function will calculate the average weight loss for a user
	// based on their historical weight data.
	// It will return the average weight loss in kilograms.
	// If no data is found, it will return an error.

	// Get The starting weight for the user, loop through the user's weight history in order to calculate the average weight loss

	w, err := weight.GetAllWhere(weight.FIELD_UserID, userID)
	if err != nil {
		logHandler.ErrorLogger.Printf("Error retrieving weight data for user %d: %v", userID, err)
		return &types.Weight{}, &types.Weight{}, err
	}
	if len(w) == 0 {
		logHandler.ErrorLogger.Printf("No weight data found for user %d", userID)
		return &types.Weight{}, &types.Weight{}, nil
	}

	w = weight.FilterDeletedRecords(w)
	w = weight.SortByDateAscending(w)
	//var totalLoss *types.Weight
	//var totalMass *types.Weight
	//totalLoss := types.NewWeight(0.0)
	totalMass := types.NewWeight(0.0)
	var count int
	startingWeight := types.NewWeight(w[0].Weight.KGs)
	endingWeight := types.NewWeight(w[len(w)-1].Weight.KGs)
	for i := 0; i < len(w)-1; i++ {
		if i == 0 {
			startingWeight.Set(w[i].Weight.KGs)
		}
		if w[i].Weight.GT(w[i+1].Weight.KGs) {
			//loss := w[i].Weight.Minus(w[i+1].Weight)
			//totalLoss.Add(loss)
			count++
		}
		logHandler.InfoLogger.Printf("Weight record %d: %v", i, w[i].Weight.String())
		logHandler.InfoLogger.Printf("Weight recorded %d: %v", i+1, w[i+1].Audit.CreatedAt.String())
		totalMass.Add(w[i].Weight)
	}
	logHandler.InfoLogger.Printf("Starting weight for user %d: %v", userID, startingWeight.String())
	logHandler.InfoLogger.Printf("Ending weight for user %d: %v", userID, endingWeight.String())

	totalLoss := endingWeight.MinusFloat(startingWeight.KGs)
	totalLoss = *totalLoss.Invert()
	//godump.Dump(totalLoss)
	totalLossString := totalLoss.KgAsString()
	logHandler.InfoLogger.Printf("Total weight loss for user %d: %v over %d records", userID, totalLossString, count)
	//logHandler.InfoLogger.Printf("Total weight mass for user %d: %v over %d records", userID, totalMass.String(), count)
	logHandler.InfoLogger.Printf("Total weight loss for user %d: %v over %d records", userID, totalLoss.KgAsString(), count)
	logHandler.InfoLogger.Printf("Average weight loss for user %d: %v kg", userID, totalLoss.Kg()/float64(count))
	stas, _ := totalLoss.StonesAsString()
	logHandler.InfoLogger.Printf("Total weight loss for user %d: %v", userID, stas)
	avg = types.NewWeight(totalLoss.Kg() / float64(count))
	//os.Exit(0) // Exit the program successfully
	return avg, &totalLoss, nil
}

func FetchLatestWeightRecord(userID int) (weight.Weight, error) {
	// This function will return the latest weight record for a user
	// If no records are found, it will return an error

	w, err := weight.GetAllWhere(weight.FIELD_UserID, userID)
	if err != nil {
		logHandler.ErrorLogger.Printf("Error retrieving weight data for user %d: %v", userID, err)
		return weight.Weight{}, err
	}
	if len(w) == 0 {
		logHandler.ErrorLogger.Printf("No weight data found for user %d", userID)
		return weight.Weight{}, fmt.Errorf("no weight data found for user %d", userID)
	}

	// Loop through the weight records to find the latest one
	var latest weight.Weight
	for _, record := range w {

		if record.RecordedAt.After(time.Now()) {
			logHandler.WarningLogger.Printf("Future weight record found for user %d: %v", userID, record.RecordedAt.String())
			continue // Skip future records
		}

		if record.UserID != userID {
			logHandler.WarningLogger.Printf("Weight record for user %d does not match requested user %d %d", record.UserID, userID)
			continue // Skip records for other users
		}

		if latest.ID == 0 || record.RecordedAt.After(latest.RecordedAt) {
			latest = record
		}

	}
	logHandler.InfoLogger.Printf("Latest weight record for user %d: %v", userID, latest.String())
	if latest.ID == 0 {
		return weight.Weight{}, fmt.Errorf("no weight records found for user %d", userID)
	}

	return latest, nil
}
