func maxArea(heights []int) int {
	max := 0
	l , r := 0, len(heights)-1

	min := func(a , b int) int {
		if a < b {
			return a
		}
		return b
	}
	for l <= r {
		area := min(heights[l], heights[r]) * (r-l)
		if area > max {
			max = area
		}
		if min(heights[l], heights[r]) == heights[l] {
			l++
		}else {
			r--
		}
	}

	return max
}
