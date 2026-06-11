func isAnagram(s string, t string) bool {
    key1 := []byte(s)
    sort.Slice(key1, func(a, b int) bool {return key1[a] < key1[b]})

    key2 := []byte(t)
    sort.Slice(key2, func(a, b int) bool {return key2[a] < key2[b]})
    return string(key1) == string(key2)
}