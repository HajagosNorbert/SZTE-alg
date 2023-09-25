package main

func main() {
	// permutate([]rune("abcd"))
	combine([]rune("abcd"), 2)
}

// Produce every possible permutation of the provided characters
func permutate(chars []rune) {
	permHelp(chars, 0, len(chars)-1)
}

func permHelp(chars []rune, left, right int) {
	if left == right {
		println(string(chars))
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
		println(string(comb))
		return
	}
	if choosing >= len(chars) {
		return
	}
	comb[idx] = chars[choosing]
	combineHelp(chars, n, comb, idx+1, choosing +1)
	combineHelp(chars, n, comb, idx, choosing+1)
}
