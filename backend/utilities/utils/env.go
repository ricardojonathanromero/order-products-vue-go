package utils

import "os"

func GetEnv(key, value string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}

	return value
}
