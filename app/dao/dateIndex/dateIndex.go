package dateIndex

import (
	"fmt"
	"time"

	"github.com/mt1976/frantic-core/logHandler"
)

func Today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour)
}

func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour)
}

func GetToday() (int, DateIndex, error) {
	logHandler.ServiceLogger.Println("Getting today's DateIndex record")
	today := Today()
	record, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("Error getting today's DateIndex record: %v", err)
		return 0, DateIndex{}, err
	}
	for _, r := range record {
		if r.Date.Equal(today) && r.Current.IsTrue() {
			logHandler.ServiceLogger.Printf("Today's DateIndex record found: %v", r)
			return r.ID, r, nil
		}
	}
	logHandler.ServiceLogger.Printf("No DateIndex record found for today: %v", today)
	return 0, DateIndex{}, fmt.Errorf("No DateIndex record found for today: %v", today)
}

func GetYesterday() (int, DateIndex, error) {
	logHandler.ServiceLogger.Println("Getting yesterday's DateIndex record")
	yesterday := Yesterday()
	record, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("Error getting yesterday's DateIndex record: %v", err)
		return 0, DateIndex{}, err
	}
	for _, r := range record {
		if r.Date.Equal(yesterday) && r.Current.IsFalse() {
			logHandler.ServiceLogger.Printf("Yesterday's DateIndex record found: %v", r)
			return r.ID, r, nil
		}
	}
	logHandler.ServiceLogger.Printf("No DateIndex record found for yesterday: %v", yesterday)
	return 0, DateIndex{}, fmt.Errorf("No DateIndex record found for yesterday: %v", yesterday)
}

func GetTomorrow() (int, DateIndex, error) {
	logHandler.PanicLogger.Println("Getting tomorrow's DateIndex record,why? - this should not be used")
	return 0, DateIndex{}, fmt.Errorf("No DateIndex record possibe for tomorrow: %v", Tomorrow())
}

func classifyDateIndexRecord(dateIndexRecord *DateIndex) (*DateIndex, bool, error) {

	logHandler.ServiceLogger.Printf("Categorizing dateIndex record: %v", dateIndexRecord.Date)
	// Check if the date is valid and categorize the record
	// This function checks the date and sets the Current field accordingly
	// It returns the updated record, a boolean indicating if it should be skipped, and an error if any

	// If the date is zero, log an error and skip processing
	logHandler.ServiceLogger.Printf("Processing dateIndex entry: %v", dateIndexRecord.Date)

	if dateIndexRecord.Date.IsZero() {
		logHandler.ErrorLogger.Printf("Invalid date in dateIndex entry: %v", dateIndexRecord.Date)
		return dateIndexRecord, true, nil
	}
	if dateIndexRecord.Current.IsTrue() {
		logHandler.ServiceLogger.Printf("Skipping inactive dateIndex entry: %v", dateIndexRecord.Date)
		return dateIndexRecord, true, nil
	}
	if dateIndexRecord.Date.Equal(Today()) && dateIndexRecord.Current.IsTrue() {
		logHandler.ServiceLogger.Printf("Today's (%v) dateIndex entry already exists: %v", Today(), dateIndexRecord.Date)
		return dateIndexRecord, true, nil
	}
	if dateIndexRecord.Date.Before(Today()) {
		logHandler.ServiceLogger.Printf("Processing past dateIndex entry: %v", dateIndexRecord.Date)
		dateIndexRecord.Current.Set(false)
	}
	if dateIndexRecord.Date.After(Today()) {
		logHandler.ServiceLogger.Printf("Processing future dateIndex entry: %v", dateIndexRecord.Date)
		dateIndexRecord.Current.Set(false)
	}

	dateIndexRecord.Current.Set(true)
	return dateIndexRecord, false, nil
}
