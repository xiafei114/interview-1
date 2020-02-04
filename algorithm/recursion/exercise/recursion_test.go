package main

import "testing"

func Test_Reverse_Recursion(t *testing.T) {
	node := &Node{
		Next:  &Node{
			Next:  &Node{
				Next:  &Node{
					Next:  nil,
					Value: 4,
				},
				Value: 3,
			},
			Value: 2,
		},
		Value: 1,
	}

	node = ReverseRecursion(node)
	for {
		if node == nil {
			break
		}

		t.Log(node.Value)
		node = node.Next
	}
}

type TicTacToe struct {
	board []string
	result bool
}

func Test_ValidTicTacToe(t *testing.T) {
	boards := []*TicTacToe{
		&TicTacToe{
			board: []string{"   ", "   ", "   "},
			result: true,
		},
		&TicTacToe{
			board: []string{"XOX", " X ", "   "},
			result: false,
		},
		&TicTacToe{
			board: []string{"XXX", "   ", "OOO"},
			result: false,
		},
		&TicTacToe{
			board: []string{"XOX", "O O", "XOX"},
			result: true,
		},
		&TicTacToe{
			board: []string{"XXX", "XOO", "OO "},
			result: false,
		},
	}

	for _, board := range boards {
		res := validTicTacToe(board.board)
		if res != board.result {
			t.Error("failed", board)
		}
	}
}


type JumpStruct struct {
	data int
	result int
}

func Test_Jump(t *testing.T) {
	items := []*JumpStruct{
		&JumpStruct{
			data: 1,
			result: 1,
		},
		&JumpStruct{
			data:   2,
			result: 2,
		},
		&JumpStruct{
			data:   3,
			result: 3,
		},
		&JumpStruct{
			data:   4,
			result: 5,
		},
	}

	for _, item := range items {
		res := Jump2(item.data)
		if res != item.result {
			t.Error("assert failed", res)
		}
	}
}
