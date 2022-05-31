func hasAllCodes(s string, k int) bool {
    
    slow,fast := 0,k
    expected := int(math.Pow(2, float64(k)))
    set := make(map[string]bool)
    
    for fast <= len(s) {
        dequeue := s[slow:fast]
        if _,ok := set[dequeue]; !ok {
            set[dequeue]=true
        }
        
        slow++
        fast++
    }
    
    return len(set) == expected
}
