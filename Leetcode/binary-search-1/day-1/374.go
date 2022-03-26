package day1

func guess(num int) int

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
	left, right := 1, n
	for left < right {
		mid := (left + right) / 2
		num := guess(mid)
		if num == 0 {
			return mid
		} else if num < 0 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
