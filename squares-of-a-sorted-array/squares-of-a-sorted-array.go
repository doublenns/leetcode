func mergeSortedArrays(nums1 []int, m int, nums2 []int, n int) []int  {
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
    return res
}

func prependInt(nums []int, y int) []int {
    nums = append(nums, 0)
    copy(nums[1:], nums)
    nums[0] = y
    return nums
}

func sortedSquares(nums []int) []int {
    var negative, positive []int
    for _, v := range nums {
        if v < 0 {
            negative = prependInt(negative, v*v)
        } else {
            positive = append(positive, v*v)
        }
    }
    return mergeSortedArrays(negative, len(negative), positive, len(positive))
}