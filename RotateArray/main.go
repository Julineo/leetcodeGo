package main

func rotate(nums []int, k int)  {
	if len(nums) == 0 {
		return
	}
	n := len(nums)
	if n < k {
		k = k % n
	}
	reverse(nums, 0, n - 1)
	reverse(nums, k, n - 1)
	reverse(nums, 0, k - 1)
}

func reverse(nums []int, i, j int) {
    for i < j{
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}
