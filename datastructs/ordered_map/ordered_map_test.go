package ordered_map_test

import (
	"testing"

	"github.com/MordaTeam/go-toolbox/datastructs/comparer"
	"github.com/MordaTeam/go-toolbox/datastructs/ordered_map"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrderedMap(t *testing.T) {
	m := ordered_map.New[string, int](comparer.NewStringComparer())

	require.NoError(t, m.Set("foo", 5))
	require.NoError(t, m.Set("buz", 7))
	require.NoError(t, m.Set("bar", 10))
	require.NoError(t, m.Set("baz", 7))

	require.Equal(t, 4, m.Size())

	assert.Equal(t, []string{"bar", "baz", "buz", "foo"}, m.Keys())

	val, err := m.Get("foo")
	require.NoError(t, err)
	assert.Equal(t, 5, val)

	require.NoError(t, m.Remove("foo"))
	val, err = m.Get("foo")
	require.ErrorIs(t, err, ordered_map.KeyNotFoundError[string]{})
	assert.Empty(t, val)

	m.Clear()
	assert.Equal(t, 0, m.Size())
}
