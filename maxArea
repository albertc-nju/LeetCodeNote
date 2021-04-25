func maxArea(height []int) int {
    i := 0
    j := len(height) - 1 
    max := 0
    tmp := 0
    for i < j {
        if height[i] < height[j] {
            tmp = (j - i) * height[i]
            i++
        }else {
            tmp = (j - i) * height[j]
            j--
        }
        if max < tmp {
            max = tmp
        }
    }
    return max
}
