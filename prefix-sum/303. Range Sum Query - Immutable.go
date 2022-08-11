type NumArray struct {
    Nums []int
    PrefixSum []int
}


func Constructor(nums []int) NumArray {
    
    var ps []int
    ps = append(ps, nums[0])
    
    for i:=1; i<len(nums); i++ {
      ps = append(ps, ps[i-1]+nums[i])
    }
    
    fmt.Println(ps)
    
    return NumArray{
        Nums:nums,
        PrefixSum:ps,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    
    l := 0
    
    if left-1 >=0 {
      l = this.PrefixSum[left-1]
    }
    
    return this.PrefixSum[right] - l
}
