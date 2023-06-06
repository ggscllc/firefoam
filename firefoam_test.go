package firefoam

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLimitWithMut(t *testing.T) {
	newRateLimit := NewRateLimit(5, 500*time.Millisecond)
	assert.Equal(t, true, newRateLimit.LimitWithMut())  //1
	assert.Equal(t, true, newRateLimit.LimitWithMut())  //2
	assert.Equal(t, true, newRateLimit.LimitWithMut())  //3
	assert.Equal(t, true, newRateLimit.LimitWithMut())  //4
	assert.Equal(t, true, newRateLimit.LimitWithMut())  //5
	assert.Equal(t, false, newRateLimit.LimitWithMut()) // 6
	time.Sleep(6 * time.Second)
	assert.Equal(t, int64(0), newRateLimit.GetCurrentProcs())
}

func TestMultipleLimitWithMut(t *testing.T) {
	newRateLimit_1 := NewRateLimit(5, 500*time.Millisecond)
	newRateLimit_2 := NewRateLimit(2, 750*time.Millisecond)

	assert.Equal(t, true, newRateLimit_1.LimitWithMut()) //1
	assert.Equal(t, true, newRateLimit_1.LimitWithMut()) //2
	assert.Equal(t, true, newRateLimit_1.LimitWithMut()) //3
	assert.Equal(t, true, newRateLimit_1.LimitWithMut()) //4
	assert.Equal(t, true, newRateLimit_1.LimitWithMut()) //5

	assert.Equal(t, true, newRateLimit_2.LimitWithMut()) //1
	assert.Equal(t, true, newRateLimit_2.LimitWithMut()) //2

	assert.Equal(t, false, newRateLimit_1.LimitWithMut()) //6
	assert.Equal(t, false, newRateLimit_2.LimitWithMut()) //3

	time.Sleep(6 * time.Second)

	assert.Equal(t, int64(0), newRateLimit_1.GetCurrentProcs())
	assert.Equal(t, int64(0), newRateLimit_2.GetCurrentProcs())
}

func TestLimitWithAtomics(t *testing.T) {
	newRateLimit := NewRateLimit(5, 500*time.Millisecond)
	assert.Equal(t, true, newRateLimit.LimitWithAtomic())  //1
	assert.Equal(t, true, newRateLimit.LimitWithAtomic())  //2
	assert.Equal(t, true, newRateLimit.LimitWithAtomic())  //3
	assert.Equal(t, true, newRateLimit.LimitWithAtomic())  //4
	assert.Equal(t, true, newRateLimit.LimitWithAtomic())  //5
	assert.Equal(t, false, newRateLimit.LimitWithAtomic()) // 6
	time.Sleep(6 * time.Second)
	assert.Equal(t, int64(0), newRateLimit.GetCurrentProcs())
}

func TestMultipleLimitWithAtomics(t *testing.T) {
	newRateLimit_1 := NewRateLimit(5, 500*time.Millisecond)
	newRateLimit_2 := NewRateLimit(2, 750*time.Millisecond)

	assert.Equal(t, true, newRateLimit_1.LimitWithAtomic()) //1
	assert.Equal(t, true, newRateLimit_1.LimitWithAtomic()) //2
	assert.Equal(t, true, newRateLimit_1.LimitWithAtomic()) //3
	assert.Equal(t, true, newRateLimit_1.LimitWithAtomic()) //4
	assert.Equal(t, true, newRateLimit_1.LimitWithAtomic()) //5

	assert.Equal(t, true, newRateLimit_2.LimitWithAtomic()) //1
	assert.Equal(t, true, newRateLimit_2.LimitWithAtomic()) //2

	assert.Equal(t, false, newRateLimit_1.LimitWithAtomic()) //6
	assert.Equal(t, false, newRateLimit_2.LimitWithAtomic()) //3

	time.Sleep(6 * time.Second)

	assert.Equal(t, int64(0), newRateLimit_1.GetCurrentProcs())
	assert.Equal(t, int64(0), newRateLimit_2.GetCurrentProcs())

}
