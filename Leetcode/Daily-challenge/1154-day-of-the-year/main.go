package dayoftheyear

import (
	"strconv"
)

func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	yearDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year != 1900 && year%4 == 0 {
		yearDays[1]++
	}
	yearDays = yearDays[:month-1]
	sum := 0
	for i := 0; i < len(yearDays); i++ {
		sum += yearDays[i]
	}
	return sum + day
}
