package utils

import "encoding/json"

// convert a json.RawMessage to a string
func RawMessageToString(raw json.RawMessage) string {
	return string(raw)
}
