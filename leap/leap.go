// Package leap provides some useful stuff to check if the year is a leap year.
package leap

// IsLeapYear checks if the year is a leap year
func IsLeapYear(year int) bool {
	leapYear := false
	// year is evenly divisible by 4
	if year % 4 == 0 {
		// and not evenly divisible by 100
		if year % 100 != 0 {
			leapYear = true
		} else {
			// unless the year is evenly divisible by 400
			if year % 400 == 0 {
				leapYear = true
			}
		}
	}
	return leapYear
}
