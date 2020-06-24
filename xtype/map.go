package xtype

import "encoding/json"

// StringMap is name of map[string]string for custom db driver
type StringMap map[string]string

// String print strings as json format: ["a","b"]
func (m StringMap) String() string {
	if m == nil {
		return "{}"
	}
	tmp, _ := json.Marshal(m)
	return string(tmp)
}
