package viewProvider

import (
	"context"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func Weight(ctx context.Context, weightId string) (contentProvider.WeightView, error) {

	logHandler.EventLogger.Println("Creating Weight view for weight ID:", weightId)
	//godump.Dump(view)
	view := contentProvider.WeightView{}
	var err error

	// Set the common fields for the view
	view.Context.PageTitle = "Progress Log Entry"
	view.Context.PageKeywords = "user, goal, weight, log"
	view.Context.PageSummary = "Display the weight log for the selected user."
	view.Context.TemplateName = "weight"
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true

	intWeightID, err := contentProvider.StringToInt(weightId)
	if err != nil {
		logHandler.ErrorLogger.Println("Error converting weight ID to int:", err)
		view.Context.AddError("Invalid weight ID format")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}

	view, err = contentProvider.GetWeight(view, intWeightID)

	return view, err
}
