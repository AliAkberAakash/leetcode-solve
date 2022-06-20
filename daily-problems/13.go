func romanToInt(s string) int {
    
    idx,sum := len(s)-1,0
    
    table := map[string]int{
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
    }
    
    for idx >= 0 {
        c := table[string(s[idx])]
        
        if idx-1 >= 0 {
          curr := string(s[idx-1])
            if table[curr] < c {
                sum += c - table[curr]
                idx--
            }else{
                sum += c
            }
        }else{
                sum += c
            }
        idx--
    }
    
    return sum
}
