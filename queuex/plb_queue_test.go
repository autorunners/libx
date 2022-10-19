package queuex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPlbQueueNormal: Normal 类消息的Push 与 Pop
func TestPlbQueueNormal(t *testing.T) {
	tests := []struct {
		name  string
		args1 plbThreshold
		args2 []interface{}
	}{
		{"less", plbThreshold{6, 4, 2}, []interface{}{0, 1, 2, 3, 4, 5, 6}},
		{"just boundary condition", plbThreshold{6, 4, 2}, []interface{}{0, 1, 2, 3, 4, 5, 6}},
		{"number is caution", plbThreshold{6, 4, 2}, []interface{}{0, 1, 2, 3, 4, 5, 6}},
		{"number is caution", plbThreshold{6, 4, 2}, []interface{}{0, 1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		prioQueue := NewPlbQueue(tt.name, tt.args1)
		threshold := tt.args1.warning
		t.Run(tt.name, func(t *testing.T) {
			// Push: 在超过caution边界前，正常插入
			for i := 0; i < threshold; i++ {
				assert.Nil(t, prioQueue.Push(tt.args2[i]))
				//t.Log(prioQueue)
			}
			// Push: 在超过 caution 边缘后，Normal 类数据插入失败
			if threshold < len(tt.args2) {
				for i := threshold; i < len(tt.args2); i++ {
					assert.NotNil(t, prioQueue.Push(tt.args2[i]))
				}
			}
			// Pop: 正常取出消息
			for i := 0; i < threshold; i++ {
				item, err := prioQueue.Pop()
				assert.Nil(t, err)
				assert.Equal(t, tt.args2[i], item)
			}
			// Pop: 消息取完后，再取抛出失败
			item, err := prioQueue.Pop()
			assert.NotNil(t, err)
			assert.Nil(t, item)
		})
	}
}

// TestPlbQueuePiro: 指定优先级消息的Push 与 Pop
func TestPlbQueuePiro(t *testing.T) {
	tests := []struct {
		name      string
		args1     plbThreshold
		args2     []interface{}
		prio      Prio
		threshold int
	}{
		{"[normal]", plbThreshold{6, 4, 3}, []interface{}{1, 2, 3, 4, 5, 6, 7}, NormalPrimary, 4},
		{"[high]", plbThreshold{6, 4, 3}, []interface{}{1, 2, 3, 4, 5, 6, 7}, HighPrimary, 6},
		{"[lower]", plbThreshold{6, 4, 3}, []interface{}{1, 2, 3, 4, 4, 5, 6, 7}, LowerPriority, 3},
		{"number is caution", plbThreshold{6, 4, 3}, []interface{}{1, 2, 3, 4, 5, 6}, LowerPriority, 3},
	}
	for _, tt := range tests {
		prioQueue := NewPlbQueue(tt.name, tt.args1)
		t.Run(tt.name, func(t *testing.T) {
			// Push: 在超过caution边界前，正常插入
			for i := 0; i < tt.threshold; i++ {
				assert.Nil(t, prioQueue.PushPrio(tt.args2[i], tt.prio))
			}
			// Push: 在超过 caution 边缘后，指定类数据插入失败
			if tt.threshold < len(tt.args2) {
				for i := tt.threshold; i < len(tt.args2); i++ {
					assert.NotNil(t, prioQueue.PushPrio(tt.args2[len(tt.args2)-1], tt.prio))
				}
			}
			// Pop: 正常取出消息
			for i := 0; i < tt.threshold; i++ {
				item, err := prioQueue.Pop()
				assert.Nil(t, err)
				assert.Equal(t, tt.args2[i], item)
			}
			// Pop: 消息取完后，再取抛出失败
			item, err := prioQueue.Pop()
			assert.NotNil(t, err)
			assert.Nil(t, item)
		})
	}
}

func Test_plbThreshold_check(t *testing.T) {
	tests := []struct {
		name         string
		plbThreshold plbThreshold
		arg          int
		want         queueState
	}{
		{"empty", plbThreshold{6, 4, 2}, 0, empty},
		{"normal", plbThreshold{6, 4, 2}, 1, normal},
		{"normal", plbThreshold{6, 4, 2}, 2, caution},
		{"caution", plbThreshold{6, 4, 2}, 3, caution},
		{"warning", plbThreshold{6, 4, 2}, 4, warning},
		{"warning", plbThreshold{6, 4, 2}, 5, warning},
		{"full", plbThreshold{6, 4, 2}, 6, full},
		{"full2", plbThreshold{6, 4, 2}, 7, full},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.plbThreshold.check(tt.arg), "check(%v)", tt.arg)
		})
	}
}
