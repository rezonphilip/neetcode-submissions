func lengthOfLongestSubstring(s string) int {
    max := 0
    unique := ""
    var currentLetter = ""
    for _,letter := range s {
        currentLetter = string(letter)
        if !strings.Contains(unique, currentLetter) {
            unique  = unique + currentLetter
            subSLen := len(unique)
            if max < subSLen {
                max = subSLen
            }
        } else {
            
            unique = unique[strings.Index(unique, currentLetter)+1:] + currentLetter
        }
    }

    return max
}
