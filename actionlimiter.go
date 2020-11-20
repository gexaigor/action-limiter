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

// NewActionLimiter returns the implementation of the ActionLimiter interface
func NewActionLimiter(limit int, delay time.Duration) ActionLimiter {
	ch := make(chan struct{}, limit)

	al := &actionLimit{
		ch: ch,
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

// Wait implements the limiter.ActionLimiter interface.
func (al *actionLimit) Wait() {
	<-al.ch
}
