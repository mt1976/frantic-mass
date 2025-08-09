package viewProvider

import (
	"context"

	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

// Users build the page content for the choose user page.
func Test(ctx context.Context) (contentProvider.Test, error) {
	view := contentProvider.Test{}
	var err error
	// Get the list of users from the user package
	view.Context.PageTitle = "Test Page"
	view.Context.PageSummary = "Test page for the application styling and layout."
	view.Context.PageKeywords = "test, styling, layout"
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	view.Context.TemplateName = "test"

	// Get the user ID from the context or request
	userID := "100" // This is a dummy user ID for testing purposes

	view, err = contentProvider.LoadTestView(view, userID) // Load the test view with a dummy test ID

	return view, err
}
