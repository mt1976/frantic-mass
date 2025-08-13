package contentActioner

import (
	"net/http"
	"time"

	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weight"
	"github.com/mt1976/frantic-mass/app/types/measures"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func NewWeightLogEntry(w http.ResponseWriter, r *http.Request, userID int) (contentProvider.WeightView, error) {
	logHandler.EventLogger.Println("NewWeightLogEntry called with method:", r.Method)
	logHandler.EventLogger.Printf("Creating new weight log entry for user ID: %d", userID)
	view := contentProvider.WeightView{}
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "weight"
	view.User = user.User{}
	//weightId := chi.URLParam(r, contentProvider.WeightWildcard)
	//userId := chi.URLParam(r, contentProvider.UserWildcard)
	// if weightId != actionHelpers.NEW {

	// 	http.Redirect(w, r, "/error?message=Invalid Weight ID", http.StatusSeeOther)
	// 	return contentProvider.WeightView{}, fmt.Errorf("invalid weight ID %v", weightId)
	// }

	logHandler.InfoLogger.Printf("Creating new weight log entry for user ID: %d", userID)
	inputDate := r.FormValue("dateControl")
	inputWeight := r.FormValue("bmiWeightInput")
	inputNote := r.FormValue("note")

	logHandler.InfoLogger.Printf("Creating new weight log entry for user ID: %d, Date: %s, Weight: %s, Note: %s", userID, inputDate, inputWeight, inputNote)
	logHandler.InfoLogger.Printf("Creating new weight log entry for user ID: %d, Date: %s, Weight: %s, Note: %s", userID, inputDate, inputWeight, inputNote)

	if inputDate == "" || inputWeight == "" {
		http.Redirect(w, r, "/error?message=Invalid Input", http.StatusSeeOther)
		return contentProvider.WeightView{}, nil
	}
	// Check this is a valid user
	_, err := user.GetBy(user.FIELD_ID, userID)
	if err != nil {
		logHandler.ErrorLogger.Printf("Error fetching user %d: %v", userID, err)
		http.Redirect(w, r, "/error?message=Invalid User", http.StatusSeeOther)
		return contentProvider.WeightView{}, err
	}

	// convert date to internal
	dateInternal, err := time.Parse(dateHelpers.Format.YMD, inputDate)
	if err != nil {
		http.Redirect(w, r, "/error?message=Invalid Date", http.StatusSeeOther)
		return contentProvider.WeightView{}, err
	}

	if dateInternal.IsZero() {
		http.Redirect(w, r, "/error?message=Invalid Date", http.StatusSeeOther)
		return contentProvider.WeightView{}, err
	}
	if dateInternal.After(time.Now().AddDate(0, 0, 1)) {
		http.Redirect(w, r, "/error?message=Invalid Date", http.StatusSeeOther)
		return contentProvider.WeightView{}, err
	}
	weightFloat, err := StringToFloat(inputWeight)
	if err != nil {
		http.Redirect(w, r, "/error?message=Invalid Weight", http.StatusSeeOther)
		return contentProvider.WeightView{}, err
	}
	weightInternal := measures.NewWeight(weightFloat)

	_, err = weight.Create(r.Context(), userID, *weightInternal, inputNote, dateInternal)
	if err != nil {
		logHandler.ErrorLogger.Printf("Error creating weight entry for user %d: %v", userID, err)
		http.Redirect(w, r, "/error?message=Failed to Create Weight Entry", http.StatusSeeOther)
		return contentProvider.WeightView{}, err
	}
	userIDStr := contentProvider.IntToString(userID)
	nextURI := contentProvider.ReplacePathParam(contentProvider.DashboardURI, contentProvider.UserWildcard, userIDStr)
	logHandler.InfoLogger.Printf("Redirecting to [%s] after creating new weight log entry for user ID: %d", nextURI, userID)
	logHandler.EventLogger.Println("New Weight Log created successfully")
	http.Redirect(w, r, nextURI, http.StatusSeeOther)

	return view, nil
}

func UpdateWeightLogEntry(w http.ResponseWriter, r *http.Request, userID int, weightID int) (contentProvider.WeightView, error) {
	logHandler.EventLogger.Printf("Updating weight log entry %d for user ID: %d", weightID, userID)
	return contentProvider.WeightView{}, nil
}
