package xtype

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSet(t *testing.T) {
	var src Set

	// testing json marshal nil set
	tmp, err := json.Marshal(src)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null"), tmp)

	// testing yaml marshal nil set
	tmp, err = yaml.Marshal(src)
	assert.NoError(t, err)
	assert.Equal(t, []byte("[]\n"), tmp)

	// testing json unmarshal
	source := []byte(`["a","b","c"]`)
	err = json.Unmarshal(source, &src)
	assert.NoError(t, err)
	assert.Equal(t, Set{"a", "b", "c"}, src)

	// testing json marshal
	tmp, err = json.Marshal(src)
	assert.NoError(t, err)
	assert.Equal(t, source, tmp)

	var s1 = Set{"c", "d", "e"}
	var s2 = Set{"d", "e", "f"}
	var s3 = Set{"a", "b"}

	assert.Equal(t, true, src.Has("a"))
	assert.Equal(t, false, src.Has("d"))
	assert.Equal(t, true, src.IsSuperSet(s3))
	assert.Equal(t, false, src.IsSuperSet(s2))
	assert.Equal(t, false, src.IsDisjoint(s1))
	assert.Equal(t, true, src.IsDisjoint(s2))
	assert.Equal(t, Set{"a", "b", "c"}, src.Add("a"))
	assert.Equal(t, Set{"a", "b", "c", "d"}, src.Add("d"))
	assert.Equal(t, Set{"a", "b", "c", "d", "e"}, src.Union(s1))
	assert.Equal(t, Set{"a", "b", "c", "d", "e", "f"}, src.Union(s2))
	assert.Equal(t, Set{"b", "c"}, src.Remove("a"))
	assert.Equal(t, Set{"a", "b"}, src.Sub(s1))
}
