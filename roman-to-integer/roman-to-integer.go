func romanToInt(s string) int {
    var result int
    for i, v := range s {
        switch {
            case v == 'M':
                result += 1000
            case v == 'D':
                result += 500
            case v == 'C' && i+1 < len(s) && (s[i+1] == 'D' || s[i+1] == 'M'):
                result -= 100
            case v == 'C':
                result += 100
            case v == 'L':
                result += 50
            case v == 'X' && i+1 < len(s) && (s[i+1] == 'L' || s[i+1] == 'C'):
                result -= 10
            case v == 'X':
                result += 10
            case v == 'V':
                result += 5
            case v == 'I' && i+1 < len(s) && (s[i+1] == 'V' || s[i+1] == 'X'):
                result -= 1
            case v == 'I':
                result += 1
        }
    }
    return result
}