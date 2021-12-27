package srtp

import "sync/atomic"

type count32 struct {
	val int32
}

func (c *count32) Inc() int32 {
	return atomic.AddInt32((*int32)(&c.val), 1)
}

func (c *count32) Get() int32 {
	return atomic.LoadInt32((*int32)(&c.val))
}
