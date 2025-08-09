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

var BMIURI = "/bmi/calculate"                                                 // Define the URI for the BMI endpoint
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
	height := getUserHeight(userIDInt, w)

	bmi := measures.BMI{}
	bmi.SetBMIByWeightAndHeight(userWeightFloat, height.CMs)

	jsonBytes := generateBMIResponse(w, bmi, userIDInt)

	// Add json Response to response body
	_, err = w.Write(jsonBytes)
	if err != nil {
		logHandler.ErrorLogger.Printf("Response Rendering Error %s", err.Error())
	}

	logHandler.EventLogger.Printf("BMI calculated for User ID %d: %.2f (%s)", userIDInt, bmi.Float(), bmi.Text())
}

func generateBMIResponse(w http.ResponseWriter, bmi measures.BMI, userIDInt int) []byte {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	// Build the JSON API response
	attributes := map[string]interface{}{
		"bmi":         bmi.BMI,
		"description": bmi.Description,
		"glyph":       bmi.Glyph,
	}

	jsonBytes, err := BuildJSONAPIResponse("user", fmt.Sprintf("%d", userIDInt), attributes)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func getUserHeight(userIDInt int, w http.ResponseWriter) measures.Height {
	userBaselineRecord, err := baseline.GetBy(baseline.FIELD_UserID, userIDInt)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return measures.Height{}
	}

	height := userBaselineRecord.Height // Assuming height is stored in the user record
	if height.CMs <= 0 {
		http.Error(w, "User height is not set or invalid", http.StatusBadRequest)
		return measures.Height{}
	}
	return height
}

var BMIEnrichmentURI = "/bmi/enrichment"                                                                           // Define the URI for the BMI enrichment endpoint
var BMIEnrichmentUserWildcard = "{userID}"                                                                         // Define the user ID wildcard for enrichment
var BMIEnrichmentBMIWildcard = "{bmi}"                                                                             // Define the BMI wildcard for enrichment
var BMIEnrichmentWildcardURI = BMIEnrichmentURI + "/" + BMIEnrichmentUserWildcard + "/" + BMIEnrichmentBMIWildcard // Define the wildcard URI for the BMI enrichment endpoint

func BMIEnrichment(w http.ResponseWriter, r *http.Request) {
	logHandler.EventLogger.Printf("Entered BMI Enrichment")
	// Implementation for BMI enrichment goes here
	// just like the BMI except we will return the enrichment data based on the BMI value
	userID := chi.URLParam(r, strip(BMIEnrichmentUserWildcard))
	bmiValue := chi.URLParam(r, strip(BMIEnrichmentBMIWildcard))

	logHandler.InfoLogger.Printf("UserID [%v] BMI Value [%v]", userID, bmiValue)
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID parameter", http.StatusBadRequest)
		return
	}
	if userIDInt <= 0 {
		http.Error(w, "User ID must be a positive integer", http.StatusBadRequest)
		return
	}
	bmiFloat, err := strconv.ParseFloat(bmiValue, 64)
	if err != nil {
		http.Error(w, "Invalid BMI value parameter", http.StatusBadRequest)
		return
	}
	if bmiFloat <= 0 {
		http.Error(w, "BMI value must be a positive number", http.StatusBadRequest)
		return
	}

	bmi := measures.BMI{}
	bmi.Set(bmiFloat)

	jsonBytes := generateBMIResponse(w, bmi, userIDInt)

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
