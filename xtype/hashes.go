package xtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Hashes is hash slice
type Hashes []uint64

func (t Hashes) String() string {
	if t == nil {
		return ""
	}
	tmp, _ := json.Marshal([]uint64(t))
	return string(tmp)
}

// Has a item or not
func (t Hashes) Has(s uint64) bool {
	for _, item := range t {
		if item == s {
			return true
		}
	}
	return false
}

// MarshalJSON 转换为json类型
func (t Hashes) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]uint64(t))
}

// UnmarshalJSON 不做处理
func (t *Hashes) UnmarshalJSON(data []byte) error {
	var tmp []uint64
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = tmp
	return nil
}

// Scan implements the Scanner interface.
func (t *Hashes) Scan(src interface{}) error {
	*t = make([]uint64, 0)
	if src == nil {
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("read json int array from DB failed")
	}
	if len(tmp) == 0 {
		return nil
	}
	return t.UnmarshalJSON(tmp)
}

// Value implements the driver Valuer interface.
func (t Hashes) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return t.String(), nil
}
