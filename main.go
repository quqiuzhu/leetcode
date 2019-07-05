package main

import "fmt"

func main() {
	// fmt.Println(longestPalindrome("bababd")) //5
	// fmt.Println(convert("PAYPALISHIRING", 4)) //6
	// fmt.Println(reverse(-123)) //7
	// fmt.Println(myAtoi("2147483648")) //8
	// fmt.Println(isPalindrome(121)) //9
	// fmt.Println(isMatch("accbbbaababbac", ".*...*c*.")) //10
	// fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) //11
	// fmt.Println(intToRoman(2994)) //12
	// fmt.Println(romanToInt("MMCMXCIV")) //13
	// fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) //14
	// solutions := threeSum([]int{3, 0, -2, -1, 1, 2}) //15
	// for _, s := range solutions {
	// 	fmt.Println(s[0], s[1], s[2])
	// }
	// fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) //16
	// letters := letterCombinations("23") //17
	// for _, s := range letters {
	// 	fmt.Println(s)
	// }
	solutions := fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0) //18
	for _, s := range solutions {
		fmt.Println(s[0], s[1], s[2], s[3])
	}
}
