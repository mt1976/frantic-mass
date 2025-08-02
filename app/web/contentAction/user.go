package contentAction

import (
	"net/http"
	"time"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/types/measures"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func NewUser(w http.ResponseWriter, r *http.Request) (contentProvider.UserView, error) {

	godump.Dump(r)
	view := contentProvider.UserView{}
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "user"

	view.User = user.User{}

	userName := r.PostFormValue("name")
	userUserName := r.PostFormValue("username")
	pwd := r.PostFormValue("password")
	pwd = "default" // Set a default password if none is provided

	if userName == "" || userUserName == "" || pwd == "" {
		logHandler.ErrorLogger.Println("Name, Username, and Password are required for creating a new user")
		view.Context.AddError("Name, Username, and Password are required for creating a new user")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}
	if pwd == "" {
		pwd = "default" // Set a default password if none is provided
	}
	email := r.PostFormValue("email")

	newUser, err := user.Create(r.Context(), userName, userUserName, pwd, email)
	if err != nil {
		logHandler.ErrorLogger.Println("Error creating new user:", err)
		view.Context.AddError("Error creating new user")
		view.Context.HttpStatusCode = 500 // Internal Server Error
		view.Context.WasSuccessful = false
		return view, nil
	}

	hght := r.PostFormValue("height")
	if hght == "" {
		logHandler.ErrorLogger.Println("Height is required for creating a new user")
		view.Context.AddError("Height is required for creating a new user")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}

	height, err := measures.NewHeightCMSFromString(hght)

	if err != nil {
		logHandler.ErrorLogger.Println("Error creating height from string:", err)
		view.Context.AddError("Invalid height format")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}

	projectionPeriod := r.PostFormValue("projection_period")
	if projectionPeriod == "" {
		projectionPeriod = "12" // Default to 12 months if not provided
	}
	projectionPeriod_int, err := StringToInt(projectionPeriod)
	if err != nil {
		logHandler.ErrorLogger.Println("Error converting projection period to int:", err)
		view.Context.AddError("Invalid projection period format")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}
	note := r.PostFormValue("note")
	dob := r.PostFormValue("date_of_birth")
	pivotDate := r.PostFormValue("pivot_date")
	dob_int, err := time.Parse("02 Jan 2006", dob)
	if err != nil {
		logHandler.ErrorLogger.Println("Error parsing date of birth:", err)
		view.Context.AddError("Invalid date of birth format")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}
	pivotDate_int, err := time.Parse("02 Jan 2006", pivotDate)
	if err != nil {
		logHandler.ErrorLogger.Println("Error parsing pivot date:", err)
		view.Context.AddError("Invalid pivot date format")
		view.Context.HttpStatusCode = 400 // Bad Request
		view.Context.WasSuccessful = false
		return view, nil
	}

	newBaseline, err := baseline.Create(r.Context(), newUser.ID, *height, projectionPeriod_int, note, dob_int, pivotDate_int)

	//godump.Dump(newUser, newBaseline)
	view.WeightSystem = view.User.WeightSystem
	view.HeightSystem = view.User.HeightSystem

	view.WeightSystemLookup = measures.WeightSystemsLookup
	view.HeightSystemLookup = measures.HeightSystemsLookup
	view.HeightSystemLookup.Data[view.HeightSystem].Selected = true
	view.WeightSystemLookup.Data[view.WeightSystem].Selected = true

	view.WeightSystemSelected = measures.WeightSystemsLookup.Data[view.WeightSystem]
	view.HeightSystemSelected = measures.HeightSystemsLookup.Data[view.HeightSystem]

	//view.Locales = Locales

	// Fetch the user's baseline data

	view.Baseline = newBaseline

	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true
	// Log the successful creation of the view
	//view.Context.AddMessage("Users loaded successfully")
	view.Context.AddMessage("Creating a new user")
	//uri := DashboardURI // Use the defined URI for the dashboard
	//uri = ReplacePathParam(uri, UserWildcard, fmt.Sprintf("%d", view.User.ID))
	view.Context.PageActions.Clear()         // Clear any existing page actions
	view.Context.PageActions.AddBackAction() // Add a back action to the page actions
	// view.Context.PageActions.Add(actionHelpers.NewAction("Add", "Add User", glyphs.Add, contentProvider.ReplacePathParam(contentProvider.UserURI, contentProvider.UserWildcard, actionHelpers.NEW), actionHelpers.CREATE, "", style.NONE(), css.NONE()))
	// logHandler.InfoLogger.Println("User New view created successfully")
	// // Return the populated view

	// view.Context.AddBreadcrumb(contentProvider.LauncherName, fmt.Sprintf(contentProvider.LauncherHover, view.Context.AppName), contentProvider.LauncherURI, contentProvider.LauncherIcon)
	// view.Context.AddBreadcrumb(contentProvider.UserChooserName, contentProvider.UserChooserHover, contentProvider.UserChooserURI, contentProvider.UserChooserIcon)
	// view.Context.AddBreadcrumb("New User", "", "", glyphs.User)

	logHandler.InfoLogger.Println("New User view created successfully")
	http.Redirect(w, r, contentProvider.UserChooserURI, http.StatusSeeOther)

	//os.Exit(0) // Exit the program for debugging purposes, remove this in production

	return view, nil
}
