package xtype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
	var src = Strings{"a", "b", "c"}
	var s1 = Strings{"c", "d", "e"}
	var s2 = Strings{"d", "e", "f"}
	var s3 = Strings{"a", "b"}

	assert.Equal(t, src.Has("a"), true)
	assert.Equal(t, src.Has("d"), false)
	assert.Equal(t, src.Contains(s3), true)
	assert.Equal(t, src.Contains(s2), false)
	assert.Equal(t, src.Intersect(s1), true)
	assert.Equal(t, src.Intersect(s2), false)
	assert.Equal(t, src.SAdd("a"), Strings{"a", "b", "c"})
	assert.Equal(t, src.SAdd("d"), Strings{"a", "b", "c", "d"})
	assert.Equal(t, src.Union(s1), Strings{"a", "b", "c", "d", "e"})
	assert.Equal(t, src.Union(s2), Strings{"a", "b", "c", "d", "e", "f"})
	assert.Equal(t, src.Remove("a"), Strings{"b", "c"})
	assert.Equal(t, src.Sub(s1), Strings{"a", "b"})
}
