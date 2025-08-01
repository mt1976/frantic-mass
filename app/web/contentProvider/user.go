package contentProvider

import (
	"fmt"

	"github.com/mt1976/frantic-core/dao/lookup"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/types/measures"
	"github.com/mt1976/frantic-mass/app/web/actionHelpers"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
)

var UserWildcard = "{uid}"            // Wildcard for the user ID in the URI
var UserURI = "/user/" + UserWildcard // Define the URI for the user chooser
var UserName = "User"                 // Name for the user chooser
var UserIcon = glyphs.User            // Icon for the user chooser
var UserHover = "User %s"             // Description for the user chooser string

type UserView struct {
	ID       int
	User     user.User
	Baseline baseline.Baseline

	WeightSystemLookup   lookup.Lookup
	WeightSystem         int // Measurement system selected by the user
	WeightSystemSelected lookup.LookupData
	HeightSystemLookup   lookup.Lookup
	HeightSystem         int // Height measurement system selected by the user
	HeightSystemSelected lookup.LookupData
	Locales              lookup.Lookup

	Context AppContext
}

func GetUser(view UserView, userID string) (UserView, error) {
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "user"

	userIdInt, err := StringToInt(userID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error converting userID to int:", err)
		view.Context.AddError("Invalid user ID format")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.User = user.User{}

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
	view.WeightSystem = view.User.WeightSystem
	view.HeightSystem = view.User.HeightSystem

	view.WeightSystemLookup = measures.WeightSystemsLookup
	view.HeightSystemLookup = measures.HeightSystemsLookup
	view.HeightSystemLookup.Data[view.HeightSystem].Selected = true
	view.WeightSystemLookup.Data[view.WeightSystem].Selected = true

	view.WeightSystemSelected = measures.WeightSystemsLookup.Data[view.WeightSystem]
	view.HeightSystemSelected = measures.HeightSystemsLookup.Data[view.HeightSystem]

	view.Locales = Locales

	// Fetch the user's baseline data
	baselineRecord, err := baseline.GetBy(baseline.FIELD_UserID, userIdInt)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching baseline data:", err)
		view.Context.AddError("Error fetching baseline data")
		view.Context.AddMessage("An error occurred while fetching baseline data. Please try again later.")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.Baseline = baselineRecord

	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	// Log the successful creation of the view
	//view.Context.AddMessage("Users loaded successfully")
	view.Context.AddMessage(fmt.Sprintf("Found user %s", view.User.Username))
	//uri := DashboardURI // Use the defined URI for the dashboard
	//uri = ReplacePathParam(uri, UserWildcard, fmt.Sprintf("%d", view.User.ID))
	view.Context.PageActions.Clear()         // Clear any existing page actions
	view.Context.PageActions.AddBackAction() // Add a back action to the page actions
	view.Context.PageActions.Add(actionHelpers.NewAction("Save", "Save Changes", glyphs.Save, ReplacePathParam(UserURI, UserWildcard, IntToString(view.User.ID)), actionHelpers.UPDATE, "", style.NONE(), css.NONE()))
	logHandler.InfoLogger.Println("UserEdit view created successfully with user", view.User.Username)
	// Return the populated view

	view.Context.AddBreadcrumb(LauncherName, fmt.Sprintf(LauncherHover, view.Context.AppName), LauncherURI, LauncherIcon)
	view.Context.AddBreadcrumb(UserChooserName, UserChooserHover, UserChooserURI, UserChooserIcon)
	view.Context.AddBreadcrumb(view.User.Name, fmt.Sprintf(UserHover, view.User.Name), ReplacePathParam(DashboardURI, UserWildcard, IntToString(view.User.ID)), UserIcon)
	view.Context.AddBreadcrumb(DashboardName, fmt.Sprintf(DashboardHover, view.User.Name), ReplacePathParam(DashboardURI, UserWildcard, IntToString(view.User.ID)), DashboardIcon)
	view.Context.AddBreadcrumb(view.User.Username, "", "", glyphs.User)

	return view, nil
}
