package queuex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_queue(t *testing.T) {
	tests := []struct {
		name  string
		queue Queue
		args  interface{}
	}{
		{"integer", NewQueue(), 1},
		{"string", NewQueue(), "hello world"},
		{"array", NewQueue(), []string{"hello", "world"}},
		{"map", NewQueue(), map[string]string{"hello": "world"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Nil(t, tt.queue.Push(tt.args))
			i, err := tt.queue.Pop()
			assert.Nil(t, err)
			assert.Equal(t, tt.args, i)
		})
	}
}
