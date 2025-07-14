package dateIndex

import (
	"context"

	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Data Access Object dateIndex
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "dateIndex" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Implement the validate function to process the domain entity
//TODO: Implement the calculate function to process the domain entity
//TODO: Implement the isDuplicateOf function to process the domain entity
//TODO: Implement the postGetProcessing function to process the domain entity

func (record *DateIndex) upgradeProcessing() error {
	//TODO: Add any upgrade processing here
	//This processing is triggered directly after the record has been retrieved from the database
	return nil
}

func (record *DateIndex) defaultProcessing() error {
	//TODO: Add any default calculations here
	//This processing is triggered directly before the record is saved to the database
	return nil
}

func (record *DateIndex) validationProcessing() error {
	//TODO: Add any record validation here
	//This processing is triggered directly before the record is saved to the database and after the default calculations
	return nil
}

func (h *DateIndex) postGetProcessing() error {
	//TODO: Add any post get processing here
	//This processing is triggered directly after the record has been retrieved from the database and after the upgrade processing
	return nil
}

func (record *DateIndex) preDeleteProcessing() error {
	//TODO: Add any pre delete processing here
	//This processing is triggered directly before the record is deleted from the database
	return nil
}

func dateIndexClone(ctx context.Context, source DateIndex) (DateIndex, error) {
	//TODO: Add any clone processing here
	panic("Not Implemented")
	return DateIndex{}, nil
}

func dateIndexJobProcessor(j jobs.Job) {
	clock := timing.Start(jobs.CodedName(j), actions.PROCESS.GetCode(), j.Description())
	count := 0

	//TODO: Add your job processing code here

	// Add today's date index entry
	// This is a placeholder for the actual logic to add a date index entry
	// Replace with the actual implementation as needed

	//today := time.Now()
	logHandler.InfoLogger.Printf("[%v] Adding DateIndex for %v", jobs.CodedName(j), Today())
	_, err := Create(context.TODO(), Today())
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error adding DateIndex for %v: %v", jobs.CodedName(j), Today(), err)
	} else {
		logHandler.InfoLogger.Printf("[%v] DateIndex for %v added successfully", jobs.CodedName(j), Today())
		count++
	}

	dateList, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error retrieving dateIndex entries: %v", jobs.CodedName(j), err)
		return
	}

	if len(dateList) > 0 {
		logHandler.ServiceLogger.Printf("[%v] Found %d existing dateIndex entries", jobs.CodedName(j), len(dateList))
	} else {
		logHandler.ServiceLogger.Printf("[%v] No existing dateIndex entries found", jobs.CodedName(j))
	}

	for i := range dateList {
		logHandler.ServiceLogger.Printf("[%v] Processing dateIndex entry: %v", jobs.CodedName(j), dateList[i].Raw)
		dateIndexRecord, skip, err := classifyDateIndexRecord(&dateList[i])
		if err != nil {
			logHandler.ErrorLogger.Printf("[%v] Error processing dateIndex entry %v: %v", jobs.CodedName(j), dateList[i].Raw, err)
			continue
		}
		if skip {
			continue
		} // Set to active for processing
		_ = dateIndexRecord.UpdateWithAction(context.TODO(), audit.SERVICE, "Job Processing "+j.Name())

		count++
	}

	logHandler.ServiceLogger.Printf("[%v] Processed %d %vs", jobs.CodedName(j), count, domain)

	clock.Stop(count)
}
