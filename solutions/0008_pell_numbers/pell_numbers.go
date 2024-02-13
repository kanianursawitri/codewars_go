package main

/*

The Pell sequence is the sequence of integers defined by the initial values

P(0) = 0, P(1) = 1
and the recurrence relation

P(n) = 2 * P(n-1) + P(n-2)
The first few values of P(n) are

0, 1, 2, 5, 12, 29, 70, 169, 408, 985, 2378, 5741, 13860, 33461, 80782, 195025, 470832, ..


Task
Your task is to return the nth Pell number

*/

import (
	"fmt"
	"math/big"
)

func Pell(n int) *big.Int {

	if n == 0 {
		return big.NewInt(0)
	} else if n == 1 {
		return big.NewInt(1)
	}

	return big.NewInt(0).Add(Pell(n-1).Mul(Pell(n-1), big.NewInt(2)), Pell(n-2))
}

/*

TODO: revamp to handle timeout on large n



*/

func main() {
	fmt.Println(Pell(100))
}
