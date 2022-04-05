package main

import (
	"math/bits"
	"sort"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/count-number-of-maximum-bitwise-or-subsets/
func countMaxOrSubsets(nums []int) int {
	tag := [31][]int{}
	for i, v := range nums {
		for j := 0; v > 0; v, j = v>>1, j+1 {
			if v&1 > 0 {
				tag[j] = append(tag[j], i)
			}
		}
	}
	m := map[int]bool{}
	var dfs func(i int) int
	dfs = func(i int) int {
		for ; i < 31; i++ {
			if len(tag[i]) > 0 {
				break
			}
		}
		if i == 31 {
			return 1
		}
		v := 0
		for _, j := range tag[i] {
			if m[j] {
				continue
			}
			m[j] = true
			v += dfs(i + 1)
			m[j] = false
		}
		return v
	}
	return dfs(0)
}

// https://leetcode-cn.com/problems/construct-string-from-binary-tree/
func tree2str(root *TreeNode) string {
	var bs []byte
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		bs = strconv.AppendInt(bs, int64(node.Val), 10)
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil {
			bs = append(bs, '(')
			dfs(node.Left)
			bs = append(bs, ')')
		} else {
			bs = append(bs, '(', ')')
		}
		if node.Right != nil {
			bs = append(bs, '(')
			dfs(node.Right)
			bs = append(bs, ')')
		}
	}
	if root != nil {
		dfs(root)
	}
	return string(bs)
}

//https://leetcode-cn.com/problems/word-break-ii/
func wordBreak(s string, wordDict []string) (ret []string) {
	n, m := len(s), map[string]bool{}
	for _, w := range wordDict {
		m[w] = true
	}
	var arr []string
	var bfs func(i int)
	bfs = func(i int) {
		if i >= n {
			return
		}
		for j := i + 1; j <= n; j++ {
			w := s[i:j]
			if m[w] {
				arr = append(arr, w)
				if j == n {
					ret = append(ret, strings.Join(arr, " "))
				} else {
					bfs(j)
				}
				arr = arr[:len(arr)-1]
			}
		}
	}
	bfs(0)
	sort.Strings(ret)
	return
}

// https://leetcode-cn.com/problems/max-points-on-a-line/
func maxPoints(points [][]int) int {
	return -1
}

//https://leetcode-cn.com/problems/array-of-doubled-pairs/
func canReorderDoubled(arr []int) bool {
	n := len(arr)
	m := map[int]int{}
	for _, v := range arr {
		m[v] += 1
	}
	if m[0]%2 == 1 {
		return false
	}
	sort.Ints(arr)
	left := 0
	for ; left < n && arr[left] < 0; left++ {
	}
	for i := left - 1; i > 0; i-- {
		cn := m[arr[i]]
		if cn == 0 {
			continue
		}
		v := 2 * arr[i]
		if m[v] < cn {
			return false
		}
		m[arr[i]], m[v] = 0, m[v]-cn
	}
	for ; left < n && arr[left] <= 0; left++ {
	}
	for i := left; i < n; i++ {
		cn := m[arr[i]]
		if cn == 0 {
			continue
		}
		v := 2 * arr[i]
		if m[v] < cn {
			return false
		}
		m[arr[i]], m[v] = 0, m[v]-cn
	}
	return true
}

// https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	var res [][]int
	var dfs func(sum, i int, ans []int)
	dfs = func(sum, i int, ans []int) {
		an := len(ans)
		for j := i; j < n; j++ {
			cv := sum + candidates[j]
			if cv == target {
				arr := make([]int, an+1)
				copy(arr, ans)
				arr[an] = candidates[j]
				res = append(res, arr)
			} else if cv < target {
				ans = append(ans, candidates[j])
				dfs(cv, j, ans)
				ans = ans[0:an]
			}
		}
	}
	dfs(0, 0, nil)
	return res
}

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if target < letters[0] || target >= letters[n-1] {
		return letters[0]
	}
	mv := 0
	for i, j := 0, n-1; i < j; {
		m := (i + j) >> 1
		if target < letters[m] {
			j = m - 1
		} else if target == letters[m] {
			mv = m
			break
		} else {
			i = m + 1
		}
		mv = m
	}
	for ; letters[mv] <= target; mv++ {
	}
	return letters[mv]
}

// https://leetcode-cn.com/problems/range-sum-query-mutable/
type NumArray struct {
	nums []int
	sum  []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = nums[i] + sum[i-1]
	}
	return NumArray{nums: nums, sum: sum}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.nums[index]
	this.nums[index] = val
	for ; index < len(this.nums); index++ {
		this.sum[index] += diff
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return this.sum[right] - this.sum[left-1]
	} else {
		return this.sum[right]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

//  https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/
func countPrimeSetBits(left int, right int) int {
	m := map[int]bool{2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true}
	cnt := 0
	for ; left <= right; left++ {
		if m[bits.OnesCount(uint(left))] {
			cnt++
		}
	}
	return cnt
}

//https://leetcode-cn.com/problems/minimum-number-of-operations-to-convert-time/
func convertTime(current string, correct string) int {
	parse := func(s string) (h, m int) {
		if s[0] == '0' {
			h = int(s[1] - '0')
		} else {
			h = int((s[0]-'0')*10 + s[1] - '0')
		}
		if s[3] == '0' {
			m = int(s[4] - '0')
		} else {
			m = int((s[3]-'0')*10 + s[4] - '0')
		}
		return
	}
	curH, curM := parse(current)
	corH, corM := parse(correct)
	cnt, stageVal := 0, corM-curM
	if curM > corM {
		stageVal += 60
		curH++
	}
	if stageVal >= 15 {
		v := stageVal / 15
		cnt += v
		stageVal -= v * 15
	}
	if stageVal >= 5 {
		v := stageVal / 5
		cnt += v
		stageVal -= v * 5
	}
	if stageVal >= 1 {
		cnt += stageVal
	}
	cnt += corH - curH
	if curH > corH {
		cnt += 24
	}
	return cnt
}

// https://leetcode-cn.com/problems/find-players-with-zero-or-one-losses/
func findWinners(matches [][]int) [][]int {
	return nil
}
