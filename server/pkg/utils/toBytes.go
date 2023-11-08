package utils

import "encoding/json"

func toBytes(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
