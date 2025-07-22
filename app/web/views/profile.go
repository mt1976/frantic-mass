package views

import (
	"time"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

type Profile struct {
	User          User
	Context       AppContext
	BMI           string
	CurrentWeight string
	Height        string
	DateOfBirth   string
	Age           int
}

func BuildProfile(userId int) Profile {
	view := Profile{}
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "profile"
	view.User = User{}
	view.Context.PageActions.Add(helpers.NewAction("Back", "Back", glyphs.Back, "/users", helpers.GET, ""))

	// Here you would typically fetch the user data based on userId
	UserRecord, err := user.GetBy(user.FIELD_ID, userId)
	if err != nil {
		view.Context.WasSuccessful = false
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.AddError("User not found")
		view.Context.AddMessage("Please check the user ID and try again.")
		logHandler.ErrorLogger.Println("Error fetching user:", err)
		return view
	}
	if UserRecord.ID == 0 {
		view.Context.WasSuccessful = false
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.AddError("User not found")
		view.Context.AddMessage("Please check the user ID and try again.")
		logHandler.ErrorLogger.Println("Error fetching user: user record not found")
		return view
	}

	view.User.ID = UserRecord.ID
	view.User.Name = UserRecord.Username
	// Add more fields as necessary

	view.Context.PageTitle = "User Profile"
	view.Context.PageKeywords = "user, profile"
	view.Context.PageSummary = "Display the user information for the selected user."
	view.Context.HttpStatusCode = 200 // OK
	view.Context.WasSuccessful = true

	// Get the current baseline for the user
	baseline, err := baseline.GetByUserID(userId)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching user baseline:", err)
		view.Context.AddError("Error fetching user baseline")
		view.Context.AddMessage("Please try again later.")
		return view
	}

	view.BMI = ""                // Placeholder for BMI calculation
	view.CurrentWeight = "100kg" // Placeholder for current weight
	if baseline.Height.IsZero() {
		logHandler.ErrorLogger.Println("Height not set for user ID:", userId)
		view.Height = "Not set"
	} else {
		logHandler.InfoLogger.Println("Height for user ID:", userId, "is", baseline.Height.CmAsString())
		view.Height = baseline.Height.CmAsString() // Convert height to string representation
	}
	if baseline.DateOfBirth.IsZero() {
		logHandler.ErrorLogger.Println("Date of Birth not set for user ID:", userId)
		view.DateOfBirth = "Not set"
	} else {
		logHandler.InfoLogger.Println("Date of Birth for user ID:", userId, "is", baseline.DateOfBirth)
		view.DateOfBirth = baseline.DateOfBirth.Format("02/01/2006") // Format the date as needed
	}
	view.DateOfBirth = baseline.DateOfBirth.Format("02/01/2006") // Format the date as needed
	view.Age = AgeFromDOB(baseline.DateOfBirth)

	logHandler.InfoLogger.Println("Profile view created for user ID:", userId)

	return view
}

// AgeFromDOB calculates the age in years given a date of birth
func AgeFromDOB(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// Check if the birthday has occurred yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}
