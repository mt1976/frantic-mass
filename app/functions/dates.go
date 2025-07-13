package functions

import (
	"time"

	"github.com/mt1976/frantic-mass/app/dao/dateIndex"
)

func Today() time.Time {
	return dateIndex.Today()
}
