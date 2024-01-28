package transform

import "encoding/json"

func StructToString(data any) string {
	m, _ := json.Marshal(data)
	return string(m)
}
