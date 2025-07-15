package sliding_window_test

import (
	"testing"

	"github.com/MordaTeam/go-toolbox/datastructs/sliding_window"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindow(t *testing.T) {
	w := sliding_window.New[int](3)

	t.Run("check_data_on_empty", func(t *testing.T) {
		d := w.Data()
		assert.Empty(t, d)
	})

	t.Run("check_push_clear", func(t *testing.T) {
		for i := range 3 {
			w.Push(i)
		}
		assert.Equal(t, []int{0, 1, 2}, w.Data())
		w.Clear()
		assert.Empty(t, w.Data())
	})

	t.Run("check_push_overlimit", func(t *testing.T) {
		for i := range 10 {
			w.Push(i)
		}
		assert.Equal(t, []int{7, 8, 9}, w.Data())
	})
}
