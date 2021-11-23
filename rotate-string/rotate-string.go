func rotateString(s string, goal string) bool {
    if len(goal) != len(s) {
        return false
    }
    check := s + s
    return strings.Contains(check, goal)
}