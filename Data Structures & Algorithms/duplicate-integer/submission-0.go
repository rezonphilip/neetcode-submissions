func hasDuplicate(nums []int) bool {
    lookup := make(map[int]bool, len(nums))
	for _, val := range nums {
		if lookup[val] {
			return true
		}
		lookup[val] = true
	}

	return false
}

