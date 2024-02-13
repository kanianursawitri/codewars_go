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

/*
this code gives timeout on large n

func Pell(n int) *big.Int {

	if n == 0 {
		return big.NewInt(0)
	} else if n == 1 {
		return big.NewInt(1)
	}

	return big.NewInt(0).Add(Pell(n-1).Mul(Pell(n-1), big.NewInt(2)), Pell(n-2))
}
*/

func Pell(n int) *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)
	for i := 0; i < n; i++ {
		a, b = b, a.Add(a, b).Add(a, b)
	}
	return a
}

/*

Pell(4)

n=4
a,b = 0,1

i=0
-> a=0, b=1
-> a=0.add(0,1).add(a,b)
-> a=1.add(1,1)
-> a=2
-> b,a = a,b
-> b,a= 2,1

i=1
-> a=1, b=2
-> a=1.add(1,2).add(a,b)
-> a=3.add(3,2)
-> a=5
-> b,a = a,b
-> b,a= 5,2

i=2
-> a=2, b=5
-> a=2.add(2,5).add(a,b)
-> a=7.add(7,5)
-> a=12
-> b,a = a,b
-> b,a= 12,5

i=3
-> a=5, b=12
-> a=5.add(5,12).add(a,b)
-> a=17.add(17,12)
-> a=29
-> b,a = a,b
-> b,a= 29,12


return a
return 12


*/

func main() {
	fmt.Println(Pell(100))
}
