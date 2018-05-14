func flipAndInvertImage(A [][]int) [][]int {
    for k, _ := range(A) {
        for i, j := 0, len(A) - 1; i <= j; i, j = i + 1, j - 1 {
            A[k][i], A[k][j] = invert(A[k][j]), invert(A[k][i])
        }
    }
    return A
}

func invert(i int) int {
    if i == 0 {
        return 1
    }
    return 0
}
