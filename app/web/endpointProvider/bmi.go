package endpointProvider

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/types/measures"
)

var BMIURI = "/bmi"                                                           // Define the URI for the BMI endpoint
var BMIUserWildcard = "{userID}"                                              // Define the user ID wildcard
var BMIWeightWildcard = "{weight}"                                            // Define the weight wildcard
var BMIWildcardURI = BMIURI + "/" + BMIUserWildcard + "/" + BMIWeightWildcard // Define the wildcard URI for the BMI endpoint

func BMI(w http.ResponseWriter, r *http.Request) {
	// This is a dummy router function for BMI calculation
	logHandler.EventLogger.Printf("Entered BMI")
	// lets use chi to get the values
	userID := chi.URLParam(r, strip(BMIUserWildcard))
	userWeight := chi.URLParam(r, strip(BMIWeightWildcard))

	logHandler.InfoLogger.Printf("UserID [%v] Weight [%v]", userID, userWeight)

	// Convert userWeight to float64 if necessary
	// For example, if it's a string, you might need to parse it:
	userWeightFloat, err := strconv.ParseFloat(userWeight, 64)
	if err != nil {
		http.Error(w, "Invalid weight parameter", http.StatusBadRequest)
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID parameter", http.StatusBadRequest)
		return
	}

	if userIDInt <= 0 {
		http.Error(w, "User ID must be a positive integer", http.StatusBadRequest)
		return
	}

	// Get the users height from the db
	userBaselineRecord, err := baseline.GetBy(baseline.FIELD_UserID, userIDInt)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	height := userBaselineRecord.Height // Assuming height is stored in the user record
	if height.CMs <= 0 {
		http.Error(w, "User height is not set or invalid", http.StatusBadRequest)
		return
	}

	bmi := measures.BMI{}
	bmi.SetBMIByWeightAndHeight(userWeightFloat, height.CMs)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	attributes := map[string]interface{}{
		"bmi":         bmi.BMI,
		"description": bmi.Description,
		"glyph":       bmi.Glyph,
	}

	jsonBytes, err := BuildJSONAPIResponse("user", fmt.Sprintf("%d", userIDInt), attributes)
	if err != nil {
		panic(err)
	}

	// Add json Response to response body
	_, err = w.Write(jsonBytes)
	if err != nil {
		logHandler.ErrorLogger.Printf("Response Rendering Error %s", err.Error())
	}

	logHandler.EventLogger.Printf("BMI calculated for User ID %d: %.2f (%s)", userIDInt, bmi.Float(), bmi.Text())
}

func strip(in string) string {
	rtn := strings.Trim(in, "{")
	rtn = strings.Trim(rtn, "}")
	return rtn
}
