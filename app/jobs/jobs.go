package jobs

import "github.com/mt1976/frantic-core/jobs"

var OvernightUserJob jobs.Job = &UserJob{}
var ProjectionsPruningJob jobs.Job = &WeightProjectionJob{}
var ProjectionHistoryPruningJob jobs.Job = &weightProjectionHistoryJob{}
var OvernightRolloverJob jobs.Job = &DateIndexJob{}
