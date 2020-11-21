package limiter

import "time"

// ActionLimiter is the interface that wraps the basic Wait method.
// Wait method blocks the thread if more limit actions have already been performed
type ActionLimiter interface {
	Wait()
}

// actionLimit implements the ActionLimiter interface and allows you to limit actions
type actionLimit struct {
	ch chan struct{}
}

// New returns the implementation of the ActionLimiter interface
func New(limit int, delay time.Duration) ActionLimiter {
	al := &actionLimit{
		ch: make(chan struct{}, limit),
	}
	go func() {
		for {
			for i := len(al.ch); i < cap(al.ch); i++ {
				al.ch <- struct{}{}
			}
			time.Sleep(delay)
		}
	}()

	return al
}

// Wait implements the limiter.ActionLimiter interface
func (al *actionLimit) Wait() {
	<-al.ch
}
