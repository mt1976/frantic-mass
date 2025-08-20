package contentActioner

import (
	"net/http"
	"time"

	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/htmlHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/types/measures"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
	"github.com/mt1976/frantic-mass/app/web/errorHandler"
)

func NewGoal(w http.ResponseWriter, r *http.Request, userID int) (contentProvider.GoalView, error) {
	logHandler.EventLogger.Println("NewGoal called with method:", r.Method)
	logHandler.EventLogger.Printf("Creating new goal for user ID: %d", userID)
	view := contentProvider.GoalView{}
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "goal"
	view.User = user.User{}
	//weightId := chi.URLParam(r, contentProvider.WeightWildcard)
	//userId := chi.URLParam(r, contentProvider.UserWildcard)
	// if weightId != actionHelpers.NEW {

	// 	http.Redirect(w, r, "/error?message=Invalid Weight ID", http.StatusSeeOther)
	// 	return view, fmt.Errorf("invalid weight ID %v", weightId)
	// }

	logHandler.InfoLogger.Printf("Creating new weight log entry for user ID: %d", userID)
	// inputDate := r.FormValue("dateControl")
	// inputWeight := r.FormValue("bmiWeightInput")
	// inputNote := r.FormValue("note")

	inputName := r.FormValue("name")
	inputTargetWeight := r.FormValue("targetWeight")
	inputTargetBMI := r.FormValue("bmiValue")
	inputTargetDate := r.FormValue("targetDate")
	inputIsAverageWeightLossTarget := r.FormValue("averageWeightLoss")
	inputLossPerWeek := r.FormValue("lossPerWeek")
	inputNoProjections := r.FormValue("noProjections")
	inputNote := r.FormValue("note")

	logHandler.InfoLogger.Printf("Creating new weight log entry for user ID: %d, Name: %s, Target Weight: %s", userID, inputName, inputTargetWeight)

	if inputTargetDate == "" || inputTargetWeight == "" {
		http.Redirect(w, r, "/error?message=Invalid Input", http.StatusSeeOther)
		return view, nil
	}
	// Check this is a valid user
	_, err := user.GetBy(user.FIELD_ID, userID)
	if err != nil {
		logHandler.ErrorLogger.Printf("Error fetching user %d: %v", userID, err)
		errorHandler.Error(w, r, contentProvider.IntToString(userID), "Invalid User", "User not found or invalid", "404")
		return view, err
	}
	goalNoProjections := htmlHelpers.ValueToInt(inputNoProjections)

	// convert date to internal
	targetDate_internal, err := time.Parse(dateHelpers.Format.YMD, inputTargetDate)
	logHandler.InfoLogger.Printf("Parsed target date: %v, Input: %v", targetDate_internal, inputTargetDate)
	if err != nil {
		errorHandler.Error(w, r, contentProvider.IntToString(userID), "Invalid Date", "Failed to parse date", "400")
		return view, err
	}

	if targetDate_internal.IsZero() {
		errorHandler.Error(w, r, contentProvider.IntToString(userID), "Invalid Date", "Target date is zero", "400")
		return view, err
	}

	targetWeight_internal, err := StringToFloat(inputTargetWeight)
	if err != nil {
		errorHandler.Error(w, r, contentProvider.IntToString(userID), "Invalid Weight", "Failed to parse target weight", "400")
		return view, err
	}
	targetWeight_measure := measures.NewWeight(targetWeight_internal)

	targetBMI_internal, err := StringToFloat(inputTargetBMI)
	if err != nil {
		errorHandler.Error(w, r, contentProvider.IntToString(userID), "Invalid BMI", "Failed to parse target BMI", "400")
		return view, err
	}
	targetBMI_measure := measures.NewBMI(targetBMI_internal)
	logHandler.InfoLogger.Printf("Target Weight: %f, Target BMI: %f, AverageWeightLoss: %b, Loss Per Week: %f", targetWeight_internal, targetBMI_measure.BMI, inputIsAverageWeightLossTarget, inputLossPerWeek)

	targetLossPerWeek_measure, err := measures.NewWeightFromString(inputLossPerWeek)
	if err != nil {
		errorHandler.Error(w, r, contentProvider.IntToString(userID), "Invalid Loss Per Week", "Failed to parse loss per week", "400")
		logHandler.ErrorLogger.Printf("Error creating loss per week measure: %v", err)
		return view, err
	}

	isAverageType := htmlHelpers.ValueToBool(inputIsAverageWeightLossTarget)

	if goalNoProjections < 1 {
		errorHandler.Error(w, r, contentProvider.IntToString(userID), "Invalid No Projections", "No projections must be at least 1", "400")
		return view, err
	}

	// Create the goal
	_, err = goal.Create(r.Context(), userID, inputName, *targetWeight_measure, *targetBMI_measure, targetDate_internal, *targetLossPerWeek_measure, inputNote, isAverageType, goalNoProjections)

	if err != nil {
		logHandler.ErrorLogger.Printf("Error creating weight entry for user %d: %v", userID, err)
		errorHandler.Error(w, r, contentProvider.IntToString(userID), "Failed to Create Goal", "Error creating goal entry", "500")
		return view, err
	}
	logHandler.InfoLogger.Printf("No Projections: %v", inputNoProjections)
	userIDStr := contentProvider.IntToString(userID)
	nextURI := contentProvider.ReplacePathParam(contentProvider.DashboardURI, contentProvider.UserWildcard, userIDStr)
	logHandler.InfoLogger.Printf("Redirecting to [%s] after creating new goal for user ID: %d", nextURI, userID)
	logHandler.EventLogger.Println("New Goal created successfully")
	logHandler.EventLogger.Println("New Goal created successfully")
	logHandler.EventLogger.Println("New Goal created successfully")
	logHandler.EventLogger.Println("New Goal created successfully")

	http.Redirect(w, r, nextURI, http.StatusSeeOther)

	return view, nil
}

func UpdateGoal(w http.ResponseWriter, r *http.Request, userID int, goalID int) (contentProvider.GoalView, error) {
	logHandler.EventLogger.Printf("Updating goal entry %d for user ID: %d", goalID, userID)
	view := contentProvider.GoalView{}
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "goal"
	view.User = user.User{}

	return view, nil
}
