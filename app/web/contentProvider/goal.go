package contentProvider

import (
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/web/actionHelpers"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
)

var GoalWildcard = "{goalId}"         // Wildcard for the goal ID in the URI
var GoalURI = "/goal/" + GoalWildcard // Define the URI for the user chooser
var GoalName = "Goal"                 // Name for the goal chooser
var GoalIcon = glyphs.Goal            // Icon for the goal chooser
var GoalHover = "Goal %s for %s"      // Description for the goal chooser

type GoalView struct {
	ID       int
	User     user.User
	UserID   int    // User ID for the goal
	UserName string // Name of the user for display purposes
	Goal     goal.Goal

	Context AppContext
}

func GetGoal(view GoalView, goalID string) (GoalView, error) {

	thisURI := ReplacePathParam(GoalURI, GoalWildcard, goalID)
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "goal"
	view.Context.PageTitle = "Goal Details"
	view.Context.PageDescription = "View and manage your goal details"

	userGoalID, err := StringToInt(goalID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error converting goalID to int:", err)
		view.Context.AddError("Invalid goal ID format")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.User = user.User{}

	UserGoalRecord, err := goal.GetBy(goal.FIELD_ID, userGoalID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching goal:", err)
		view.Context.AddError("Error fetching goal")
		view.Context.AddMessage("An error occurred while fetching goal details. Please try again later.")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.Goal = UserGoalRecord

	userIdInt := UserGoalRecord.UserID
	UserRecord, err := user.GetBy(user.FIELD_ID, userIdInt)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching user:", err)
		view.Context.AddError("Error fetching user")
		view.Context.AddMessage("An error occurred while fetching user details. Please try again later.")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.User = UserRecord
	view.UserID = UserRecord.ID
	view.UserName = UserRecord.Name
	if view.UserName == "" {
		view.UserName = UserRecord.Username // Fallback to username if name is not set
	}

	// Fetch the user's baseline data

	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	// Log the successful creation of the view
	//view.Context.AddMessage("Users loaded successfully")
	view.Context.AddMessage(fmt.Sprintf("Found goal %s", view.Goal.Name))
	uri := DashboardURI // Use the defined URI for the dashboard
	uri = ReplacePathParam(uri, GoalWildcard, fmt.Sprintf("%d", view.Goal.UserID))
	view.Context.PageActions.Clear()          // Clear any existing page actions
	view.Context.PageActions.AddBackAction()  // Add a back action to the page actions
	view.Context.PageActions.AddPrintAction() // Add a print action to the page actions
	view.Context.PageActions.Add(actionHelpers.NewAction("Submit", "Submit Goal Changes", glyphs.Save, thisURI, actionHelpers.UPDATE, "", style.DEFAULT(), css.NONE()))
	ProjectionPath := ReplacePathParam(ProjectionURI, UserWildcard, IntToString(view.User.ID))
	ProjectionPath = ReplacePathParam(ProjectionPath, GoalWildcard, IntToString(view.Goal.ID))
	logHandler.InfoLogger.Println("Projection Path:", ProjectionPath)
	logHandler.InfoLogger.Println("Projection Path:", ProjectionPath)
	logHandler.InfoLogger.Println("Projection Path:", ProjectionPath)
	logHandler.InfoLogger.Println("Projection Path:", ProjectionPath)
	logHandler.InfoLogger.Println("Projection Path:", ProjectionPath)
	logHandler.InfoLogger.Println("Projection Path:", ProjectionPath)
	logHandler.InfoLogger.Println("Projection Path:", ProjectionPath)

	logHandler.InfoLogger.Println("GoalEdit view created successfully with goal", view.Goal.Name)

	if view.User.ID > 0 && view.Goal.ID > 0 {
		view.Context.PageActions.Add(actionHelpers.NewAction("Projection", fmt.Sprintf(ProjectionHover, view.Goal.Name, view.UserName), glyphs.Projection, ProjectionPath, actionHelpers.READ, "", style.SECONDARY(), css.NONE()))
	}
	logHandler.InfoLogger.Println("GoalEdit view created successfully with goal", view.Goal.Name)
	// Return the populated view

	// view.Context.Breadcrumbs = []Breadcrumb{
	// 	{Title: "Dashboard", URL: DashboardURI},
	// 	{Title: "Goal", URL: GoalURI},
	// 	{Title: view.Goal.Name, URL: fmt.Sprintf("/goal/%d", view.Goal.ID)},
	// }
	view.Context.AddBreadcrumb(LauncherName, fmt.Sprintf(LauncherHover, view.Context.AppName), LauncherURI, LauncherIcon)
	view.Context.AddBreadcrumb(UserChooserName, UserChooserHover, UserChooserURI, UserChooserIcon)
	view.Context.AddBreadcrumb(view.UserName, fmt.Sprintf(UserHover, view.UserName), ReplacePathParam(DashboardURI, UserWildcard, IntToString(view.User.ID)), UserIcon)
	view.Context.AddBreadcrumb(DashboardName, fmt.Sprintf(DashboardHover, view.UserName), ReplacePathParam(DashboardURI, UserWildcard, IntToString(view.User.ID)), DashboardIcon)
	view.Context.AddBreadcrumb(view.Goal.Name, fmt.Sprintf(GoalHover, view.Goal.Name, view.UserName), "", GoalIcon)

	return view, nil
}
