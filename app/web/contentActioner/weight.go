package contentActioner

import (
	"net/http"

	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

func NewWeightLogEntry(w http.ResponseWriter, r *http.Request, userID int) (contentProvider.WeightView, error) {
	return contentProvider.WeightView{}, nil
}

func UpdateWeightLogEntry(w http.ResponseWriter, r *http.Request, userID int, weightID int) (contentProvider.WeightView, error) {
	return contentProvider.WeightView{}, nil
}
