package clock

// A clock abstraction for testable functionality dependant on time.Now.

import "time"

// Now a thin wrapper for time.Now, which implements the Clock interface.
var Now = Func(time.Now)

// Clock abstracts an advancing clock.
type Clock interface {
	// Now returns the current time.Time of the implementation, which is not
	// necessarily the same as the actual wall clock time.
	Now() time.Time
}

// Func a function which implements the Clock interface.
type Func func() time.Time

// Now returns the current time.Time of the Func implementation, which is not
// necessarily the same as the actual wall clock time.
func (fn Func) Now() time.Time {
	return fn()
}
