func judgeCircle(moves string) bool {
    runes := []rune(moves)   
    var h, v int
    for _, r := range(runes) {
        switch r {
            case 'U':
                v++
            case 'D':
                v--
            case 'L':
                h--
            case 'R':
                h++
        }
    }
    return h + v == 0
}
