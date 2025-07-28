package contentProvider

import (
	"fmt"
	"time"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/dao/lookup"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weight"
	"github.com/mt1976/frantic-mass/app/functions"
	"github.com/mt1976/frantic-mass/app/types/measures"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/graphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

var DashboardURI = "/dash/{id}" // Define the URI for the user profile

type Profile struct {
	User                 User
	Context              AppContext
	BMI                  string
	BMINote              string
	BMIStatus            string
	CurrentWeight        string
	Height               string
	DateOfBirth          string
	Age                  int
	NoGoals              int
	Goals                []Goal
	TotalWeightLoss      string // Total weight loss in kg
	AverageWeightLoss    string // Average weight loss in kg
	WeightSystemLookup   lookup.Lookup
	WeightSystem         int // Measurement system selected by the user
	WeightSystemSelected lookup.LookupData
	HeightSystemLookup   lookup.Lookup
	HeightSystem         int // Height measurement system selected by the user
	HeightSystemSelected lookup.LookupData
	Measurements         []Measurement // List of measurements for the user
}

type Measurement struct {
	ID                       int
	RecordedAt               time.Time
	Weight                   measures.Weight // Weight in kilograms
	BMI                      measures.BMI    // Body Mass Index
	Note                     string
	LossSinceLastMeasurement measures.Weight
	Actions                  helpers.Actions // Actions available for the measurement, such as edit or delete
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

func BuildUserDashboard(view Profile, userId int) (Profile, error) {
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "dashboard"
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

	view.Context.PageTitle = "User Dashboard - " + userDetails.Username
	view.Context.PageDescription = "Dashboard for user " + userDetails.Username
	view.Context.PageKeywords = "user, dashboard"
	view.Context.PageSummary = "Dashboard of the user's information for the selected user as well as their goals and measurements."
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

	view = setWeightSystem(view, userDetails, userId)
	view = setupHeightSystem(view, userDetails, userId)
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

	// Fetch the user's measurements
	userWeights, err := weight.GetAllWhere(weight.FIELD_UserID, userId)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching user weights:", err)
		view.Context.AddError("Error fetching user weights")
		view.Context.AddMessage("Please try again later.")
		return view, nil
	}
	userWeights = weight.SortByDateAscending(userWeights)
	if len(userWeights) == 0 {
		logHandler.InfoLogger.Println("No weight records found for user ID:", userId)
		view.Context.AddMessage("No weight records found for this user.")
	} else {
		logHandler.InfoLogger.Printf("Found %d weight records for user ID: %d", len(userWeights), userId)
		view.Measurements = make([]Measurement, len(userWeights))
		for i, w := range userWeights {
			view.Measurements[i] = Measurement{
				ID:         w.ID,
				RecordedAt: w.RecordedAt,
				Weight:     w.Weight,
				BMI:        w.BMI,
				Note:       w.Note,
			}
			if i > 0 {
				// Calculate the weight loss since the last measurement
				view.Measurements[i].LossSinceLastMeasurement = view.Measurements[i-1].Weight.Minus(w.Weight)
			} else {
				view.Measurements[i].LossSinceLastMeasurement = *measures.NewWeight(0) // No previous measurement, so set to zero
			}
			view.Measurements[i].Actions.Add(helpers.NewAction("View", "View Measurement", glyphs.View, "/weight/view/"+IntToString(w.ID), helpers.GET, ""))
			logHandler.InfoLogger.Printf("Measurement %d: Recorded At: %s, Weight: %s, BMI: %s", w.ID, w.RecordedAt.Format("02 Jan 2006"), w.Weight.KgAsString(), w.BMI.String())
		}
	}
	//godump.Dump(view, "Profile View")
	view = buildDashboardChart(view, userWeights, goals, "Weight Loss Progress")

	return view, nil

}

func buildDashboardChart(view Profile, weights []weight.Weight, goals []goal.Goal, chartTitle string) Profile {
	view.Context.PageHasChart = true
	view.Context.ChartID = "weightLossChart"
	view.Context.ChartTitle = chartTitle

	traces := []graphs.Trace{}

	config := graphs.NewLegendConfig(0.5, "reversed", 16, "paper")

	// Refactor to use the graphs package
	graphData := graphs.Trace{}
	graphData.Name = "Weight Loss"
	graphData.Shape = "scatter"
	graphData.XIsTime = true // X values are time-based
	noWeights := len(weights)
	for _, w := range weights {
		ds := w.RecordedAt.Format(graphs.TIMESERIES_FORMAT)
		ws := fmt.Sprintf("%v", w.Weight.Kg())
		graphData.AddXYText(ds, ws, w.Weight.KgAsString())
	}
	traces = append(traces, graphData)

	// For each goal, add a horizontal line at the target weight
	for _, goal := range goals {
		if !goal.TargetWeight.IsZero() {
			// String to float conversion
			logHandler.InfoLogger.Printf("Processing goal: %s with target weight %s", goal.Name, goal.TargetWeight)
			graphData := graphs.Trace{}
			graphData.Name = fmt.Sprintf("%s - %v", goal.Name, goal.TargetWeight.KgAsString())
			graphData.Shape = "scatter"
			graphData.XIsTime = true // X values are time-based
			// Add a horizontal line for the target weight
			for i := 0; i < noWeights; i++ {
				ds := weights[i].RecordedAt.Format(graphs.TIMESERIES_FORMAT)
				ws := fmt.Sprintf("%v", goal.TargetWeight.Kg())

				graphData.AddXYText(ds, ws, fmt.Sprintf("Goal: %s", goal.Name))
			}
			traces = append(traces, graphData)
			logHandler.InfoLogger.Printf("Added goal line for %s at target weight %s", goal.Name, goal.TargetWeight.KgAsString())
		} else {
			logHandler.WarningLogger.Printf("Goal %s has no target weight set, skipping", goal.Name)
		}
	}

	var err error
	view.Context.ChartData, err = graphs.GeneratePlotlyScript(traces, config, view.Context.ChartID)

	if err != nil {
		logHandler.ErrorLogger.Println("Error generating Plotly script:", err)
		view.Context.AddError("Error generating chart data")
		view.Context.AddMessage("Please try again later.")
		return view
	}
	godump.Dump(view.Context.ChartData)

	return view
}

func setWeightSystem(view Profile, userDetails user.User, userId int) Profile {
	view.WeightSystemLookup = measures.WeightSystemsLookup
	view.WeightSystem = userDetails.WeightSystem
	if view.WeightSystem < 0 || view.WeightSystem >= len(measures.WeightMeasurementSystems) {
		logHandler.ErrorLogger.Println("Invalid measurement system for user ID:", userId)
		view.WeightSystem = 0 // Default to the first measurement system
	} else {
		logHandler.InfoLogger.Println("Measurement system for user ID:", userId, "is", measures.WeightMeasurementSystems[view.WeightSystem].Value)
		view.WeightSystemSelected = view.WeightSystemLookup.Data[view.WeightSystem]
		logHandler.InfoLogger.Println("Measurement system selected:", view.WeightSystemSelected.Value)
		view.WeightSystemLookup.Data[view.WeightSystem].Selected = true // Mark the selected measurement system
	}
	return view
}

func setupHeightSystem(view Profile, userDetails user.User, userId int) Profile {
	view.HeightSystemLookup = measures.HeightSystemsLookup
	view.HeightSystem = userDetails.HeightSystem
	if view.HeightSystem < 0 || view.HeightSystem >= len(measures.HeightMeasurementSystems) {
		logHandler.ErrorLogger.Println("Invalid height measurement system for user ID:", userId)
		view.HeightSystem = 0 // Default to the first height measurement system
	} else {
		logHandler.InfoLogger.Println("Height measurement system for user ID:", userId, "is", measures.HeightMeasurementSystems[view.HeightSystem].Value)
		view.HeightSystemSelected = view.HeightSystemLookup.Data[view.HeightSystem]
		logHandler.InfoLogger.Println("Height measurement system selected:", view.HeightSystemSelected.Value)
		view.HeightSystemLookup.Data[view.HeightSystem].Selected = true // Mark the selected height measurement system
	}
	return view
}
