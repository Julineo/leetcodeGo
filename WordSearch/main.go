package main

import (
	"fmt"
)

func main () {
	board := [][]byte{{'A','B','C','E'},
			{'S','F','C','S'},
			{'A','D','E','E'}}
	word := "ABCCED"//true
	word = "AFS"//false

	board = [][]byte{{'C','A','A'},
			{'A','A','A'},
			{'B','C','D'}}
	word = "AAB"//true

	board = [][]byte{{'A','B','C','E'},
			{'S','F','E','S'},
			{'A','D','E','E'}}

	word = "ABCESEEEFS"//true


	fmt.Println(exist(board, word))

}

func exist(board [][]byte, word string) bool {

	for m := 0; m < len(board); m++ {
		for n := 0; n < len(board[m]); n++ {
			if helper(board, m, n, word, 0) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, m int, n int, word string, d int) bool {
	//fmt.Println("m:", m, "n:", n, "d:", d)
	if d == len(word) {
		return true
	}
	if m < 0 || n < 0 || m >= len(board) || n >= len(board[m]) {
		return false
	}
	if word[d] != board[m][n] {
		return false
	}

	//fmt.Println(string(board[m][n]))
	//fmt.Println(word[d])
	board[m][n] ^= 255//to avoid geting to the same character. XOR to save memory
	fmt.Printf("%c/n", board)
	if      helper(board, m-1, n, word, d+1) ||
		helper(board, m+1, n, word, d+1) ||
		helper(board, m, n-1, word, d+1) ||
		helper(board, m, n+1, word, d+1) {
		return true
	}

	board[m][n] ^= 255

	return false
}
