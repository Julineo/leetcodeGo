package main

import "testing"

func TestOne(t *testing.T) {
	board := [][]byte{{'A','B','C','E'},
			{'S','F','C','S'},
			{'A','D','E','E'}}
	word := "ABCCED"//true

	if !exist(board, word) {
		t.Error(`testOne fails`, "ABCCED")
	}

	word = "AFS"//false

	if exist(board, word) {
		t.Error(`testOne fails`, "AFS")
	}
}

func TestTwo(t *testing.T) {
	board := [][]byte{{'C','A','A'},
			{'A','A','A'},
			{'B','C','D'}}
	word := "AAB"//true

	if !exist(board, word) {
		t.Error(`testTwo fails`, "AAB")
	}

	board = [][]byte{{'A','B','C','E'},
			{'S','F','E','S'},
			{'A','D','E','E'}}

	word = "ABCESEEEFS"//true

	if !exist(board, word) {
		t.Error(`TestTwo fails`, "ABCESEEEFS")
	}
}
