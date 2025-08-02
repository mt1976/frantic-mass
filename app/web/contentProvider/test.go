package contentProvider

import (
	"fmt"

	"github.com/mt1976/frantic-core/dao/lookup"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/actionHelpers"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
)

var TestWildcard = "{tid}"            // Wildcard for the test ID in the URI
var TestURI = "/test/" + TestWildcard // Define the URI for the test chooser
var TestName = "Test"                 // Name for the test chooser
var TestIcon = glyphs.User            // Icon for the test chooser
var TestHover = "Test %s"             // Description for the test chooser string

type Test struct {
	ID int

	WeightSystemLookup   lookup.Lookup
	WeightSystem         int // Measurement system selected by the user
	WeightSystemSelected lookup.LookupData
	HeightSystemLookup   lookup.Lookup
	HeightSystem         int // Height measurement system selected by the user
	HeightSystemSelected lookup.LookupData
	Locales              lookup.Lookup

	Context AppContext
}

func LoadTestView(view Test, testID string) (Test, error) {
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "test"

	testIdInt, err := StringToInt(testID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error converting testID to int:", err)
		view.Context.AddError("Invalid test ID format" +
			fmt.Sprintf(" - %s", testID) + err.Error())
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}

	view.ID = testIdInt

	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	// Log the successful creation of the view
	//view.Context.AddMessage("Users loaded successfully")
	view.Context.AddMessage(fmt.Sprintf("Test Message for %s %d", cache.GetApplication_Name(), view.ID))
	view.Context.AddError(fmt.Sprintf("Test Error for %s %d", cache.GetApplication_Name(), view.ID))
	view.Context.AddNotification(fmt.Sprintf("Test Notification for %s %d", cache.GetApplication_Name(), view.ID))
	view.Context.AddSuccess(fmt.Sprintf("Test Success for %s %d", cache.GetApplication_Name(), view.ID))
	//uri := DashboardURI // Use the defined URI for the dashboard
	//uri = ReplacePathParam(uri, UserWildcard, fmt.Sprintf("%d", view.User.ID))
	view.Context.PageActions.Clear()         // Clear any existing page actions
	view.Context.PageActions.AddBackAction() // Add a back action to the page actions
	view.Context.PageActions.Add(actionHelpers.NewAction("Save", "Save Changes", glyphs.Save, ReplacePathParam(UserURI, UserWildcard, IntToString(testIdInt)), actionHelpers.UPDATE, "", style.NONE(), css.NONE()))
	// Return the populated view

	view.Context.AddBreadcrumb(LauncherName, fmt.Sprintf(LauncherHover, view.Context.AppName), LauncherURI, LauncherIcon)
	view.Context.AddBreadcrumb(TestName, fmt.Sprintf(TestHover, cache.GetApplication_Name()), ReplacePathParam(TestURI, TestWildcard, IntToString(view.ID)), TestIcon)

	return view, nil
}
