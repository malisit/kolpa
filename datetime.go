package kolpa

import (
	"math/rand"
	"time"
)

// DateTimeBetween returns a time.Time typed variable which is a pseudo-random
// datetime that points to a date and time pair between given datetime interval.
func (g *Generator) DateTimeBetween(afterObj time.Time, beforeObj time.Time) time.Time {
	afterUnix := afterObj.Unix()
	beforeUnix := beforeObj.Unix()

	diff := beforeUnix - afterUnix

	rand.Seed(time.Now().UTC().UnixNano())
	sec := rand.Int63n(diff) + afterUnix

	return time.Unix(sec, 0)
}

// DateTimeBetweenWithString is same with DateTimeBetween but takes interval as strings with time.RFC3339 builtin format.
func (g *Generator) DateTimeBetweenWithString(after string, before string) time.Time {
	afterObj, err := time.Parse(time.RFC3339, after)
	beforeObj, err2 := time.Parse(time.RFC3339, before)

	if err != nil && err2 != nil {
		return time.Time{}
	}

	return g.DateTimeBetween(afterObj, beforeObj)
}

// DateTimeAfter returns a time.Time typed variable which is a pseudo-random datetime that points
// to a date and time pair that is after the given datetime parameter.
func (g *Generator) DateTimeAfter(afterObj time.Time) time.Time {
	return g.DateTimeBetween(afterObj, time.Now())
}

// DateTimeAfterWithString is same with DateTimeAfter but takes datetime parameter as string with time.RFC3339 builtin format.
func (g *Generator) DateTimeAfterWithString(after string) time.Time {
	afterObj, err := time.Parse(time.RFC3339, after)

	if err != nil {
		return time.Time{}
	}

	return g.DateTimeBetween(afterObj, time.Now())
}

// DateTimeBefore returns a time.Time typed variable which is a pseudo-random datetime that points
// to a date and time pair that is before the given datetime parameter.
func (g *Generator) DateTimeBefore(beforeObj time.Time) time.Time {
	return g.DateTimeBetween(time.Date(0, 1, 0, 0, 0, 0, 0, time.UTC), beforeObj)
}

// DateTimeBeforeWithString is same with DateTimeBefore but takes datetime parameter as string with time.RFC3339 builtin format.
func (g *Generator) DateTimeBeforeWithString(before string) time.Time {
	beforeObj, err := time.Parse(time.RFC3339, before)
	if err != nil {
		return time.Time{}
	}

	return g.DateTimeBetween(time.Date(0, 1, 0, 0, 0, 0, 0, time.UTC), beforeObj)
}

// DateFormatter formats given datetime string with given layout.
// To define your own format, write down what the reference time would look like formatted your way;
// see the values of constants like ANSIC, StampMicro or Kitchen of time package for examples.
func (g *Generator) DateFormatter(layout, datetime string) string {
	t, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", datetime)

	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	return t.Format(layout)
}
