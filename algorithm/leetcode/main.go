package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(loudAndRich([][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}))
}

// https://leetcode-cn.com/problems/accounts-merge/
func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}

	union := func(x, y int) {
		p[find(x)] = find(y)
	}

	ap := map[string]int{}
	for i, v := range accounts {
		for _, email := range v[1:] {
			pindex, ok := ap[email]
			if ok {
				union(i, pindex)
			} else {
				ap[email] = i
			}
		}
	}

	for i, v := range accounts {
		pi := find(i)
		if pi != i {
			accounts[pi] = append(accounts[pi], v[1:]...)
			accounts[pi] = nil
		}
	}

	var ret [][]string
	for _, v := range accounts {
		if v != nil {
			sort.Strings(v[1:])
			i, j := 2, 1
			r := []string{v[0]}
			for ; i < len(v); i++ {
				if v[i] == v[j] {
					continue
				} else {
					r = append(r, v[j])
				}
				j = i
			}
			r = append(r, v[j])
			ret = append(ret, r)
		}
	}
	return ret
}

// https://leetcode-cn.com/problems/max-area-of-island/
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var bfs func(i, j int) int
	bfs = func(i, j int) int {
		grid[i][j] = 0
		v := 1
		if i > 0 && grid[i-1][j] == 1 {
			v += bfs(i-1, j)
		}
		if i < m-1 && grid[i+1][j] == 1 {
			v += bfs(i+1, j)
		}
		if j > 0 && grid[i][j-1] == 1 {
			v += bfs(i, j-1)
		}
		if j < n-1 && grid[i][j+1] == 1 {
			v += bfs(i, j+1)
		}
		return v
	}
	maxVal := 0
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				maxVal = max(maxVal, bfs(i, j))
			}
		}
	}
	return maxVal
}

// https://leetcode-cn.com/problems/surrounded-regions/
func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		board[i][j] = 0
		if i > 0 && board[i-1][j] == 'O' {
			dfs(i-1, j)
		}
		if i < m-1 && board[i+1][j] == 'O' {
			dfs(i+1, j)
		}
		if j > 0 && board[i][j-1] == 'O' {
			dfs(i, j-1)
		}
		if j < n-1 && board[i][j+1] == 'O' {
			dfs(i, j+1)
		}
	}
	for i := 0; i < n; i += 1 {
		if board[0][i] == 'O' {
			dfs(0, i)
		}
		if board[m-1][i] == 'O' {
			dfs(m-1, i)
		}
	}
	for i := 1; i < m-1; i += 1 {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}
	for i, row := range board {
		for j, v := range row {
			if v == 'O' {
				board[i][j] = 'X'
			} else if v == 0 {
				board[i][j] = 'O'
			}
		}
	}
}
