package priority_queue_test

import (
	"testing"

	"github.com/MordaTeam/go-toolbox/datastructs/priority_queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPriorityQueue(t *testing.T) {
	pq := priority_queue.New[string]()

	pq.Push(priority_queue.Node[string]{Payload: "foo", Priority: 5})
	pq.Push(priority_queue.Node[string]{Payload: "buz", Priority: 7})
	pq.Push(priority_queue.Node[string]{Payload: "bar", Priority: 10})
	pq.Push(priority_queue.Node[string]{Payload: "baz", Priority: 7})

	require.Equal(t, 4, pq.Size())

	lastPrior := 10
	for pq.Size() > 0 {
		el, err := pq.Pop()
		require.NoError(t, err)
		assert.LessOrEqual(t, int(el.Priority), lastPrior)
		lastPrior = int(el.Priority)
	}
}
