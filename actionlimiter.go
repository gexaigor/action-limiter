package limiter

import "time"

// ActionLimiter is the interface that wraps the basic Wait and Try methods.
type ActionLimiter interface {
	Wait()
	Try() bool
}

// actionLimit implements the ActionLimiter interface and allows you to limit actions
type actionLimit struct {
	ch       chan struct{}
	interval time.Duration
}

// New returns the implementation of the ActionLimiter interface
func New(limit int, interval time.Duration) ActionLimiter {
	al := &actionLimit{
		ch:       make(chan struct{}, limit),
		interval: interval,
	}
	al.fillChan()
	go func() {
		for {
			time.Sleep(al.interval)
			al.fillChan()
		}
	}()

	return al
}

func (al *actionLimit) fillChan() {
	for i := len(al.ch); i < cap(al.ch); i++ {
		al.ch <- struct{}{}
	}
}

// Wait implements the limiter.ActionLimiter interface.
// Wait method blocks the thread if more limit actions have already been performed.
func (al *actionLimit) Wait() {
	<-al.ch
}

// Try implements the limiter.ActionLimiter interface.
// Try method returns false if more limit actions have already been performed.
func (al *actionLimit) Try() bool {
	select {
	case <-al.ch:
		return true
	default:
		return false
	}
}
