func hammingDistance(x int, y int) int {
    fmt.Printf("x: %08b\n", x)
    fmt.Printf("y: %08b\n", y)
    fmt.Printf("x^y: %08b\n", x^y)
    z := x^y
    s := 1
    count := 0
    for i := 0; i < 32; i++ {
        fmt.Printf("s: %08b\n", s)
        fmt.Printf("%08b\n", z & s)
        if z & s != 0 {
            count++
        }
        s = s << 1
    }
    return count
}
