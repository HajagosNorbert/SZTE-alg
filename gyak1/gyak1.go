package main

import "fmt"

func main() {
	permutate([]rune("abc"))
	// combine([]rune("abcd"), 2)
	// subset([]rune("abc"))
	// subsetIter([]rune("abc"))
}

// Produce every possible permutation of the provided characters
func permutate(chars []rune) {
	permHelp(chars, 0, len(chars)-1)
}

func permHelp(chars []rune, left, right int) {
	if left == right {
		fmt.Println(string(chars))
		return
	}
	for i := left; i <= right; i++ {
		chars[left], chars[i] = chars[i], chars[left]
		permHelp(chars, left+1, right)
		chars[left], chars[i] = chars[i], chars[left]
	}
}

// Produce all the combinations of characters with length n from a char list.
// Should produce `len(chars) choose n` ammount of combinations
func combine(chars []rune, n int) {
	combineHelp(chars, n, make([]rune, n), 0, 0)
}

/*
	chars is the pool of choosable characters
	comb is the actual combinations that get printed out.
	idx is the index of where the "cursor" is at. It does not influence what gets put in there.
	we always modify the idx -th position of comb character
	choosing is the character that is put into the idx -th position.

	The first rec call makes sure that the char after idx is choosen from the chars which are after idx.
	So if idx = 0 with choose = position of "a", chars = "abcd", n = 2 after that call, we are working with "a(b|c|d)"

	The second rec call makes sure that all the posible chars get enumerated at position idx.
	So if we are calling it with idx = 1, choosing = position of "b", n = 2, chars = "abcd" and assume that comb[0] == "a"
	then inside that next function call, possible combinations are "ac"

if n = 3 then "ac(d)"
*/
func combineHelp(chars []rune, n int, comb []rune, idx int, choosing int) {
	if n == idx {
		fmt.Println(string(comb))
		return
	}
	if choosing >= len(chars) {
		return
	}
	comb[idx] = chars[choosing]
	combineHelp(chars, n, comb, idx+1, choosing+1)
	combineHelp(chars, n, comb, idx, choosing+1)
}

// Cheeky solution, but probably not the best. combHelp makes a lot of rec calls anyways
func subsetWithComb(set []rune) {
	for length := 0; length <= len(set); length++ {
		combine(set, length)
	}
}

/*
at every element, we decide if we include it in the set or not.
This is a binary decision. For "ab", it could be:
00, 01, 10, 11
So we enumerate from 0 to 2^len(set) and instead of printing 0s or 1s,
we not print the set elelment or we print it.
*/
func subset(set []rune) {
	subsetHelp(set, make([]rune, len(set)), 0)
}

func subsetHelp(set []rune, subset []rune, idx int) {
	if idx >= len(set) {
		fmt.Println(string(subset))
		return
	}
	subsetHelp(set, subset, idx+1)
	subsetHelp(set, append(subset, set[idx]), idx+1)
}

/*
Enumerate from 0 to 2^len(set) in binary. This is the first for statement.
Every bit represents a rune from the set. There will be len(set) bits.
1 = include, 0 = not include in the subset. So for set = "abc", when
i = 5 = 0b101 we map those bits to "cba" and it'll produce "ac".
(it is easyer to think about it if the bit with the lowest value (written as the right-most bit) corresponds to the first element of the set (which is a slice, so it can have a first elemenet))

The inner loop is responsible for deciding which chars get put into the subset,
based on the value of i. It check the first, second, ..., len(set)-1 th positioned bit.
For the i = 5 example, it would be:

			 	j=0		j=1		j=2
				  a      b      c
				--------------------
			 i:	101 | 	101 | 	101
		ckdBit:	001 |  	010 | 	100
			 AND--- |AND--- |AND---
	 ckdBit& i:	001 |	000 | 	100
				  a             c
*/
func subsetIter(set []rune) {
	subsetCount := 1 << len(set)
	subset := make([]rune, len(set))
	for i := 0; i < subsetCount; i++ {
		clear(subset)
		for j := 0; j < len(set); j++ {
			checkedBit := 1 << j
			if i&checkedBit > 0 {
				subset = append(subset, set[j])
			}
		}
		fmt.Printf("{ %s }\n", string(subset))
	}
}
