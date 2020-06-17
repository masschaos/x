package xtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gopkg.in/yaml.v2"
)

// Strings is string slice, can be varchar,text or json in mysql column
type Strings []string

// String print strings as json format: ["a","b"]
func (t Strings) String() string {
	if t == nil {
		return ""
	}
	tmp, _ := json.Marshal([]string(t))
	return string(tmp)
}

// Has a item or not
func (t Strings) Has(s string) bool {
	for _, item := range t {
		if item == s {
			return true
		}
	}
	return false
}

// Has another strings or not
func (t Strings) Contains(s Strings) bool {
	for _, ss := range s {
		if !t.Has(ss) {
			return false
		}
	}
	return true
}

// Intersect another strings or not
func (t Strings) Intersect(s Strings) bool {
	for _, ss := range s {
		if t.Has(ss) {
			return true
		}
	}
	return false
}

// SAdd add an item to set t and return the new set
func (t Strings) SAdd(s string) Strings {
	if t.Has(s) {
		return t
	}
	return append(t, s)
}

// Remove the items if it's value is s
func (t Strings) Remove(s string) Strings {
	tmp := make(Strings, 0, len(t))
	for _, tt := range t {
		if tt != s {
			tmp = append(tmp, tt)
		}
	}
	return tmp
}

// Union two strings
func (t Strings) Union(s Strings) Strings {
	tmp := t
	for _, ss := range s {
		tmp = tmp.SAdd(ss)
	}
	return tmp
}

// Sub s from t and return the new strings
func (t Strings) Sub(s Strings) Strings {
	tmp := t
	for _, ss := range s {
		tmp = tmp.Remove(ss)
	}
	return tmp
}

// MarshalJSON for json interface
func (t Strings) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]string(t))
}

// UnmarshalJSON for json interface
func (t *Strings) UnmarshalJSON(data []byte) error {
	var tmp []string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = tmp
	return nil
}

// MarshalYAML for yaml interface
func (t Strings) MarshalYAML() ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}
	return yaml.Marshal([]string(t))
}

// UnmarshalYAML for yaml interface
func (t *Strings) UnmarshalYAML(data []byte) error {
	var tmp []string
	if err := yaml.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = tmp
	return nil
}

// Scan implements the Scanner interface.
func (t *Strings) Scan(src interface{}) error {
	*t = make([]string, 0)
	if src == nil {
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("read json string array from DB failed")
	}
	if len(tmp) == 0 {
		return nil
	}
	return t.UnmarshalJSON(tmp)
}

// Value implements the driver Valuer interface.
func (t Strings) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return t.String(), nil
}
