/*
Package leap includes functionality for calculation of leap years.
Leap days and hours may be implemented in the future.
*/
package leap

import "log"

// IsLeapYear takes a year as a positive integer
// and returns a boolean representing whether this year is leap or not.
func IsLeapYear(year int) bool {

	if year < 0 {
		log.Fatal("sorry, A.D. years only")
	}

	// every year that is evenly divisible by 4
	//   except every year that is evenly divisible by 100
	//     unless the year is also evenly divisible by 400
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}

	return false
}
