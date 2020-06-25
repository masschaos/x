package xtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// Set is a string set, can be varchar,text or json array in mysql column
type Set []string

// NewSet init a set from items
func NewSet(items ...string) Set {
	var s = Set(make([]string, 0))
	for _, item := range items {
		s = s.Add(item)
	}
	return s
}

// String print strings as json format: ["a","b"], use this method in printing only
func (t Set) String() string {
	if t == nil {
		return ""
	}
	tmp, _ := json.Marshal([]string(t))
	return string(tmp)
}

// Has a item or not
func (t Set) Has(s string) bool {
	for _, item := range t {
		if item == s {
			return true
		}
	}
	return false
}

// IsSuperSet of s or not
func (t Set) IsSuperSet(s Set) bool {
	for _, ss := range s {
		if !t.Has(ss) {
			return false
		}
	}
	return true
}

// IsSubSet of s or not
func (t Set) IsSubSet(s Set) bool {
	for _, tt := range t {
		if !s.Has(tt) {
			return false
		}
	}
	return true
}

// IsDisjoint returns whether two sets have a intersection(false) or not(true)
func (t Set) IsDisjoint(s Set) bool {
	for _, ss := range s {
		if t.Has(ss) {
			return false
		}
	}
	return true
}

// Add an item to set t and return the new set
func (t Set) Add(s string) Set {
	if t.Has(s) {
		return t
	}
	return append(t, s)
}

// Remove the specified item and return the new set
func (t Set) Remove(s string) Set {
	tmp := make(Set, 0, len(t))
	for _, tt := range t {
		if tt != s {
			tmp = append(tmp, tt)
		}
	}
	return tmp
}

// Union two strings
func (t Set) Union(s Set) Set {
	tmp := t
	for _, ss := range s {
		tmp = tmp.Add(ss)
	}
	return tmp
}

// Sub s from t and return the new strings
func (t Set) Sub(s Set) Set {
	tmp := t
	for _, ss := range s {
		tmp = tmp.Remove(ss)
	}
	return tmp
}

// Scan implements the Scanner interface.
func (t *Set) Scan(src interface{}) error {
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
	return json.Unmarshal(tmp, t)
}

// Value implements the driver Valuer interface.
func (t Set) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	tmp, err := json.Marshal([]string(t))
	if err != nil {
		return nil, fmt.Errorf("save json to db error: %w", err)
	}
	return string(tmp), nil
}
