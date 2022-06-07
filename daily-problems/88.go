func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    if n==0 {
        return
    }
    
    currIndex := len(nums1)-1
    
    index1 := m-1
    index2 := n-1
    
    for index1 >= 0 && index2 >= 0 {
        if nums1[index1] > nums2[index2] {
            nums1[currIndex] = nums1[index1]
            index1--
        }else{
            nums1[currIndex] = nums2[index2]
            index2--
        }
        currIndex--
    }
    
    for index1 >= 0 {
        nums1[currIndex] = nums1[index1]
        index1--
        currIndex--
    }
    
    for index2 >= 0 {
        nums1[currIndex] = nums2[index2]
        index2--
        currIndex--
    }
    
}
