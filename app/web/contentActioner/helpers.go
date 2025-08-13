package contentActioner

import "strconv"

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
