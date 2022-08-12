func subsets(nums []int) [][]int {
    var set [][]int
    
    set = append(set,[]int{})

    for _,num := range nums {
        length := len(set)
        
        for i:=0; i<length; i++ {
            item := make([]int, len(set[i]))
            copy(item, set[i])
            
            item = append(item, num)
            set = append(set,item)
        }
    }
    
    return set
}
