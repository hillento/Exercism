package diffsquares


//Difference takes the difference between the square of the sums and the sum of squares
func Difference(n int)int{
	return SquareOfSum(n) - SumOfSquares(n) 
}

//SumOfSquares uses the formula to add 1^2 + 2^2 + 3^2....+n^2
func SumOfSquares(n int) int {
	sumOSqrs := (n*(2*n+1)*(n+1))/6
	return sumOSqrs
}

//SquareOfSum uses the Gauss approach to sum all numbers from 1 to n and then squares the sum
func SquareOfSum(n int)int{
	sum := (n*(n+1))/2
	sqr := sum * sum
	return sqr
}


