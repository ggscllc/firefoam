package firefoam

import (
	"sync"
	"time"
)

type RateLimit struct {
	duration    time.Duration
	maxProc     int
	currentProc int
	mut         *sync.RWMutex
}

func NewRateLimit(maxProc int, duration time.Duration) RateLimit {
	var mut sync.RWMutex

	return RateLimit{
		duration:    duration,
		maxProc:     maxProc,
		currentProc: 0,
		mut:         &mut,
	}
}

func (r *RateLimit) GetCurrentProcs() int {
	r.mut.RLock()
	defer r.mut.RUnlock()
	return r.currentProc
}

func (r *RateLimit) TakeItToTheLimit() bool {
	r.mut.Lock()
	defer r.mut.Unlock()

	if r.currentProc >= r.maxProc {
		return false
	}

	r.currentProc += 1

	time.AfterFunc(r.duration, func() {
		r.mut.Lock()
		defer r.mut.Unlock()
		r.currentProc -= 1
	})
	return true
}
