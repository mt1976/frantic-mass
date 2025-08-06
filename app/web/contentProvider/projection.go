package contentProvider

import (
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weightProjection"
	"github.com/mt1976/frantic-mass/app/functions"
	"github.com/mt1976/frantic-mass/app/web/actionHelpers"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
)

var ProjectionWildcard = ""                                            // Wildcard for the goal ID in the URI
var ProjectionURI = "/projection/" + UserWildcard + "/" + GoalWildcard // Define the URI for the projection
var ProjectionName = "Projection"                                      // Name for the projection
var ProjectionIcon = glyphs.Projection                                 // Icon for the projection
var ProjectionHover = "Projection of goal %s for %s"                   // Description for the projection

type Projection struct {
	User        User
	Goal        Goal
	Projections []weightProjection.WeightProjection

	Context AppContext
}

func BuildProjection(view Projection, userId int, goalId int) (Projection, error) {

	logHandler.EventLogger.Println("Building Projection view for user ID:", userId, "and goal ID:", goalId)

	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "projection"
	view.User = User{}
	view.Context.PageActions = actionHelpers.Actions{} // Initialize the PageActions
	view.Context.PageTitle = "Projection Details"
	view.Context.PageDescription = "View and manage your projection details"
	view.Context.PageActions.AddBackAction()
	view.Context.PageActions.AddPrintAction()
	view.Context.PageActions.AddSubmitButton(DashboardName, DashboardHover, glyphs.Dashboard, ReplacePathParam(DashboardURI, UserWildcard, IntToString(userId)), actionHelpers.READ, "", style.NONE(), css.NONE())

	// Here you would typically fetch the user data based on userId
	UserRecord, err := user.GetBy(user.FIELD_ID, userId)
	if err != nil {
		view.Context.WasSuccessful = false
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.AddError("User not found")
		view.Context.AddMessage("Please check the user ID and try again.")
		logHandler.ErrorLogger.Println("Error fetching user:", err)
		return view, nil
	}
	if UserRecord.ID == 0 {
		view.Context.WasSuccessful = false
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.AddError("User not found")
		view.Context.AddMessage("Please check the user ID and try again.")
		logHandler.ErrorLogger.Println("Error fetching user: user record not found")
		return view, nil
	}

	view.User.ID = UserRecord.ID
	view.User.Name = UserRecord.Username
	// Add more fields as necessary

	view.Context.PageTitle = "User Profile"
	view.Context.PageKeywords = "user, profile"
	view.Context.PageSummary = "Display the user information for the selected user."
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true

	// Fetch the user's goals
	goalrec, err := goal.GetBy(goal.FIELD_ID, goalId)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching user goal:", err)
		view.Context.AddError("Error fetching user goal")
		view.Context.AddMessage("Please try again later.")
		return view, nil
	}

	view.Goal.ID = goalrec.ID
	view.Goal.Note = goalrec.Note
	view.Goal.Name = goalrec.Name
	view.Goal.TargetWeight = goalrec.TargetWeight.KgAsString()
	view.Goal.TargetBMI = goalrec.TargetBMI.String()
	view.Goal.TargetBMINote = goalrec.TargetBMI.Description
	view.Goal.TargetBMIStatus = goalrec.TargetBMI.Glyph
	view.Goal.TargetDate = goalrec.TargetDate
	view.Goal.Description = goalrec.Description
	if view.Goal.Description == "" {
		view.Goal.Description = view.Goal.Note // Fallback to Note if Description is empty
	}
	if view.Goal.Description == "" {
		view.Goal.Description = view.Goal.Name // Fallback to Name if both Note and Description are empty
	}
	if view.Goal.Description == "" {
		view.Goal.Description = "No description provided" // Default message if all are empty
	}

	// Get Avg Weight Loss
	if goalrec.AverageWeightLoss.IsTrue() {
		avgWeightLoss, _, err := functions.AverageWeightLoss(userId)
		if err != nil {
			logHandler.ErrorLogger.Println("Error calculating average weight loss for user:", userId, err)
			view.Context.AddError("Error calculating average weight loss")
			view.Context.AddMessage("Please try again later.")
			return view, nil
		}
		view.Goal.LossPerWeek = avgWeightLoss.KgAsString()
	} else {
		view.Goal.LossPerWeek = goalrec.LossPerWeek.KgAsString()
	}
	view.Goal.IsDefault = goalrec.AverageWeightLoss.IsTrue()

	// Fetch the weight projections for the user and goal
	projections, err := weightProjection.GetAllWhere(weightProjection.FIELD_UserID, userId)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching weight projections:", err)
		view.Context.AddError("Error fetching weight projections")
		view.Context.AddMessage("Please try again later.")
		return view, nil
	}

	view.Projections = projections

	if len(view.Projections) == 0 {
		logHandler.InfoLogger.Println("No weight projections found for user ID:", userId)
		view.Context.AddMessage("No weight projections found for this user.")
	} else {
		logHandler.InfoLogger.Println("Found", len(view.Projections), "weight projections for user ID:", userId)
	}

	// Sort and filter projections by date ascending, and where goal ID matches
	view.Projections = weightProjection.FilterByGoalID(view.Projections, goalId)
	view.Projections = weightProjection.FilterByUserID(view.Projections, userId)

	view.Projections = weightProjection.SortByDateAscending(view.Projections)

	if len(view.Projections) == 0 {
		logHandler.InfoLogger.Println("No weight projections found for user ID:", userId)
		view.Context.AddMessage("No weight projections found for this user.")
	} else {
		logHandler.InfoLogger.Println("Found", len(view.Projections), "weight projections for user ID:", userId)
	}

	logHandler.InfoLogger.Println("Projection view created for user ID:", userId, "and goal ID:", goalId)

	view.Context.PageActions.AddSubmitButton(GoalName, fmt.Sprintf(GoalHover, view.Goal.Name, view.User.Name), glyphs.Goal, ReplacePathParam(GoalURI, GoalWildcard, IntToString(view.Goal.ID)), actionHelpers.READ, "", style.NONE(), css.NONE())

	view.Context.AddBreadcrumb(LauncherName, fmt.Sprintf(LauncherHover, view.Context.AppName), LauncherURI, LauncherIcon)
	view.Context.AddBreadcrumb(UserChooserName, UserChooserHover, UserChooserURI, UserChooserIcon)
	view.Context.AddBreadcrumb(view.User.Name, fmt.Sprintf(UserHover, view.User.Name), ReplacePathParam(DashboardURI, UserWildcard, IntToString(view.User.ID)), UserIcon)
	view.Context.AddBreadcrumb(DashboardName, fmt.Sprintf(DashboardHover, view.User.Name), ReplacePathParam(DashboardURI, UserWildcard, IntToString(view.User.ID)), DashboardIcon)
	view.Context.AddBreadcrumb(view.Goal.Name, fmt.Sprintf(ProjectionHover, view.Goal.Name, view.User.Name), "", glyphs.Projection)

	//godump.Dump(view, "Projection View")
	//os.Exit(0) // Remove this line in production code

	return view, nil
}
