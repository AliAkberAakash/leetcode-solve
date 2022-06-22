func rob(nums []int) int {
    
    house1,house2 := 0,0
    
    for _,num := range nums {
        tmp := max(house1+num, house2)
        house1 = house2
        house2 = tmp
    }
    
    return house2
}

func max(x,y int) int {
    if x >= y {
        return x
    }
    return y
}
