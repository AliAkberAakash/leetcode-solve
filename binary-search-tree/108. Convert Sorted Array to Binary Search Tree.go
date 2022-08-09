func sortedArrayToBST(nums []int) *TreeNode {
    
    if len(nums)==1 {
        return &TreeNode {
            Val: nums[0],
        }
    }
    
    var root *TreeNode
    root = findNode(nums, 0, len(nums)-1)
    
    return root
}

func findNode(nums []int, start int, end int) *TreeNode{
    
    if start > end {
        return nil
    }
    if start == end {
        return &TreeNode {
            Val: nums[start],
        }
    }
    
    var root *TreeNode
    mid := (start+end)/2
    
    root = &TreeNode {
        Val: nums[mid],
    }
    
    root.Left = findNode(nums, start, mid-1)
    root.Right = findNode(nums, mid+1, end)
    
    return root
}
