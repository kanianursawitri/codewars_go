package main

import (
	"fmt"
	"sort"
)

/*

You will be given a list of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars) and then return the first value.

The returned value must be a string, and have "***" between each of its letters.

You should not remove or add elements from/to the array.

*/

/*

func TwoSort(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
}

*/

func TwoSort(arr []string) string {
	var result string

	sort.Strings(arr)

	for idx, item := range arr[0] {
		result += string(item)
		if idx < len(arr[0])-1 {
			result += "***"
		}
	}

	return result
}

func main() {
	fmt.Println(TwoSort([]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}))
}
