// Brute force solution
func isPalindrome(head *ListNode) bool {
    // 1st load the Linked List into a slice
    var arr []int
    for curr := head; curr != nil; curr = curr.Next {
        arr = append(arr, curr.Val)
    }
    
    // Check if the slice is a palindrome
    for i, j := 0, len(arr) -1; i < j; i,j = i+1, j-1 {
        if arr[i] != arr[j] {
            return false
        }
    }
    return true
}