package functions

import (
	"context"

	"github.com/mt1976/frantic-core/commonConfig"
	cron "github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/jobs"
)

func StartJobs(thisContext context.Context, common *commonConfig.Settings) error {
	logHandler.InfoLogger.Println("Starting Jobs")
	// Initialize the cron scheduler
	err := cron.Initialise(common)
	if err != nil {
		logHandler.ErrorLogger.Println("Error initializing cron scheduler:", err)
		return err
	}
	cron.AddJobToScheduler(jobs.UserJobInstance)
	cron.AddJobToScheduler(jobs.WeightProjectionJobInstance)
	cron.AddJobToScheduler(jobs.DateIndexJobInstance)

	logHandler.InfoLogger.Println("Jobs started successfully")
	return nil
}
