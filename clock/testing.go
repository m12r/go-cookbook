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
