func twoSum(nums []int, target int) []int {
	index := make([][2]int, len(nums))
	for i, val := range nums {
		index[i] = [2]int{val, i}
	}

    sort.Slice(index, func(a, b int) bool {return index[a][0] < index[b][0]})
	i, j := 0, len(nums) - 1
	for i < j {
		candidate := index[i][0] + index[j][0]
		if candidate == target {
			if index[i][1] < index[j][1] {
				return []int{index[i][1], index[j][1]}
			}else {
				return []int{index[j][1], index[i][1]}
			}
			
		}else if candidate > target {
			j = j - 1
		}else {
			i = i + 1
		}
	}
	return []int{}
}
