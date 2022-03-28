package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problem-list/2ckc81c/
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

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sv := mean * (m + n)
	for _, v := range rolls {
		sv -= v
	}
	if sv < n || sv > 6*n {
		return nil
	}
	ret := make([]int, n)
	for sv >= n {
		for i := 0; i < n; i++ {
			ret[i] += 1
		}
		sv -= n
	}
	for i := 0; i < sv; i++ {
		ret[i] += 1
	}
	return ret
}

//https://leetcode-cn.com/problems/binary-number-with-alternating-bits/
func hasAlternatingBits(n int) bool {
	a := n ^ (n >> 1) + 1
	return a&(a-1) == 0
}

//https://leetcode-cn.com/problems/find-the-difference-of-two-arrays/
func findDifference(nums1 []int, nums2 []int) [][]int {
	m1, m2 := map[int]bool{}, map[int]bool{}
	for _, v := range nums1 {
		m1[v] = true
	}
	for _, v := range nums2 {
		m2[v] = true
	}
	var r1, r2 []int
	for _, v := range nums1 {
		if !m2[v] {
			r1 = append(r1, v)
			m2[v] = true
		}
	}
	for _, v := range nums2 {
		if !m1[v] {
			r2 = append(r2, v)
			m1[v] = true
		}
	}
	return [][]int{r1, r2}
}

// https://leetcode-cn.com/problems/minimum-deletions-to-make-array-beautiful/
func minDeletion(nums []int) int {
	n := len(nums)
	c := 0
	for i := 0; i < n; {
		j := i + 1
		for ; j < n && nums[i] == nums[j]; j++ {
			c++
		}
		if j == n {
			c++
			break
		}
		i = j + 1
	}
	return c
}

//https://leetcode-cn.com/problems/find-palindrome-with-fixed-length/
func kthPalindrome(queries []int, intLength int) []int64 {
	// ret := make([]int64, len(queries))
	// n := intLength >> 1
	// p10Val := int(math.Pow10(n))
	// stageMaxCnt := p10Val - p10Val/10
	// bs := make([]byte, intLength)
	// if intLength%2 == 1 {
	// 	maxCnt := stageMaxCnt * 10
	// 	for i, q := range queries {
	// 		if q > maxCnt {
	// 			ret[i] = -1
	// 			continue
	// 		}
	// 		cnt := q % stageMaxCnt
	// 		c := (q / stageMaxCnt)
	// 		if cnt == 0 && c > 0 {
	// 			c--
	// 			cnt = stageMaxCnt
	// 		}
	// 		bs[n] = byte('0' + c)
	// 		cnt += (cnt - 1) / 9
	// 		s := strconv.Itoa(cnt)
	// 		for j, sl := 1, n-len(s); j <= n; j++ {
	// 			if j > sl {
	// 				bs[n-j], bs[n+j] = s[j-sl-1], s[j-sl-1]
	// 			} else {
	// 				bs[n-j], bs[n+j] = '0', '0'
	// 			}
	// 		}
	// 		ret[i], _ = strconv.ParseInt(string(bs), 10, 0)
	// 	}
	// } else {
	// 	for i, q := range queries {
	// 		if q > stageMaxCnt {
	// 			ret[i] = -1
	// 			continue
	// 		}
	// 		q += (q - 1) / 9
	// 		s := strconv.Itoa(q)
	// 		for j, sl := 1, n-len(s); j <= n; j++ {
	// 			if j > sl {
	// 				bs[n-j], bs[n+j] = s[j-sl-1], s[j-sl-1]
	// 			} else {
	// 				bs[n-j], bs[n+j] = '0', '0'
	// 			}
	// 		}
	// 		ret[i], _ = strconv.ParseInt(string(bs), 10, 0)
	// 	}
	// }
	// return ret
	return nil
}

// https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam/
func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	if k >= n>>1 {
		return n
	}
	find := func(t byte) int {
		ret, si, kv := 0, 0, 0
		for i, c := range answerKey {
			if byte(c) != t {
				kv++
				for kv > k {
					if answerKey[si] != t {
						kv--
					}
					si++
				}
			}
			ret = max(ret, i-si+1)
		}
		return ret
	}
	return max(find('F'), find('T'))
}
