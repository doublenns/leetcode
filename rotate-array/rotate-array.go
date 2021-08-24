func reverseSlice(sl []int) {
    for i, j := 0, len(sl) - 1; i<j; i,j = i+1, j-1{
        sl[i], sl[j] = sl[j], sl[i]
    }
}

func rotate(nums []int, k int)  {
    l := len(nums)
    if k > l {
        k = k % l
    }
    if l > 1 {
        reverseSlice(nums[l-k:])
        reverseSlice(nums[:l-k])
        reverseSlice(nums)
    }
}