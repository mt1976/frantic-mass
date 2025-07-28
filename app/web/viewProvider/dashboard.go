package viewProvider

import (
	"context"

	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func Dashboard(ctx context.Context, userId int) (contentProvider.Dashboard, error) {

	//godump.Dump(view)
	view := contentProvider.Dashboard{}
	var err error
	// Set the common fields for the view
	view.Context.PageTitle = "User Profile"
	view.Context.PageKeywords = "user, profile"
	view.Context.PageSummary = "Display the user information for the selected user."
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	view.Context.TemplateName = "profile"

	view, err = contentProvider.BuildUserDashboard(view, userId)

	return view, err
}
