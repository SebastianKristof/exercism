package diffsquares

func SquareOfSum(n int) (sqs int) {
	for i := 1; i <= n; i++ {
		sqs = sqs + i
	}
	sqs = sqs * sqs
	return
}

func SumOfSquares(n int) (ssq int) {
	// for i := 1; i <= n; i++ {
	// 	ssq += i * i
	// }
	ssq = n * (n + 1) * (2*n + 1) / 6
	return
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
