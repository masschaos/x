package xtype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
	var src = Set{"a", "b", "c"}
	var s1 = Set{"c", "d", "e"}
	var s2 = Set{"d", "e", "f"}
	var s3 = Set{"a", "b"}

	assert.Equal(t, src.Has("a"), true)
	assert.Equal(t, src.Has("d"), false)
	assert.Equal(t, src.IsSuperSet(s3), true)
	assert.Equal(t, src.IsSuperSet(s2), false)
	assert.Equal(t, src.IsDisjoint(s1), false)
	assert.Equal(t, src.IsDisjoint(s2), true)
	assert.Equal(t, src.Add("a"), Set{"a", "b", "c"})
	assert.Equal(t, src.Add("d"), Set{"a", "b", "c", "d"})
	assert.Equal(t, src.Union(s1), Set{"a", "b", "c", "d", "e"})
	assert.Equal(t, src.Union(s2), Set{"a", "b", "c", "d", "e", "f"})
	assert.Equal(t, src.Remove("a"), Set{"b", "c"})
	assert.Equal(t, src.Sub(s1), Set{"a", "b"})
}
