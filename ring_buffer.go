package godtaix

import (
	"errors"
)

var (
	ErrRingBufferIsEmpty = errors.New("ring buffer is empty")
)

// RingBuffer is ring buffer
type RingBuffer struct {
	buf        []interface{} // buf队列
	size       int
	readIndex  int
	writeIndex int
}

func NewRingBuffer(size int) *RingBuffer {
	initialSize := 8
	for initialSize < size {
		initialSize <<= 1
	}
	r := &RingBuffer{
		size: initialSize,
	}
	r.buf = make([]interface{}, r.size)
	return r
}

// IsEmpty judge ring buffer is empty
func (r *RingBuffer) IsEmpty() bool {
	return r.readIndex == r.writeIndex
}

// Read read value
func (r *RingBuffer) Read() (value interface{}, err error) {
	if r.IsEmpty() {
		err = ErrRingBufferIsEmpty
		return
	}
	value = r.buf[r.readIndex]
	r.readIndex = (r.readIndex + 1) % r.size
	return
}

// Expand expand ring buffer size
func (c *RingBuffer) expand() {
	newSize := c.size * 2
	c.size = newSize
	newBuff := make([]interface{}, c.size)
	copy(newBuff, c.buf)
	c.buf = newBuff
}

// Write write value
func (r *RingBuffer) Write(value interface{}) {
	r.buf[r.writeIndex] = value
	newIndex := (r.writeIndex + 1) % r.size
	if newIndex == r.readIndex {
		r.expand()
	}
	r.writeIndex = (r.writeIndex + 1) % r.size
}

// Capacity
func (r *RingBuffer) Capacity() int {
	return r.size
}

// Len
func (r *RingBuffer) Len() int {
	if r.IsEmpty() {
		return 0
	}
	if r.readIndex < r.writeIndex {
		return r.writeIndex - r.readIndex
	}
	return r.size - r.readIndex + r.writeIndex
}

// Reset
func (r *RingBuffer) Reset() {
	r.readIndex = 0
	r.writeIndex = 0
}
