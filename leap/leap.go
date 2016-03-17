// Package leap provides a leap year checking function
package leap

const testVersion = 2

//IsLeapYear returns true if year given is a leap year, false if it isn't
func IsLeapYear(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
