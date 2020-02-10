package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type path struct {
	m, n int
	output int 
}
func Test_uniquePaths(t *testing.T) {
	items := []*path{
		&path{
			m:      3,
			n:      2,
			output: 3,
		},
		&path{
			m:      7,
			n:      3,
			output: 28,
		},
	}

	for _, item := range items {
		if uniquePaths(item.m, item.n) != item.output {
			t.Error("assert failed", item.output)
		}
	}
}

type TestMinPathSum struct {
	data [][]int
	result int
}

func Test_MinPathSum(t *testing.T) {
	items := []*TestMinPathSum{
		{
			data:[][]int{{1,3,1},{1,5,1},{4,2,1}},
			result: 7,
		},
		{
			data:[][]int{{1,2,5},{3,2,1}},
			result: 6,
		},
	}

	art := assert.New(t)
	for _, item := range items {
		res := minPathSum(item.data)
		art.Equal(item.result, res, "assert failed")
	}

	for _, item := range items {
		res := minPathSum2(item.data)
		art.Equal(item.result, res, "assert failed")
	}
}

type profit struct {
	input []int
	output int
}
func Test_MaxProfit(t *testing.T) {
	items := []*profit{
		{
			input: []int{7,1,5,3,6,4},
			output:5,
		},
		{
			input: []int{7,6,4,3,1},
			output: 0,
		},
		{
			input: []int{},
			output: 0,
		},
	}

	ast := assert.New(t)
	for _, item := range items {
		ast.Equal(item.output, maxProfit(item.input), "assert failed")
	}
}


type robTest struct {
	input []int
	output int
}
func Test_rob(t *testing.T) {
	items := []*robTest {
		{
			input: []int{1,2,3,1},
			output: 4,
		},
		{
			input: []int{2,7,9,3,1},
			output: 12,
		},
		{
			input: []int{0},
			output: 0,
		},
		{
			input: []int{2,1,1,2},
			output: 4,
		},
	}
	for _, item := range items {
		assert.Equal(t, item.output, rob(item.input), "assert failed")
	}
}

type TestMinCost struct {
	input []int
	output int
}

func Test_MinCost(t *testing.T) {
	items := []*TestMinCost {
		{
			input: []int{10, 15, 20},
			output: 15,
		},
		{
			input: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			output: 6,
		},
		{
			input: []int{0, 0, 0, 0},
			output: 0,
		},
	}

	for _, item := range items {
		assert.Equal(t, item.output, minCostClimbingStairs(item.input), "assert failed")
	}
}