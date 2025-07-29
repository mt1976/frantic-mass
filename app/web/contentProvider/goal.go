package contentProvider

import (
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

var GoalURI = "/goal/{id}" // Define the URI for the user chooser

type GoalView struct {
	ID   int
	User user.User
	Goal goal.Goal

	Context AppContext
}

func GetGoal(view GoalView, goalID string) (GoalView, error) {
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "goal"

	userIdInt, err := StringToInt(goalID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error converting goalID to int:", err)
		view.Context.AddError("Invalid goal ID format")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.User = user.User{}

	UserGoalRecord, err := goal.GetBy(goal.FIELD_ID, userIdInt)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching goal:", err)
		view.Context.AddError("Error fetching goal")
		view.Context.AddMessage("An error occurred while fetching goal details. Please try again later.")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.Goal = UserGoalRecord

	// Fetch the user's baseline data

	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	// Log the successful creation of the view
	//view.Context.AddMessage("Users loaded successfully")
	view.Context.AddMessage(fmt.Sprintf("Found goal %s", view.Goal.Name))
	uri := DashboardURI // Use the defined URI for the dashboard
	uri = ReplacePathParam(uri, "id", fmt.Sprintf("%d", view.Goal.UserID))
	view.Context.PageActions.Add(helpers.NewAction("Back", "Back to Dashboard", glyphs.Back, uri, helpers.READ, "", style.SECONDARY(), css.NONE()))
	view.Context.PageActions.Add(helpers.NewAction("Submit", "Submit Goal Changes", glyphs.Save, "/goal/edit/"+fmt.Sprintf("%d", view.Goal.ID), helpers.CREATE, "", style.DEFAULT(), css.NONE()))
	logHandler.InfoLogger.Println("GoalEdit view created successfully with goal", view.Goal.Name)
	// Return the populated view

	return view, nil
}
