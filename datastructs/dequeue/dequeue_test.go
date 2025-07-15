package dequeue_test

import (
	"testing"

	"github.com/MordaTeam/go-toolbox/datastructs/dequeue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDequeue(t *testing.T) {
	d := dequeue.New[int]()

	d.PushHead(1)
	d.PushHead(2)
	d.PushTail(3)
	d.PushTail(4)

	t.Run("check_get", func(t *testing.T) {
		data, err := d.GetHead()
		require.NoError(t, err)
		require.Equal(t, 2, data)

		data, err = d.GetTail()
		require.NoError(t, err)
		require.Equal(t, 4, data)
	})

	expectedRes := []int{2, 1, 3, 4}

	t.Run("check_get_all", func(t *testing.T) {
		res := d.GetAll()
		assert.Equal(t, expectedRes, res)
		assert.Equal(t, uint(4), d.Size())
	})

	t.Run("check_pop_order", func(t *testing.T) {
		res := []int{}
		for range 4 {
			data, err := d.PopHead()
			require.NoError(t, err)
			res = append(res, data)
		}

		assert.Equal(t, expectedRes, res)
		assert.Equal(t, uint(0), d.Size())
	})

	t.Run("check_pop_on_empty", func(t *testing.T) {
		require.Equal(t, uint(0), d.Size())

		data, err := d.PopHead()
		require.Error(t, err)
		require.Empty(t, data)

		data, err = d.PopTail()
		require.Error(t, err)
		require.Empty(t, data)
	})

	t.Run("check_get_on_empty", func(t *testing.T) {
		require.Equal(t, uint(0), d.Size())

		data, err := d.GetHead()
		require.Error(t, err)
		require.Empty(t, data)

		data, err = d.GetTail()
		require.Error(t, err)
		require.Empty(t, data)
	})
}
