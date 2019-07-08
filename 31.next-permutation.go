package main

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */
func nextPermutation(nums []int) {
	ln := len(nums)
	if ln == 0 {
		return
	}

	var i, j, k, v int
	found := false
	for i = 0; i < ln-1; i++ {
		if nums[i] < nums[i+1] {
			v = nums[i]
			j = i
			found = true
		}
		if nums[i+1] > v {
			k = i + 1
		}
	}

	if found {
		nums[j], nums[k] = nums[k], nums[j]
		i = j + 1
		j = ln - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	} else {
		for i := 0; i < ln/2; i++ {
			nums[i], nums[ln-1-i] = nums[ln-1-i], nums[i]
		}
	}
}
