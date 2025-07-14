package weightTag

import (
	"context"

	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Data Access Object weightTag
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "weightTag" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Implement the validate function to process the domain entity
//TODO: Implement the calculate function to process the domain entity
//TODO: Implement the isDuplicateOf function to process the domain entity
//TODO: Implement the postGetProcessing function to process the domain entity

func (record *WeightTag) upgradeProcessing() error {
	//TODO: Add any upgrade processing here
	//This processing is triggered directly after the record has been retrieved from the database
	return nil
}

func (record *WeightTag) defaultProcessing() error {
	//TODO: Add any default calculations here
	//This processing is triggered directly before the record is saved to the database
	return nil
}

func (record *WeightTag) validationProcessing() error {
	//TODO: Add any record validation here
	//This processing is triggered directly before the record is saved to the database and after the default calculations
	return nil
}

func (h *WeightTag) postGetProcessing() error {
	//TODO: Add any post get processing here
	//This processing is triggered directly after the record has been retrieved from the database and after the upgrade processing
	return nil
}

func (record *WeightTag) preDeleteProcessing() error {
	//TODO: Add any pre delete processing here
	//This processing is triggered directly before the record is deleted from the database
	return nil
}

func weightTagClone(ctx context.Context, source WeightTag) (WeightTag, error) {
	//TODO: Add any clone processing here
	panic("Not Implemented")
	return WeightTag{}, nil
}

func weightTagJobProcessor(j jobs.Job) {
	clock := timing.Start(jobs.CodedName(j), actions.PROCESS.GetCode(), j.Description())
	count := 0

	//TODO: Add your job processing code here

	// Get all the sessions
	// For each session, check the expiry date
	// If the expiry date is less than now, then delete the session

	weightTagEntries, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error=[%v]", jobs.CodedName(j), err.Error())
		return
	}

	noweightTagEntries := len(weightTagEntries)
	if noweightTagEntries == 0 {
		logHandler.ServiceLogger.Printf("[%v] No %vs to process", jobs.CodedName(j), domain)
		clock.Stop(0)
		return
	}

	for weightTagEntryIndex, weightTagRecord := range weightTagEntries {
		logHandler.ServiceLogger.Printf("[%v] %v(%v/%v) %v", jobs.CodedName(j), domain, weightTagEntryIndex+1, noweightTagEntries, weightTagRecord.Raw)
		weightTagRecord.UpdateWithAction(context.TODO(), audit.GRANT, "Job Processing")
		weightTagRecord.UpdateWithAction(context.TODO(), audit.SERVICE, "Job Processing "+j.Name())
		count++
	}
	clock.Stop(count)
}
