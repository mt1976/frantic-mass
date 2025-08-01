package viewProvider

import (
	"context"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func Projection(ctx context.Context, userId int, goalId int) (contentProvider.Projection, error) {

	logHandler.EventLogger.Println("Creating Projection view for user ID:", userId, "and goal ID:", goalId)
	//godump.Dump(view)
	view := contentProvider.Projection{}
	var err error

	// Set the common fields for the view
	view.Context.PageTitle = "Goal Projection"
	view.Context.PageKeywords = "user, goal, projection"
	view.Context.PageSummary = "Display the goal projection for the selected user."
	view.Context.TemplateName = "projection"
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true

	view, err = contentProvider.BuildProjection(view, userId, goalId)

	return view, err
}
