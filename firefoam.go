// © Golden Grizzly Software Consulting LLC 2023
package firefoam

import (
	"sync"
	"sync/atomic"
	"time"
)

type RateLimiter struct {
	duration    time.Duration
	maxProc     int64
	currentProc int64
	mut         *sync.RWMutex
}

func NewRateLimit(maxProc int64, duration time.Duration) RateLimiter {
	var mut sync.RWMutex

	if maxProc < 1 {
		panic("maxProc was less than 1")
	}

	return RateLimiter{
		duration:    duration,
		maxProc:     maxProc,
		currentProc: 0,
		mut:         &mut,
	}
}

func (r *RateLimiter) GetCurrentProcs() int64 {
	return atomic.LoadInt64(&r.currentProc)
}

// 'Hard limit'
func (r *RateLimiter) LimitWithMut() bool {
	r.mut.Lock()
	defer r.mut.Unlock()

	if r.currentProc >= r.maxProc {
		return false
	}

	r.currentProc += 1

	r.decAfterTime()

	return true
}

// 'Soft' limit
func (r *RateLimiter) LimitWithAtomic() bool {
	if atomic.LoadInt64(&r.currentProc) >= atomic.LoadInt64(&r.maxProc) {
		return false
	}

	atomic.AddInt64(&r.currentProc, 1)

	r.decAfterTime()

	return true
}

func (r *RateLimiter) decAfterTime() {
	time.AfterFunc(r.duration, func() {
		atomic.AddInt64(&r.currentProc, -1)
	})
}
