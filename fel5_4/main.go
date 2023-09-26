/*
4. Adott egy tömb ami csak 0, 1 és 2 elemeket tartalmazhat. Tervezzen algoritmust
és készítsen programot, ami lineáris időben és állandó tárhasználattal rendezi a tömböt.
Példa:
Input: { 0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0 }
Output: { 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2 }

4. Given an array that can only contain 0, 1 and 2 elements. Design an algorithm
and create a program that sorts the array in linear time complexity and in constant storage complexity
*/
package main

import "fmt"

func main() {
	sortedNums := sort([]int{0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0})
	fmt.Println(sortedNums)
}

func sort(nums []int) []int {
	numCounts := [3]int{}
	for _, num := range nums {
		numCounts[num]++
	}

	sortedNums := make([]int, 0, len(nums))
	for i, count := range numCounts {
		for j := 0; j < count; j++{
			sortedNums = append(sortedNums, i)
		}
	}
	return sortedNums
}
