// Package gigasecond can be used to work out your gigasecond birthday
package gigasecond

import "time"

// testVersion declaration for exercism.io tests
const testVersion = 4

// AddGigasecond adds a gigasecond 10^9 to a time date
func AddGigasecond(tm time.Time) time.Time {
	return tm.Add(1e9 * time.Second)
}
