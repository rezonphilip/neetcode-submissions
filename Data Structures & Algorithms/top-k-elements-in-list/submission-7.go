func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	max := 0
	for _, val := range nums {
		freq[val]++
		if freq[val] > max {
			max = freq[val]
		}
	}
	result := make([][]int, max+1)
	
	for key, val := range freq {
		result[val] = append(result[val], key)
	}
	topK := []int{}
	for i := max; i >= 0 && k > 0; {
		if len(result[i]) > k {
			topK = append(topK, result[i][:k]...)
			break
		}else {
			topK = append(topK, result[i]...)
			k = k - len(result[i])
		}
		i--
	}
	return topK

}
