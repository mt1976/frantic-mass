package viewProvider

import (
	"context"

	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

// Users build the page content for the choose user page.
func Users(ctx context.Context) (contentProvider.UserChooser, error) {
	view := contentProvider.UserChooser{}
	var err error
	// Get the list of users from the user package
	view.Context.PageTitle = "Choose User"
	view.Context.PageSummary = "Select a user to proceed with the application."
	view.Context.PageKeywords = "user, choose, select"
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "users"

	view, err = contentProvider.CreateUserChooser(view)

	return view, err
}
