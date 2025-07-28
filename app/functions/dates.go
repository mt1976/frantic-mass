package functions

import (
	"fmt"
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

// CheckSameWeekday returns nil if date1 and date2 are on the same weekday.
// Otherwise, returns an error indicating the mismatch.
func CheckSameWeekday(date1, date2 time.Time) error {
	if date1.Weekday() != date2.Weekday() {
		return fmt.Errorf("weekday mismatch: %s vs %s", date1.Weekday(), date2.Weekday())
	}
	return nil
}
