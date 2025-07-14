package weight

import (
	"context"

	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Data Access Object weight
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "weight" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Implement the validate function to process the domain entity
//TODO: Implement the calculate function to process the domain entity
//TODO: Implement the isDuplicateOf function to process the domain entity
//TODO: Implement the postGetProcessing function to process the domain entity

func (record *Weight) upgradeProcessing() error {
	//TODO: Add any upgrade processing here
	//This processing is triggered directly after the record has been retrieved from the database
	return nil
}

func (record *Weight) defaultProcessing() error {
	//TODO: Add any default calculations here
	//This processing is triggered directly before the record is saved to the database
	return nil
}

func (record *Weight) validationProcessing() error {
	//TODO: Add any record validation here
	//This processing is triggered directly before the record is saved to the database and after the default calculations
	return nil
}

func (record *Weight) postGetProcessing() error {
	//TODO: Add any post get processing here
	//This processing is triggered directly after the record has been retrieved from the database and after the upgrade processing
	record.BMI = record.GetBMI() // Ensure BMI is calculated after retrieval
	return nil
}

func (record *Weight) preDeleteProcessing() error {
	//TODO: Add any pre delete processing here
	//This processing is triggered directly before the record is deleted from the database
	return nil
}

func weightClone(ctx context.Context, source Weight) (Weight, error) {
	//TODO: Add any clone processing here
	panic("Not Implemented")
	return Weight{}, nil
}

func weightJobProcessor(j jobs.Job) {
	clock := timing.Start(jobs.CodedName(j), actions.PROCESS.GetCode(), j.Description())
	count := 0

	//TODO: Add your job processing code here

	// Get all the sessions
	// For each session, check the expiry date
	// If the expiry date is less than now, then delete the session

	weightEntries, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error=[%v]", jobs.CodedName(j), err.Error())
		return
	}

	noweightEntries := len(weightEntries)
	if noweightEntries == 0 {
		logHandler.ServiceLogger.Printf("[%v] No %vs to process", jobs.CodedName(j), domain)
		clock.Stop(0)
		return
	}

	for weightEntryIndex, weightRecord := range weightEntries {
		logHandler.ServiceLogger.Printf("[%v] %v(%v/%v) %v", jobs.CodedName(j), domain, weightEntryIndex+1, noweightEntries, weightRecord.Raw)
		weightRecord.UpdateWithAction(context.TODO(), audit.GRANT, "Job Processing")
		weightRecord.UpdateWithAction(context.TODO(), audit.SERVICE, "Job Processing "+j.Name())
		count++
	}
	clock.Stop(count)
}
