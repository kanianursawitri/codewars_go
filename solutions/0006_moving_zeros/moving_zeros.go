package main

import "fmt"

/*

Write an algorithm that takes an array and moves all of the zeros to the end, preserving the order of the other elements.

MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }

*/

/*
another solution:

func MoveZeros(arr []int) []int {
 	res:= make([]int,len(arr))
	ind:=0
	for i:=0;i<len(arr);i++{
		if arr[i]!=0{
			res[ind]=arr[i]
			ind++
		}
	}
	return res
}

*/

func MoveZeros(arr []int) []int {
	// TODO: Program me
	result := make([]int, len(arr))
	var zeros []int

	for _, item := range arr {
		if item == 0 {
			zeros = append(zeros, item)
		} else {
			result = append(result, item)
		}
	}

	result = append(result, zeros...)

	return result
}

func main() {
	fmt.Println(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))
}
