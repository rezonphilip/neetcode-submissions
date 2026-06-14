func search(nums []int, target int) int {
	n := len(nums)
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if target == nums[mid] {
			return mid
		}else if target > nums[mid] {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	return -1
}
