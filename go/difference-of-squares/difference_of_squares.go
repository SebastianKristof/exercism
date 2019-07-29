/*
Package diffsquares teaches us that an efficient algorithm can make programs significantly faster.
The difference of the speed of execution compared to the iterative solutions (commented out)
is over 100-fold.
*/
package diffsquares

// SquareOfSum calculates the squared sum of the first n natural numbers.
func SquareOfSum(n int) int {
	// for i := 1; i <= n; i++ {
	// 	sum = sum + i
	// }
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of squares of the first n natural numbers.
func SumOfSquares(n int) int {
	// for i := 1; i <= n; i++ {
	// 	ssq += i * i
	// }
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between the squared sum and the sum of squares of the first n natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
