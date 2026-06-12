func topKFrequent(nums []int, k int) []int {
    lookup := make(map[int]int)
    for _,v := range nums {
        lookup[v]++
    }
    result := []int{}
    for k, _ := range lookup {
        result = append(result, k)
    }
    sort.Slice(result, func(a,b int) bool {return lookup[result[a]] > lookup[result[b]]})
    return result[:k]
}
