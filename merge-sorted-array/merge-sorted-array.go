func merge(nums1 []int, m int, nums2 []int, n int)  {
    res := []int{}
    for p1, p2 := 0,0 ; p1 < len(nums1) || p2 < len(nums2); p1, p2 = p1+1, p2+1 {
        if p1 >= len(nums1) {
            res = append(res, nums2[p2])
            p1--
        } else if p2 >= len(nums2) {
            res = append(res, nums1[p1])
            p2--
        } else if nums1[p1] <= nums2[p2] && p1 < m {
            res = append(res, nums1[p1])
            p2--
        } else {
            res = append(res, nums2[p2])
            p1--
        }
    }
    copy(nums1, res)
    
}