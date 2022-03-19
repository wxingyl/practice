package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
