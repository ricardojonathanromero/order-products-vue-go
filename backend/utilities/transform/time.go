package transform

import "time"

func GetInSec(sec int) time.Duration {
	return time.Duration(sec) * time.Second
}
