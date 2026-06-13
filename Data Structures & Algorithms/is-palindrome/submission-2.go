func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	length := len(s)
	left, right := 0, length-1

	var valid func(c rune)bool
	valid = func(c rune) bool {
		if (c >= '0'&& c <= '9') ||
		(c >= 'a' && c <= 'z') {
			return true
		}
		return false
	}

	for left <= right {
		if !valid(rune(s[left])) {
			left++
			continue
		}

		if !valid(rune(s[right])){
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true

}
