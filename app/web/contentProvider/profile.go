package contentProvider

import (
	"time"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/dao/lookup"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/functions"
	"github.com/mt1976/frantic-mass/app/types"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

type Profile struct {
	User                      User
	Context                   AppContext
	BMI                       string
	BMINote                   string
	BMIStatus                 string
	CurrentWeight             string
	Height                    string
	DateOfBirth               string
	Age                       int
	NoGoals                   int
	Goals                     []Goal
	TotalWeightLoss           string // Total weight loss in kg
	AverageWeightLoss         string // Average weight loss in kg
	MeasurementSystemsLookup  lookup.Lookup
	MeasurementSystem         int // Measurement system selected by the user
	MeasurementSystemSelected lookup.LookupData
}

type Goal struct {
	ID              int
	Description     string
	Name            string          // Name of the goal
	TargetWeight    string          // Target weight in kilograms
	TargetBMI       string          // Target BMI
	TargetBMINote   string          // Note for the target BMI
	TargetBMIStatus string          // Status for the target BMI
	TargetDate      time.Time       // Target date for achieving the goal
	LossPerWeek     string          // Desired weight loss per week in kilograms
	IsDefault       bool            // Type of goal, e.g., user-defined or average weight loss goal
	Actions         helpers.Actions // Actions available for the user, such as edit or delete
}

func BuildProfile(view Profile, userId int) (Profile, error) {
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "profile"
	view.User = User{}
	view.Context.PageActions.Add(helpers.NewAction("Back", "Back", glyphs.Back, "/users", helpers.GET, ""))

	// Here you would typically fetch the user data based on userId
	userDetails, err := user.GetBy(user.FIELD_ID, userId)
	if err != nil {
		view.Context.WasSuccessful = false
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.AddError("User not found")
		view.Context.AddMessage("Please check the user ID and try again.")
		logHandler.ErrorLogger.Println("Error fetching user:", err)
		return view, nil
	}
	if userDetails.ID == 0 {
		view.Context.WasSuccessful = false
		view.Context.HttpStatusCode = 404 // Not Found
		view.Context.AddError("User not found")
		view.Context.AddMessage("Please check the user ID and try again.")
		logHandler.ErrorLogger.Println("Error fetching user: user record not found")
		return view, nil
	}

	view.User.ID = userDetails.ID
	view.User.Name = userDetails.Username
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
		return view, nil
	}

	view.MeasurementSystemsLookup = types.MeasurementSystemsLookup
	view.MeasurementSystem = userDetails.MeasurementSystem
	if view.MeasurementSystem < 0 || view.MeasurementSystem >= len(types.MeasurementSystems) {
		logHandler.ErrorLogger.Println("Invalid measurement system for user ID:", userId)
		view.MeasurementSystem = 0 // Default to the first measurement system
	} else {
		logHandler.InfoLogger.Println("Measurement system for user ID:", userId, "is", types.MeasurementSystems[view.MeasurementSystem].Value)
		view.MeasurementSystemSelected = view.MeasurementSystemsLookup.Data[view.MeasurementSystem]
		logHandler.InfoLogger.Println("Measurement system selected:", view.MeasurementSystemSelected.Value)
		view.MeasurementSystemsLookup.Data[view.MeasurementSystem].Selected = true // Mark the selected measurement system
	}
	// Get the latest weight record for the user
	latestWeight, err := functions.FetchLatestWeightRecord(userId)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching latest weight record:", err)
		view.Context.AddError("Error fetching latest weight record")
		view.Context.AddMessage("Please try again later.")
		return view, nil
	}

	view.BMI = latestWeight.BMI.String()                  // Placeholder for BMI calculation
	view.BMINote = latestWeight.BMI.Description           // Placeholder for BMI note
	view.BMIStatus = latestWeight.BMI.Glyph               // Placeholder for BMI status
	view.CurrentWeight = latestWeight.Weight.KgAsString() // Placeholder for current weight
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

	// Calculate total weight loss & average weight loss
	avgWeightLoss, totalWeightLoss, err := functions.AverageWeightLoss(userId)
	if err != nil {
		logHandler.ErrorLogger.Println("Error calculating weight loss for user:", userId, err)
		view.Context.AddError("Error calculating weight loss")
		view.Context.AddMessage("Please try again later.")
		return view, nil
	}
	view.TotalWeightLoss = totalWeightLoss.KgAsString() // Total weight loss in kg
	view.AverageWeightLoss = avgWeightLoss.KgAsString() // Average weight loss in

	// Fetch the user's goals
	goals, err := functions.GetGoals(userId)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching user goals:", err)
		view.Context.AddError("Error fetching user goals")
		view.Context.AddMessage("Please try again later.")
		return view, nil
	}
	view.NoGoals = len(goals)
	if view.NoGoals > 0 {
		view.Goals = make([]Goal, view.NoGoals)
		for i, g := range goals {
			view.Goals[i] = Goal{
				ID:              g.ID,
				Description:     g.Note,
				Name:            g.Name,
				TargetWeight:    g.TargetWeight.KgAsString(),
				TargetBMI:       g.TargetBMI.String(),
				TargetBMINote:   g.TargetBMI.Description,
				TargetBMIStatus: g.TargetBMI.Glyph,
				TargetDate:      g.TargetDate,
				//LossPerWeek:     g.LossPerWeek.KgAsString(),
				IsDefault: g.AverageWeightLoss.IsTrue(),
			}
			view.Goals[i].LossPerWeek = g.LossPerWeek.KgAsString()
			if g.AverageWeightLoss.IsTrue() {
				view.Goals[i].LossPerWeek = view.AverageWeightLoss
			}
			view.Goals[i].Actions = helpers.Actions{}
			view.Goals[i].Actions.Add(helpers.NewAction("Projection", "View Projection", glyphs.Projection, "/goal/projection/"+IntToString(userId)+"/"+IntToString(g.ID), helpers.GET, ""))
			view.Goals[i].Actions.Add(helpers.NewAction("Edit", "Edit Goal", glyphs.Edit, "/goal/edit/"+IntToString(g.ID), helpers.GET, ""))
			view.Goals[i].Actions.Add(helpers.NewAction("Delete", "Delete Goal", glyphs.Delete, "/goal/delete/"+IntToString(g.ID), helpers.GET, ""))
			logHandler.InfoLogger.Printf("Goal %d: %s, Target Weight: %s, Target Date: %s", g.ID, g.Name, g.TargetWeight.KgAsString(), g.TargetDate.Format("02 Jan 2006"))
		}
	}

	logHandler.InfoLogger.Println("Profile view created for user ID:", userId)

	godump.Dump(view, "Profile View")

	return view, nil
}
