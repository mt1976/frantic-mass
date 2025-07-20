package controllers

import (
	"context"

	"github.com/mt1976/frantic-mass/app/web/views"
)

func Profile(ctx context.Context, userId int) (views.Profile, error) {
	view := views.BuildProfile(userId)
	//godump.Dump(view)

	// Set the common fields for the view
	view.Context.PageTitle = "User Profile"
	view.Context.PageKeywords = "user, profile"
	view.Context.PageSummary = "Display the user information for the selected user."
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true

	return view, nil
}
