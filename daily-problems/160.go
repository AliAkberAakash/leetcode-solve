func getIntersectionNode(headA, headB *ListNode) *ListNode {
    set := make(map[*ListNode]bool) 
    
    temp := headA
    
    for temp != nil {
        set[temp] = true
        temp = temp.Next
    }
    
    temp = headB
    
    for temp != nil {
        if _,ok := set[temp]; ok {
            return temp
        }
        
        temp = temp.Next
    }
    
    return nil
}
