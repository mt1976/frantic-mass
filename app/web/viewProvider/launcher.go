package viewProvider

import (
	"context"

	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func Launcher(ctx context.Context) (contentProvider.DisplayLauncher, error) {

	view := contentProvider.DisplayLauncher{}
	var err error
	// Set the common fields for the view
	view.Context.PageTitle = "Display Launcher"
	view.Context.PageKeywords = "display, launcher"
	view.Context.PageSummary = "Launch the application"
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	view.Context.TemplateName = "launcher"

	view, err = contentProvider.CreateDisplayLauncher(view)

	return view, err
}
