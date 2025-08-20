package viewProvider

import (
	"context"

	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func Error(ctx context.Context, userID int, message, description, code string) (contentProvider.DisplayError, error) {

	view := contentProvider.DisplayError{}
	var err error
	// Set the common fields for the view
	view.Context.PageTitle = "Display Error"
	view.Context.PageKeywords = "display, error"
	view.Context.PageSummary = "An error occurred"
	view.Context.HttpStatusCode = 500 // Internal Server Error
	view.Context.WasSuccessful = false
	view.Context.TemplateName = "error"

	view, err = contentProvider.CreateDisplayError(view, userID, message, description, code)

	return view, err
}
