func intersection(nums1 []int, nums2 []int) []int {
    seen := make(map[int]struct{})
    exists := struct{}{}
    
    for _, v := range nums1 {
        seen[v] = exists
    }
    
    var result []int
    for _, v := range nums2 {
        if _, ok := seen[v]; ok {
            result = append(result, v)
            delete(seen, v)
        }
    }
    
    return result
}