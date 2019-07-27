package main

/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */
func getPermutation(n int, k int) string {
	nums := make([]byte, n)
	total := 1
	for i := 0; i < n; i++ {
		nums[i] = byte(i + 1)
		total *= (i + 1)
	}

	var j, m int
	for i := n; i > 0; i-- {
		total = total / i
		j = (k - 1) / total
		// fmt.Println(total, j, k)
		nums[n-i], nums[n-i+j] = nums[n-i+j], nums[n-i]
		for m = n - i + j; m > n-i+1; m-- {
			nums[m], nums[m-1] = nums[m-1], nums[m]
		}
		k = k - j*total
	}

	for i := 0; i < n; i++ {
		nums[i] = '0' + nums[i]
	}
	return string(nums)
}
