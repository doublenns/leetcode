// Bottom-up iterative solution
func fib(n int) int {
    prev := 0
    curr := 1
    
    if n <= 1 {
        return n
    }
    
    for i := 2; i <= n ; i++ {
        temp := curr + prev
        prev = curr
        curr = temp
    }
    return curr
}