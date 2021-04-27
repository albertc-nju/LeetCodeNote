func nextPermutation(nums []int)  {
    flag := false
    n := len(nums)
    tmpValue := math.MaxInt32
    tmpIndex := 0
    for i := n - 1; i >= 0; i-- {
        for j := n - 1; i < j; j-- {
            if nums[i] < nums[j] {//发现顺序对
                if tmpValue > nums[j] {
                    tmpValue = nums[j] //记录最小值
                    tmpIndex = j
                }
                flag = true
            }
        }
        if flag {
            nums[i], nums[tmpIndex] = nums[tmpIndex], nums[i]
            sort.Ints(nums[(i + 1):])
            break
        }               
    }
    if !flag {
        i := 0
        j := len(nums) - 1
        for i < j {
            nums[i], nums[j] = nums[j], nums[i]
            i++
            j--
        }
    }
}
