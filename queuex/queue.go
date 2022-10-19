package queuex

import (
	"container/list"
	"errors"
)

var (
	ErrQueueFull    = errors.New("queue is full")
	ErrQueueEmpty   = errors.New("queue is empty")
	ErrQueueWarning = errors.New("queue is in warning")
	ErrQueueCaution = errors.New("queue is in caution")

	ErrQueueWrong = errors.New("queue is wrong: shold not exist")
)

type PrioQueue interface {
	Push(interface{}) error
	PushPrio(interface{}, Prio) error
	Pop() (interface{}, error)
}

type Queue interface {
	Push(interface{}) error
	Pop() (interface{}, error)
}

func NewQueue() Queue {
	return &queue{List: new(list.List)}
}

type queue struct {
	List *list.List
}

var _ Queue = new(queue)

func (q queue) Push(i interface{}) error {
	q.List.PushFront(i)
	return nil
}

func (q queue) Pop() (interface{}, error) {
	if q.List.Len() <= 0 {
		return nil, ErrQueueEmpty
	}
	return q.List.Remove(q.List.Back()), nil
}
