package firefoam

import (
	"sync"
	"time"
)

type RateLimit struct {
	Duration    time.Duration
	MaxProc     int
	CurrentProc *int
	Mut         *sync.Mutex
}

func NewRateLimit(maxProc int, duration time.Duration) RateLimit {
	count := 0
	var mut sync.Mutex

	return RateLimit{
		Duration:    duration,
		MaxProc:     maxProc,
		CurrentProc: &count,
		Mut:         &mut,
	}
}

func (r *RateLimit) TakeItToTheLimit() bool {
	r.Mut.Lock()
	defer r.Mut.Unlock()

	if *r.CurrentProc >= r.MaxProc {
		return false
	}

	*r.CurrentProc += 1

	time.AfterFunc(r.Duration, func() {
		r.Mut.Lock()
		defer r.Mut.Unlock()
		*r.CurrentProc -= 1
	})
	return true
}
