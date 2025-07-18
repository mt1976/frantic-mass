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
	view.Common.Title = "Display Launcher"
	view.Common.Keywords = "display, launcher"
	view.Common.Status = 200 // OK
	view.Common.Success = true

	return view, nil
}
