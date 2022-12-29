package helper

import (
	"encoding/json"
)

func AppearJSON(info interface{}) string {
	var m = map[string]interface{}{
		"info": info,
	}
	byte, _ := json.Marshal(m)
	return string(byte)
}
