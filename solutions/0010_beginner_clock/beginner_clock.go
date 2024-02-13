package main

import "fmt"

/*

Clock shows h hours, m minutes and s seconds after midnight.

Your task is to write a function which returns the time since midnight in milliseconds.

Example:
h = 0
m = 1
s = 1

result = 61000
Input constraints:

0 <= h <= 23
0 <= m <= 59
0 <= s <= 59

*/

func Past(h, m, s int) int {
	// your code here
	return (h*3600 + m*60 + s) * 1000
}

func main() {
	fmt.Println(Past(0, 1, 1))
}
