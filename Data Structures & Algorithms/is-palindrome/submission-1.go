func isPalindrome(s string) bool {
	length := len(s)
	left, right := 0, length-1

	var valid func(c rune)bool
	valid = func(c rune) bool {
		if (c >= '0'&& c <= '9') ||
		(c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') {
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

		if s[left] != s[right] && (rune(s[left])  >= '0' && rune(s[left]) <= '9' || rune(s[right])  >= '0' && rune(s[right]) <= '9') {
			return false
		}
		if s[left] != s[right] && s[left] != s[right]-32 && s[left] != s[right] + 32 {
			return false
		}

		left++
		right--
	}
	return true

}
