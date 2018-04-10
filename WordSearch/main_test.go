package main

import "testing"

func TestExist(t *testing.T) {
	var tests = []struct {
		board[][]byte
		word string
		want bool
	}{
		{
			[][]byte{{'A','B','C','E'},
				{'S','F','C','S'},
				{'A','D','E','E'}},
			"ABCCED",
			true,
		},
		{
			[][]byte{{'A','B','C','E'},
				{'S','F','C','S'},
				{'A','D','E','E'}},
			"AFS",
			false,
		},
		{
			[][]byte{{'C','A','A'},
				{'A','A','A'},
				{'B','C','D'}},
			"AAB",
			true,
		},
		{
			[][]byte{{'A','B','C','E'},
				{'S','F','E','S'},
				{'A','D','E','E'}},
			"ABCESEEEFS",
			true,
		},
	}
	for _, test := range tests {
		//copy slice to show it later in test case output
		origin := make([][]byte, len(test.board))
		for i := range test.board {
			origin[i] = make([]byte, len(test.board[i]))
			copy(origin[i], test.board[i])
		}

		//run test
		if got := exist(test.board, test.word); got != test.want {
			t.Errorf("exist(%c,%q) = %v", origin, test.word, got)
		}
	}
}
