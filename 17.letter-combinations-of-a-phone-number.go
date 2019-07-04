package main

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	x := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	ld := len(digits)
	length := 1
	for i := 0; i < ld; i++ {
		s := x[digits[i]]
		length *= len(s)
	}

	letters := make([][]byte, length)
	for i := 0; i < length; i++ {
		letters[i] = make([]byte, ld)
	}

	each := length
	for i := 0; i < ld; i++ {
		s := x[digits[i]]
		ls := len(s)
		each /= ls
		for j := 0; j < length; j++ {
			letters[j][i] = s[(j/each)%ls]
		}
	}
	results := make([]string, length)
	for i := 0; i < length; i++ {
		results[i] = string(letters[i])
	}
	return results
}
