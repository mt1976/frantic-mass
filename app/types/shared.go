package types

import "strconv"

func IntToString(i int) string {
	if i == 0 {
		return ""
	}
	return strconv.Itoa(i)
}
