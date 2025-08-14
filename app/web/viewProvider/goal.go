package viewProvider

import (
	"context"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func ViewGoal(ctx context.Context, goalID string) (contentProvider.GoalView, error) {
	logHandler.EventLogger.Println("ViewGoal called with goalID:", goalID)
	view := contentProvider.GoalView{}
	var err error
	// Set the common fields for the view
	view.Context.PageTitle = "View/Edit Goal " + goalID
	view.Context.PageSummary = "View or edit the goal information."
	view.Context.PageKeywords = "goal, edit"
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	view.Context.TemplateName = "goal"

	view, err = contentProvider.GetGoal(view, goalID)

	return view, err
}

func NewGoal(ctx context.Context, userID int) (contentProvider.GoalView, error) {
	logHandler.EventLogger.Println("NewGoal called with userID:", userID)
	//godump.Dump(view)
	view := contentProvider.GoalView{}
	var err error

	// Set the common fields for the view
	view.Context.PageTitle = "Progress Log Entry"
	view.Context.PageKeywords = "user, goal, weight, log"
	view.Context.PageSummary = "Enter the weight log for the selected user."
	view.Context.TemplateName = "weight"
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true

	view, err = contentProvider.NewGoal(view, userID)

	return view, err
}
