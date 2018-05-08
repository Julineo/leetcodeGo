func selfDividingNumbers(left int, right int) []int {
    var r []int
    var n int
    for i := left; i <= right; i++ {
        aa := strconv.Itoa(i)
        nn, _ := strconv.Atoi(aa)
        len := len(aa)
        for _, a := range(aa) {
            n, _ = strconv.Atoi(string(a))
            if n > 0 {
            if nn % n == 0 {
                len--
            }
            }
        }
        if len == 0 {
            r = append(r, nn)
        }
    }
    return r
}
