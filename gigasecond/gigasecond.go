// Work out your gigasecond birthday
package gigasecond

import "time"

// testVersion declaration for exercism.io tests
const testVersion = 4

// Add a gigasecond to a time date
func AddGigasecond(tm time.Time) time.Time {
    return tm.Add(1e9 * time.Second)
}
