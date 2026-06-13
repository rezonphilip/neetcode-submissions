func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	left[0] = 1
	right := make([]int, n)
	right[n-1] = 1
	l, r := 1, n-2
	for l < n && r >= 0 {
		left[l] = left[l - 1] * nums[l - 1]
		right[r] = right[r + 1] * nums[r + 1]
		l = l + 1
		r = r - 1 
	}
	result := make([]int, n)
	for i := range n {
		result[i] = left[i] * right[i]
	}
	return result
}
