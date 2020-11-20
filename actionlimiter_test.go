package limiter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestActionLimit_Wait(t *testing.T) {
	al := &actionLimit{
		ch: make(chan struct{}, 2),
	}
	go func() {
		for {
			for i := len(al.ch); i < cap(al.ch); i++ {
				al.ch <- struct{}{}
			}
			time.Sleep(time.Second)
		}
	}()

	testCases := []struct {
		name          string
		quantityCalls int
		expected      bool
	}{
		{
			name:          "one call",
			quantityCalls: 1,
			expected:      true,
		},
		{
			name:          "two calls",
			quantityCalls: 2,
			expected:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := true
			for i := 0; i < tc.quantityCalls; i++ {
				al.Wait()
			}
			go func() {
				time.Sleep(time.Millisecond * 200)
				c = false
			}()
			al.Wait()
			assert.Equal(t, c, tc.expected)
		})
	}
}
