package main

import (
	"fmt"
	"time"
)

func timeFunc() {
	p := fmt.Println
	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	p(then.Year())
	p(then.Minute())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Seconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
