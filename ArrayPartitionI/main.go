import "sort"

func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    fmt.Println(nums)
    var s int
    for i := 0; i < len(nums); i = i + 2 {
        s = s + min(nums[i], nums[i+1])
    }
    return s
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
