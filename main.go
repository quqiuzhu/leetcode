package main

import "fmt"

func main() {
	// fmt.Println(twoSum([]int{3, 2, 4}, 6)) //1
	// l1 := NewListNode([]int{9, 9})
	// l2 := NewListNode([]int{9})
	// l := addTwoNumbers(l1, l2) //2
	// l.Dump()
	// fmt.Println(lengthOfLongestSubstring("abba")) //3
	// fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4, 5})) //4
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
	// solutions := fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0) //18
	// for _, s := range solutions {
	// 	fmt.Println(s[0], s[1], s[2], s[3])
	// }
	// l := NewListNode([]int{1, 2}) //19
	// l = removeNthFromEnd(l, 2)
	// l.Dump()
	// fmt.Println(isValid("([]){}[]")) //20
	// l1 := NewListNode([]int{1, 2, 4}) //21
	// l := mergeTwoLists(l1, nil)
	// l.Dump()
	// results := generateParenthesis(4) //22
	// for _, r := range results {
	// 	fmt.Println(r)
	// }
	// l1 := NewListNode([]int{1, 4, 5}) //23
	// l2 := NewListNode([]int{1, 3, 4})
	// l3 := NewListNode([]int{2, 6})
	// l := mergeKLists([]*ListNode{l1, l2, l3, nil})
	// l.Dump()
	// l := swapPairs(NewListNode([]int{1, 2, 3, 4})) // 24
	// l.Dump()
	// l := reverseKGroup(NewListNode([]int{1, 2, 3, 4, 5, 6}), 2) //25
	// l.Dump()
	// nums := []int{1, 1, 2} //26
	// l := removeDuplicates(nums)
	// for i := 0; i < l; i++ {
	// 	fmt.Print(nums[i], "->")
	// }
	// fmt.Println("  l:", l)
	// nums := []int{3} //27
	// l := removeElement(nums, 3)
	// for i := 0; i < l; i++ {
	// 	fmt.Print(nums[i], "->")
	// }
	// fmt.Println("  l:", l)
	// fmt.Println(strStr("1aaaa", "a")) //28
	// fmt.Println(divide(2147483649, -1)) //29
	// rs := findSubstring("barfoothefoobarman", []string{"foo", "bar"}) //30
	// for _, r := range rs {
	// 	fmt.Println(r)
	// }
	// rs := []int{1, 2, 4, 3} //31
	// nextPermutation(rs)
	// for _, r := range rs {
	// 	fmt.Println(r)
	// }
	// fmt.Println(longestValidParentheses("((()))())")) //32
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) //33
	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 7)) //34
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0)) //35
	// sudoku := [][]byte{
	// 	[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }
	// fmt.Println(isValidSudoku(sudoku)) //36
	// solveSudoku(sudoku) //37
	// Matrix(sudoku).Dump()
	// fmt.Println(countAndSay(5)) //38
	// r := combinationSum([]int{8, 7, 4, 3}, 11) //39
	// Matrix(r).Dump()
	// r := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8) //40
	// Matrix(r).Dump()
	// fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12})) //41
	// fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) //42
	// fmt.Println(multiply("123", "0")) //43
	// fmt.Println(dpIsMatch("adceb", "*a*b")) //44
	// fmt.Println(jump([]int{2, 3, 1, 1, 4})) //45
	// r := permute([]int{1, 2, 3}) //46
	// Matrix(r).Dump()
	// r := permuteUnique([]int{0, 1, 0, 0, 9}) //47
	// Matrix(r).Dump()
	// m := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// 	[]int{7, 8, 9},
	// }
	// rotate(m) //48
	// Matrix(m).Dump()
	r := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}) //49
	fmt.Println(r)
}
