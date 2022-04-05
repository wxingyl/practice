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

func countOneBits(n int) int {
	cnt := 0
	for n > 0 {
		n &= n - 1
		cnt++
	}
	return cnt
}

func god(x, y int) int {
	if x == y {
		return x
	} else if y > x {
		x, y = y, x
	}
	for {
		x = x % y
		if x == 0 {
			return y
		}
		x, y = y, x
	}
}
