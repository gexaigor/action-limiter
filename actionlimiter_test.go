package limiter

import (
	"testing"
	"time"
)

func TestActionLimit_Wait(t *testing.T) {
	al := New(2, time.Second)

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
			if c != tc.expected {
				t.Error()
			}
		})
	}
}

func TestActionLimit_Try(t *testing.T) {
	al := New(2, time.Second)

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
			for i := 0; i < tc.quantityCalls; i++ {
				al.Wait()
			}
			if al.Try() != tc.expected {
				t.Error()
			}
		})
	}
}
