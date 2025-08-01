package contentProvider

import (
	"fmt"
	"time"

	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weight"
	"github.com/mt1976/frantic-mass/app/types/measures"
	"github.com/mt1976/frantic-mass/app/web/actionHelpers"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
)

var WeightWildcard = "{weightId}"            // Wildcard for the weight ID in the URI
var WeightURI = "/weight/" + WeightWildcard  // Define the URI for the weight measurement
var WeightName = "Weight"                    // Name for the weight measurement
var WeightIcon = glyphs.Weight               // Icon for the weight measurement
var WeightHover = "User %s Weight Log on %s" // Description for the weight measurement

type WeightView struct {
	ID          int
	User        user.User
	UserID      int             // User ID for the weight
	UserName    string          // Name of the user for display purposes
	RecordedAt  time.Time       `storm:"index"`
	Weight      measures.Weight // Weight in kilograms
	BMI         measures.BMI    // Body Mass Index
	Note        string
	Date        string // Formatted date for display
	DateControl string // Date control for user input
	// Context holds the common view context
	Context AppContext
}

func GetWeight(view WeightView, weightIdentifier int) (WeightView, error) {
	logHandler.EventLogger.Println("Creating Weight view for weight ID:", weightIdentifier)
	weightID := fmt.Sprintf("%d", weightIdentifier) // Ensure weightID is a string for URI replacement
	thisURI := ReplacePathParam(WeightURI, WeightWildcard, weightID)
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "weight"
	view.Context.PageTitle = "Weight Details"
	view.Context.PageDescription = "View and manage your weight details"
	view.Context.PageActions = actionHelpers.Actions{} // Initialize the PageActions
	view.User = user.User{}

	UserWeightRecord, err := weight.GetBy(weight.FIELD_ID, weightIdentifier)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching weight:", err)
		view.Context.AddError("Error fetching weight")
		view.Context.AddMessage("An error occurred while fetching weight details. Please try again later.")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.Weight = UserWeightRecord.Weight
	view.ID = UserWeightRecord.ID
	view.RecordedAt = UserWeightRecord.RecordedAt
	view.BMI = UserWeightRecord.BMI
	view.Note = UserWeightRecord.Note

	date := UserWeightRecord.RecordedAt
	if date.IsZero() {
		date = UserWeightRecord.Audit.CreatedAt // Fallback to CreatedAt if RecordedAt is zero
	}
	view.Date = dateHelpers.FormatHuman(date)      // Format the date for display
	view.DateControl = dateHelpers.FormatYMD(date) // Format the date for control input

	userIdInt := UserWeightRecord.UserID
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
	view.UserName = UserRecord.GetUserName()
	// Fetch the user's baseline data

	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	// Log the successful creation of the view
	//view.Context.AddMessage("Users loaded successfully")
	view.Context.AddMessage(fmt.Sprintf("Found weight %d", view.ID))
	view.Context.AddMessage(fmt.Sprintf("Found user %s", view.UserName))
	//uri := DashboardURI // Use the defined URI for the dashboard
	//uri = ReplacePathParam(uri, WeightWildcard, fmt.Sprintf("%d", view.Weight.UserID))
	view.Context.PageActions.Clear()         // Clear any existing page actions
	view.Context.PageActions.AddBackAction() // Add a back action to the page actions
	view.Context.PageActions.Add(actionHelpers.NewAction("Submit", "Submit Weight Changes", glyphs.Save, thisURI, actionHelpers.UPDATE, "", style.DEFAULT(), css.NONE()))

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

	view.Context.AddBreadcrumb(view.Date, fmt.Sprintf(WeightHover, view.UserName, view.Date), "", WeightIcon)
	return view, nil
}
