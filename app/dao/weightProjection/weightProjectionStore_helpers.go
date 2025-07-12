package weightProjection

import (
	"context"

	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Data Access Object weightProjection
// Version: 0.2.0
// Updated on: 2021-09-10

//TODO: RENAME "weightProjection" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Implement the validate function to process the domain entity
//TODO: Implement the calculate function to process the domain entity
//TODO: Implement the isDuplicateOf function to process the domain entity
//TODO: Implement the postGetProcessing function to process the domain entity

func (record *WeightProjection) upgradeProcessing() error {
	//TODO: Add any upgrade processing here
	//This processing is triggered directly after the record has been retrieved from the database
	return nil
}

func (record *WeightProjection) defaultProcessing() error {
	//TODO: Add any default calculations here
	//This processing is triggered directly before the record is saved to the database
	record.BMI = record.GetBMI() // Ensure BMI is calculated before saving
	return nil
}

func (record *WeightProjection) validationProcessing() error {
	//TODO: Add any record validation here
	//This processing is triggered directly before the record is saved to the database and after the default calculations
	return nil
}

func (h *WeightProjection) postGetProcessing() error {
	//TODO: Add any post get processing here
	//This processing is triggered directly after the record has been retrieved from the database and after the upgrade processing
	h.BMI = h.GetBMI() // Ensure BMI is calculated after retrieval
	return nil
}

func (record *WeightProjection) preDeleteProcessing() error {
	//TODO: Add any pre delete processing here
	//This processing is triggered directly before the record is deleted from the database
	return nil
}

func weightProjectionClone(ctx context.Context, source WeightProjection) (WeightProjection, error) {
	//TODO: Add any clone processing here
	panic("Not Implemented")
	return WeightProjection{}, nil
}

func weightProjectionJobProcessor(j jobs.Job) {
	clock := timing.Start(jobs.CodedName(j), actions.PROCESS.GetCode(), j.Description())
	count := 0

	//TODO: Add your job processing code here

	// Get all the sessions
	// For each session, check the expiry date
	// If the expiry date is less than now, then delete the session

	weightProjectionEntries, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error=[%v]", jobs.CodedName(j), err.Error())
		return
	}

	noweightProjectionEntries := len(weightProjectionEntries)
	if noweightProjectionEntries == 0 {
		logHandler.ServiceLogger.Printf("[%v] No %vs to process", jobs.CodedName(j), domain)
		clock.Stop(0)
		return
	}

	for weightProjectionEntryIndex, weightProjectionRecord := range weightProjectionEntries {
		logHandler.ServiceLogger.Printf("[%v] %v(%v/%v) %v", jobs.CodedName(j), domain, weightProjectionEntryIndex+1, noweightProjectionEntries, weightProjectionRecord.Raw)
		//		weightProjectionRecord.UpdateWithAction(context.TODO(), audit.SERVICE, "Job Processing "+j.Name())

		// Here we add the job processing logic

		count++
	}
	clock.Stop(count)
}
