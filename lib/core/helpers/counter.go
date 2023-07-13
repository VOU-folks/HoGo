package helpers

import "sync/atomic"

type Counter struct {
	instance atomic.Int64
}

func (c *Counter) Reset() *Counter {
	c.instance.Store(0)
	return c
}

func (c *Counter) Value() int64 {
	return c.instance.Load()
}

func (c *Counter) Add(delta int64) int64 {
	return c.instance.Add(delta)
}

func (c *Counter) Subtract(delta int64) int64 {
	return c.instance.Add(delta * -1)
}

func (c *Counter) Incr() int64 {
	return c.Add(1)
}

func (c *Counter) Decr() int64 {
	return c.Subtract(1)
}

func NewCounter() *Counter {
	return (&Counter{}).Reset()
}
