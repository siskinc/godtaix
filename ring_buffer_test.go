package godtaix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func find2nNumber(n int) int {
	var m = n
	m |= m >> 1
	m |= m >> 2
	m |= m >> 4
	m |= m >> 8
	m |= m >> 16
	m += 1 //大于N的最小的2的N次方
	// n = n >> 1 //小于N的最大的2的N次方
	if m>>1 == n {
		return n
	}
	return m
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestNewRingBuffer(t *testing.T) {
	for i := 0; i <= 10000; i++ {
		r := NewRingBuffer(i)
		topNumber := find2nNumber(i)
		topNumber = maxInt(8, topNumber)
		assert.Equal(t, r.size, maxInt(8, topNumber), "min ring buff size is 8 and ring buff size is 2^n")
	}
}

func TestRingBufferAdd(t *testing.T) {
	r := NewRingBuffer(0)
	for i := 0; i < 1000; i++ {
		r.Write(i)
		t.Logf("buf data: %v", r.buf)
	}
}

func TestRingBufferRead(t *testing.T) {
	r := NewRingBuffer(0)
	number := 10
	for i := 0; i < number; i++ {
		r.Write(i)
	}
	for i := 0; i < number+1; i++ {
		v, err := r.Read()
		if err != nil {
			t.Errorf("read have an err； %v", err)
			break
		}
		t.Logf("read value is %v", v)
		t.Logf("rest buf is %v", r.buf)
	}
}
