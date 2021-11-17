// Bottom-up iterative solution
func fib(n int) int {
    prev := 0
    curr := 1
    
    if n == 0 {
        return prev
    } else if n == 1 {
        return curr
    }
    
    for i := 2; i <= n ; i++ {
        temp := curr + prev
        prev = curr
        curr = temp
    }
    return curr
}