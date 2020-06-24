package xtype

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMap(t *testing.T) {
	var src StringMap

	dst, err := json.Marshal(src)

	assert.NoError(t, err)
	assert.Equal(t, dst, []byte("null"))
}
