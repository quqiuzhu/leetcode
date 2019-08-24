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
	// r := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}) //49
	// fmt.Println(r)
	// fmt.Println(myPow(2.10000, 3)) //50
	// rs := solveNQueens(4) //51
	// for i := 0; i < len(rs); i++ {
	// 	fmt.Println()
	// 	for j := 0; j < len(rs[i]); j++ {
	// 		fmt.Println(rs[i][j])
	// 	}
	// }
	// fmt.Println(totalNQueens(4)) //52
	// fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //53
	// fmt.Println(spiralOrder(m)) //54
	// fmt.Println(canJump([]int{3, 2, 1, 0, 4})) //55
	// intervals := [][]int{
	// 	[]int{1, 2},
	// 	[]int{3, 5},
	// 	[]int{6, 7},
	// 	[]int{8, 10},
	// 	[]int{12, 16},
	// }
	// r := merge(intervals) //56
	// Matrix(r).Dump()
	// r := insert(intervals, []int{4, 8}) //57
	// Matrix(r).Dump()
	// fmt.Println(lengthOfLastWord("Hello World")) //58
	// r := generateMatrix(5) //59
	// Matrix(r).Dump()
	// fmt.Println(getPermutation(4, 9)) //60
	// l := rotateRight(NewListNode([]int{1, 2, 3, 4, 5}), 6) //61
	// l.Dump()
	// fmt.Println(uniquePaths(7, 3)) //62
	// m := [][]int{
	// 	[]int{0, 0, 0},
	// 	[]int{0, 1, 0},
	// 	[]int{0, 0, 0},
	// }
	// fmt.Println(uniquePathsWithObstacles(m)) //63
	// Matrix(m).Dump()
	// m := [][]int{
	// 	[]int{1, 3, 1},
	// 	[]int{1, 5, 1},
	// 	[]int{4, 2, 1},
	// }
	// fmt.Println(minPathSum(m)) //64
	// Matrix(m).Dump()
	// fmt.Println(isNumber(".1")) //65
	// fmt.Println(plusOne([]int{9, 9, 9})) //66
	// fmt.Println(addBinary("1010", "1011")) //67
	// strs := []string{"This", "is", "an", "example", "of", "text", "justification."}
	// ss := fullJustify(strs, 16) //68
	// for i := 0; i < len(ss); i++ {
	// 	fmt.Println(ss[i])
	// }
	// fmt.Println(mySqrt(8)) //69
	// fmt.Println(climbStairs(3)) //70
	// fmt.Println(simplifyPath("/home/")) //71
	// fmt.Println(minDistance("intention", "execution")) //72
	// m := [][]int{
	// 	[]int{0, 1, 2, 0},
	// 	[]int{3, 4, 5, 2},
	// 	[]int{1, 3, 1, 5},
	// }
	// setZeroes(m) //73
	// Matrix(m).Dump()
	// m := [][]int{
	// 	[]int{1, 1},
	// }
	// fmt.Println(searchMatrix(m, 12)) //74
	// colors := []int{2, 0, 2, 1, 1, 0}
	// sortColors(colors) //75
	// fmt.Println(colors)
	// fmt.Println(minWindow("aa", "aa")) //76
	// Matrix(combine(4, 2)).Dump() //77
	// Matrix(subsets([]int{1, 2, 3})).Dump() //78
	// board := [][]byte{
	// 	[]byte{'A', 'B', 'C', 'E'},
	// 	[]byte{'S', 'F', 'C', 'S'},
	// 	[]byte{'A', 'D', 'E', 'E'},
	// }
	// fmt.Println(exist(board, "SEE")) //79
	// l := []int{1, 1, 1, 2, 2, 3}
	// fmt.Println(removeDuplicatesII(l)) //80
	// fmt.Println(l)
	// fmt.Println(searchII([]int{2, 5, 6, 0, 0, 1, 2}, 3)) //81
	// l := NewListNode([]int{1, 2, 3, 3, 4, 4, 5})
	// r := deleteDuplicatesII(l) //82
	// r.Dump()
	// l := NewListNode([]int{1, 2, 3, 3, 4, 4, 5, 5})
	// r := deleteDuplicates(l) //83
	// r.Dump()
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) //84
	// m := [][]byte{
	// 	[]byte{'1', '0', '1', '0', '0'},
	// 	[]byte{'1', '0', '1', '1', '1'},
	// 	[]byte{'1', '1', '1', '1', '1'},
	// 	[]byte{'1', '0', '0', '1', '0'},
	// }
	// fmt.Println(maximalRectangle(m)) //85
	// l := NewListNode([]int{1, 4, 3, 2, 5, 2})
	// partition(l, 3).Dump() //86
	// fmt.Println(isScramble("great", "rgeat")) //87
	// l1 := []int{1, 2, 4, 5, 6, 0}
	// l2 := []int{3}
	// mergeII(l1, 5, l2, 1) //88
	// fmt.Println(l1)
	// fmt.Println(grayCode(3)) //89
	// Matrix(subsetsWithDup([]int{1, 1, 2, 2})).Dump() //90
	// fmt.Println(numDecodings("301")) //91
	// l := NewListNode([]int{1, 2, 3, 4, 5})
	// reverseBetween(l, 1, 2).Dump() //92
	// fmt.Println(restoreIpAddresses("25525511135")) //93
	// inorderTraversal(nil)     //94
	// trees := generateTrees(3) //95
	// for _, tree := range trees {
	// 	tree.Dump()
	// }
	// fmt.Println(numTrees(3)) //96
	// fmt.Println(isInterleave("aa", "ab", "abaa")) //97
	// fmt.Println(isValidBST(NewTreeNode([]int{10, 5, 15, -1, -1, 6, 20}))) //98
	// tree := NewTreeNode([]int{1, 3, -1, -1, 2})
	// recoverTree(tree) //99
	// tree.Dump()
	// t1 := NewTreeNode([]int{1, 2})
	// t2 := NewTreeNode([]int{1, -1, 2})
	// fmt.Println(isSameTree(t1, t2)) //100
	fmt.Println(isSymmetric(NewTreeNode([]int{1, 2, 2, 3, 4, 4, 3}))) //101
}
