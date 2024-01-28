package transform

import "strconv"

func StrToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return v
}
