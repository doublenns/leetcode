func reverseString(s []byte)  {
    // Recursive solution
    var swap func(int, int)
    swap = func(i, j int) {
        if i<j {
            s[i], s[j] = s[j], s[i]
            i++
            j--
            swap(i, j)
        }
    }
    i, j := 0, len(s)-1
    swap(i, j)
}