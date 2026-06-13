func isValid(s string) bool {
    lookup := map[rune]int {
		'(' : 1,
		'{' : 2,
		'[' : 3,
		')' : -1,
		'}' : -2,
		']' : -3,
	}
	stack := []int{}
	head := 0
	for _, letter := range s {
		if lookup[letter] > 0 {
			stack = append(stack, lookup[letter])
			head++
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if stack[head-1] != -1 * lookup[letter] {
			return false
		}
		stack = stack[:len(stack)-1]
		head--
	}

	return head == 0
}
