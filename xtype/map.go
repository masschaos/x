package xtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// StringMap is name of map[string]string for custom db driver
type StringMap map[string]string

// String print strings as json format: {"a":"b"}, use this method in printing only
func (m StringMap) String() string {
	if m == nil {
		return "null"
	}
	tmp, _ := json.Marshal(m)
	return string(tmp)
}

// Scan implements the Scanner interface.
func (m *StringMap) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	// check src
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("read json object from DB failed")
	}
	if len(tmp) == 0 {
		return nil
	}
	// create m if it is nil
	if m == nil {
		*m = make(map[string]string)
	}
	return json.Unmarshal(tmp, m)
}

// Value implements the driver Valuer interface.
func (m StringMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	tmp, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("save json to db error: %w", err)
	}
	return string(tmp), nil
}
