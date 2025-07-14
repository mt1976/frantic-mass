package functions

import (
	"time"

	"github.com/mt1976/frantic-mass/app/dao/dateIndex"
)

func Today() time.Time {
	return dateIndex.Today()
}

func Yesterday() time.Time {
	return dateIndex.Yesterday()
}

func Tomorrow() time.Time {
	return dateIndex.Tomorrow()
}

func GetToday() (int, dateIndex.DateIndex, error) {
	return dateIndex.GetToday()
}

func GetYesterday() (int, dateIndex.DateIndex, error) {
	return dateIndex.GetYesterday()
}

func GetTomorrow() (int, dateIndex.DateIndex, error) {
	return dateIndex.GetTomorrow()
}
