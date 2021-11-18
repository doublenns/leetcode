func containsDuplicate(nums []int) bool {
    seen:= make( map[int]struct{} )
    exists := struct{}{}
    
    for _,v := range nums{
        if _, ok := seen[v]; ok {
            return true
        }
        seen[v] = exists
    }
    return false
}
