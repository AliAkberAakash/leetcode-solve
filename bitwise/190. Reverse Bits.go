func reverseBits(num uint32) uint32 {
    
    var operator uint32 = 1
    var ans uint32 = 0
    
    for i:=0; i<32; i++ {
        var tmp uint32 = num & operator
        
        ans = ans << operator
        ans = ans | tmp
        
        num = num >> operator
    }
    
    return ans
}
