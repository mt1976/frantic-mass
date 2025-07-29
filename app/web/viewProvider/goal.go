package viewProvider

import (
	"context"

	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func Goal(ctx context.Context, goalID string) (contentProvider.GoalView, error) {
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
