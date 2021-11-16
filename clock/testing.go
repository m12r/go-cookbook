package clock

// Deterministic helper useful for testing.

import "time"

// StoppedAt is a function which returns a Func, which always returns the
// given time.Time. This is very useful for testing.
func StoppedAt(t time.Time) Func {
	return func() time.Time {
		return t
	}
}

// StartedAt is a function which returns a Func, which starts the clock at
// the given time.Time. This is useful for testing, when you need a running
// clock, but you want to control, when the clock started.
func StartedAt(t time.Time) Func {
	offset := t.Sub(time.Now())
	return func() time.Time {
		return time.Now().Add(offset)
	}
}
