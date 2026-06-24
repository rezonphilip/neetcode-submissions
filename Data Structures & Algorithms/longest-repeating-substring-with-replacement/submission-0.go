func characterReplacement(s string, k int) int {
    res := 0
    for i := 0; i < len(s); i++ {
        count := make(map[byte]int)
        maxf := 0
        for j := i; j < len(s); j++ {
            count[s[j]]++
            maxf = max(maxf, count[s[j]])
            if (j - i + 1) - maxf <= k {
                res = max(res, j - i + 1)
            }
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}