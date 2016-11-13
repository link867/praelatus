package models

import "encoding/json"

func jsonString(i interface{}) string {
	b, e := json.MarshalIndent(i, "", "\t")
	if e != nil {
		return ""
	}

	return string(b)
}
