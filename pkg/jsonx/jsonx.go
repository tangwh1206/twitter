package jsonx

import "encoding/json"

func BeautifyDump(v interface{}) string {
	content, _ := json.MarshalIndent(v, "", "    ")
	return string(content)
}

func SafeDump(v interface{}) string {
	content, _ := json.Marshal(v)
	return string(content)
}
