func isValid(s string) bool {
    stack := []rune{}

    for _, ch := range s {
        switch ch {
        case '(':
            stack = append(stack, ')')
        case '[':
            stack = append(stack, ']')
        case '{':
            stack = append(stack, '}')
        default:
            if len(stack) == 0 || stack[len(stack)-1] != ch {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}