package queuex

import (
	"container/list"
	"errors"
	"sync"
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
	return &queue{
		List: new(list.List),
		Lock: &sync.Mutex{},
	}
}

type queue struct {
	List *list.List
	Lock *sync.Mutex
}

var _ Queue = new(queue)

func (q queue) Push(i interface{}) error {
	defer q.Lock.Unlock()
	q.Lock.Lock()
	q.List.PushFront(i)
	return nil
}

func (q queue) Pop() (interface{}, error) {
	defer q.Lock.Unlock()
	q.Lock.Lock()
	if q.List.Len() <= 0 {
		return nil, ErrQueueEmpty
	}
	return q.List.Remove(q.List.Back()), nil
}
