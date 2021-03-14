# Clock

This is a thin wrapper around `time.Now` to make functionality dependent on it
testable.

## Example

It is not possible to test the following service, because it directly uses
`time.Now()`.

```go
package some

import (
    "time"
)

type UntestableService struct { 
    // some fields
}

func (us *UntestableService) IsOpen() bool {
    now := time.Now()
    hour := now.Hour()
    
    switch now.Weekday() {
    case time.Monday, time.Tuesday, time.Wednesday, time.Thursday:
    	return 10 <= hour && hour < 17
    case time.Friday:
    	return 10 <= hour && hour < 15
    }
    return false
}
```

The service below is testable, as it injects a `clock.Clock`.

```go
package some

import (
    "time"
	
    "github.com/m12r/go-cookbook/clock"
)

type TestableService struct {
    Clock clock.Clock
    // some more fields
}

func (ts *TestableService) IsOpen() bool {
    now := ts.Clock.Now()
    hour := now.Hour()

    switch now.Weekday() {
    case time.Monday, time.Tuesday, time.Wednesday, time.Thursday:
        return 10 <= hour && hour < 17
    case time.Friday:
        return 10 <= hour && hour < 15
    }
    return false
}
```

Integration tests of `TestableService`:

```go
package some_test

import (
    "testing"
    "time"

    "github.com/m12r/go-cookbook/clock"
)

func TestIsOpen(t *testing.T) {
    testCases := []struct {
        name     string
        at       time.Time
        expected bool
    }{
        {
            name:     "open on monday",
            at:       timeAt(t, "2021-03-08 12:00"),
            expected: true,
        },
        // more test cases here
        {
            name:     "closed on sunday",
            at:       timeAt(t, "2021-03-14 12:00"),
            expected: false,
        },
    }
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            ts := &some.TestableService{Clock: clock.StoppedAt(tc.at)}
            got := ts.IsOpen()
            if tc.expected != got {
                t.Errorf("expected %v, got %v", tc.expected, got)
            }
        })
    }
}

func timeAt(t *testing.T, datetime string) time.Time {
    t.Helper()
	
    ti, err := time.Parse("2006-02-01 15:04", datetime)
    if err != nil {
        t.Fatalf("cannot parse supplied datetime: %q: %v", datetime, err)
    }
    return ti
}
```

Have fun and enjoy coding.

See you next time :)

---

Copyright © 2021, [Matthias Endler][me]. All rights reserved.


[me]: https://m12r.at