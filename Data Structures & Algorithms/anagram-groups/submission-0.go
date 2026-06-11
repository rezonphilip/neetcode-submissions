func groupAnagrams(strs []string) [][]string {
	dic := make(map[string][]string)

	for _, str := range strs {
		key := []byte(str)
		sort.Slice(key, func(a, b int) bool {return key[a] < key[b]})
		dic[string(key)] = append(dic[string(key)], str)
	}
	result := [][]string{}
	for _, grp := range dic {
		result = append(result, grp)
	}
	return result
}
