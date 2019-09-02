package main

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var max, i, product, firstMinus, t int
	product = 1
	max = nums[0]
	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if i+1 < len(nums) {
				t = maxProduct(nums[i+1:])
				if t > max {
					max = t
				}
			}
			if max < 0 {
				max = 0
			}
			break
		}
		product *= nums[i]
		if product < 0 && firstMinus < 0 {
			t = product / firstMinus
		} else if product < 0 {
			firstMinus = product
			t = nums[i]
		} else {
			t = product
		}
		if t > max {
			max = t
		}
	}
	return max
}
