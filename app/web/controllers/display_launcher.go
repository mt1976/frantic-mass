package controllers

import (
	"context"

	"github.com/mt1976/frantic-mass/app/web/views"
)

func Launcher(ctx context.Context) (views.DisplayLauncher, error) {
	view := views.CreateDisplayLauncher()
	//godump.Dump(view)

	// Set the common fields for the view
	view.Context.PageTitle = "Display Launcher"
	view.Context.PageKeywords = "display, launcher"
	view.Context.PageSummary = "Launch the application"
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true

	return view, nil
}
