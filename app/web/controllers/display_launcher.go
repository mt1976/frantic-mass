package controllers

import (
	"context"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-mass/app/web/views"
)

func Launcher(ctx context.Context) (views.DisplayLauncher, error) {
	view := views.CreateDisplayLauncher()
	godump.Dump(view)

	// Set the common fields for the view
	view.Common.PageTitle = "Display Launcher"
	view.Common.PageKeywords = "display, launcher"
	view.Common.PageSummary = "Launch the application"
	view.Common.HttpStatusCode = 200 // OK
	view.Common.WasSuccessful = true

	return view, nil
}
