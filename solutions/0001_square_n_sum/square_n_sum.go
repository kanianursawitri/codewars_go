/*
Complete the square sum function so that it squares each number passed into it and then sums the results together.
*/

package main

import "fmt"

func SquareSum(numbers []int) int {
	// your code here
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	return sum
}

func main() {
	fmt.Println(SquareSum([]int{0, 3, 4, 5}))
}
