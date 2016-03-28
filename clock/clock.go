// The clock package can be used to create a new 24h clock, add minutes to it
// and get a String representation of the Clock's 24hr time
package clock

import "fmt"

// testVersion is used to ensure the tests match this version of the exercism kata
const testVersion = 3

// Clock will store the no. of minutes that represent the clock time
type Clock int

// Create a new clock by providing integer hours and minutes, i.e New(3, 25) = "03:25"
// We normalize negative minutes to 24hr clock
func New(hours, minutes int) Clock {
	retClock := Clock((hours*60 + minutes) % 1440)
	if retClock < 0 {
		retClock += 1440
	}
	return retClock
}

// Get a string representation of the Clock's time in the format "hh:mm"
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add minutes to the Clocks time
// We normalize negative minutes to 24hr clock
func (c Clock) Add(minutes int) Clock {
	retClock := (c + Clock(minutes)) % 1440
	if retClock < 0 {
		retClock += 1440
	}
	return retClock
}
