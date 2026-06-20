func longestConsecutive(nums []int) int {
    set := make(map[int]bool)

    for _, num := range nums {
        set[num] = true
    }

    longest := 0

    for num := range set {
        // Only start at beginning of sequence
        if set[num-1] {
            continue
        }

        length := 1
        current := num

        for set[current+1] {
            current++
            length++
        }

        if length > longest {
            longest = length
        }
    }

    return longest
}