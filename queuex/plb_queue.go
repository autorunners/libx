package queuex

import (
	"log"
)

type Prio int

const (
	HighPrimary Prio = iota
	NormalPrimary
	LowerPriority
)

type queueState int

const (
	empty queueState = iota
	caution
	warning
	alarm
	full
)

var _ PrioQueue = &plbQueue{}

// plbQueue Prio Leaky Bucket Queue
type plbQueue struct {
	name      string
	threshold plbThreshold
	len       int

	lists map[Prio]Queue
}

func NewPlbQueue(name string, threshold plbThreshold) PrioQueue {
	return &plbQueue{
		name:      name,
		threshold: threshold,
		len:       0,
		lists: map[Prio]Queue{
			HighPrimary:   NewQueue(),
			NormalPrimary: NewQueue(),
			LowerPriority: NewQueue(),
		},
	}
}

func (p *plbQueue) Push(i interface{}) error {
	return p.push(i, NormalPrimary)
}
func (p *plbQueue) PushPrio(i interface{}, prio Prio) error {
	return p.push(i, prio)
}

func (p *plbQueue) Pop() (interface{}, error) {
	if p.len <= 0 {
		return nil, ErrQueueEmpty
	}
	p.len--
	high := p.lists[HighPrimary]
	if i, err := high.Pop(); err == nil {
		return i, err
	}
	normal := p.lists[NormalPrimary]
	if i, err := normal.Pop(); err == nil {
		return i, err
	}
	lower := p.lists[LowerPriority]
	if i, err := lower.Pop(); err == nil {
		return i, err
	}
	return nil, ErrQueueWrong
}

func (p *plbQueue) push(i interface{}, prio Prio) error {
	state := p.threshold.check(p.len + 1)
	switch state {
	case full:
		return ErrQueueFull
	case alarm:
		if prio != HighPrimary {
			return ErrQueueAlarm
		}
	case warning:
		if prio == LowerPriority {
			return ErrQueueWarning
		}
	case caution:
		log.Println(p.len)
	}
	p.len++
	p.lists[prio].Push(i)
	return nil

}

type plbThreshold struct {
	alarm   int
	warning int
	caution int
}

func (p plbThreshold) check(num int) queueState {
	if num == 0 {
		return empty
	}
	if num < p.caution {
		return caution
	}
	if num < p.warning {
		return warning
	}
	if num < p.alarm {
		return alarm
	}
	return full
}
