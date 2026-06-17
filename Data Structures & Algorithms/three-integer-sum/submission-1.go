func threeSum(nums []int) [][]int {
    result := [][]int{}
    n := len(nums)
    sort.Slice(nums, func(a , b int) bool {return nums[a] < nums[b]})
    for idx, val := range nums {
        if idx > 0 && nums[idx-1] == nums[idx] {
            continue
        }
        if val > 0 {
            return result
        }
        for i, j := idx + 1, n - 1; i < j ;  {
            comp := nums[i] + nums[j]
            if comp == -1 * val {
                result = append(result, []int{val, nums[i], nums[j]})
                i = i + 1
                j = j - 1

                for i < j && nums[i] == nums[i-1] {
                    i++
                }

                for i < j && nums[j] == nums[j+1] {
                    j--
                }
            }else if comp > -1 * val {
                j = j -1 
            } else {
                i = i + 1
            }
        }
    }

    return result
}