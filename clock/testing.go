package clock

// Deterministic helper useful for testing.

import (
	"sync"
	"time"
)

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

type Ticker struct {
	ticker *ticker
	mu     sync.Mutex
}

func NewTicker(startTime time.Time, interval time.Duration) *Ticker {
	return &Ticker{
		ticker: newTicker(startTime, interval),
	}
}

// Now returns the current time.
func (t *Ticker) Now() time.Time {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.ticker.now()
}

// Tick forwards the Ticker one tick interval.
func (t *Ticker) Tick() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.ticker.tick()
}

type AutoTicker struct {
	ticker *ticker
	mu     sync.Mutex
}

func NewAutoTicker(startTime time.Time, tickInterval time.Duration) *AutoTicker {
	return &AutoTicker{
		ticker: newTicker(startTime, tickInterval),
	}
}

func (t *AutoTicker) Now() time.Time {
	t.mu.Lock()
	defer t.mu.Unlock()

	now := t.ticker.now()
	t.ticker.tick()
	return now
}

type ticker struct {
	currentTime time.Time
	interval    time.Duration
}

func newTicker(startTime time.Time, interval time.Duration) *ticker {
	if startTime.IsZero() {
		panic("startTime can not be zero")
	}
	if interval <= 0 {
		panic("interval must be greater than 0")
	}
	return &ticker{
		currentTime: startTime,
		interval:    interval,
	}
}

func (t *ticker) now() time.Time {
	return t.currentTime
}

func (t *ticker) tick() {
	t.currentTime = t.currentTime.Add(t.interval)
}
