package contentProvider

import (
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weightProjection"
	"github.com/mt1976/frantic-mass/app/functions"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

var ProjectionURI = "/goal/projection/{id}/{goalId}" // Define the URI for the projection

type Projection struct {
	User        User
	Goal        Goal
	Projections []weightProjection.WeightProjection

	Context AppContext
}

func BuildProjection(view Projection, userId int, goalId int) (Projection, error) {
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "projection"
	view.User = User{}
	view.Context.PageActions.Add(helpers.NewAction("Back", "Back", glyphs.Back, UserChooserURI, helpers.READ, "", style.PRIMARY(), css.NONE()))

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
	view.Goal.Description = goalrec.Note
	view.Goal.Name = goalrec.Name
	view.Goal.TargetWeight = goalrec.TargetWeight.KgAsString()
	view.Goal.TargetBMI = goalrec.TargetBMI.String()
	view.Goal.TargetBMINote = goalrec.TargetBMI.Description
	view.Goal.TargetBMIStatus = goalrec.TargetBMI.Glyph
	view.Goal.TargetDate = goalrec.TargetDate

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

	return view, nil
}
