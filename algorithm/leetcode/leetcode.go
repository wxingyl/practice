package main

import (
	"math"
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
	m := map[int][]int{}
	for _, v := range matches {
		a, ok := m[v[0]]
		if ok {
			a[0]++
		} else {
			m[v[0]] = []int{1, 0}
		}
		a, ok = m[v[1]]
		if ok {
			a[1]++
		} else {
			m[v[1]] = []int{0, 1}
		}
	}
	var allWin, lostOne []int
	for k, v := range m {
		if v[1] == 0 {
			allWin = append(allWin, k)
		}
		if v[1] == 1 {
			lostOne = append(lostOne, k)
		}
	}
	sort.Ints(allWin)
	sort.Ints(lostOne)
	return [][]int{allWin, lostOne}
}

//https://leetcode-cn.com/problems/maximum-candies-allocated-to-k-children/
func maximumCandies(candies []int, k int64) int {
	var sumVal int64
	for _, v := range candies {
		sumVal += int64(v)
	}
	if sumVal < k {
		return 0
	} else if sumVal < 2*k {
		return 1
	}
	pv := int(sumVal / k)
	sort.Ints(candies)
	var sv int64
	for _, v := range candies {
		if v >= pv {
			break
		}
		sv += int64(v)
	}
	return pv
}

//https://leetcode-cn.com/problems/unique-morse-code-words/
func uniqueMorseRepresentations(words []string) int {
	index := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := map[string]bool{}
	for _, w := range words {
		sb := strings.Builder{}
		for _, c := range w {
			sb.WriteString(index[c-'a'])
		}
		m[sb.String()] = true
	}
	return len(m)
}

//https://leetcode-cn.com/problems/largest-number-after-digit-swaps-by-parity/
func largestInteger(num int) int {
	var bs []int
	for v := num; v > 0; v /= 10 {
		bs = append(bs, v%10)
	}
	for i := len(bs) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if (bs[i]^bs[j])&1 == 0 && bs[i] < bs[j] {
				bs[i], bs[j] = bs[j], bs[i]
			}
		}
	}
	v := 0
	for i := len(bs) - 1; i >= 0; i-- {
		v = 10*v + bs[i]
	}
	return v
}

// https://leetcode-cn.com/problems/minimize-result-by-adding-parentheses-to-expression/
func minimizeResult(expression string) string {
	i := 1
	for ; expression[i] != '+'; i++ {
	}
	convert := func(s string) int {
		if len(s) == 0 {
			return 1
		} else {
			v, _ := strconv.Atoi(s)
			return v
		}
	}
	a, b := expression[:i], expression[i+1:]
	an, bn := len(a), len(b)
	v := -1
	var s string
	for i := 0; i < an; i++ {
		v1, v2 := convert(a[:i]), convert(a[i:])
		for j := 1; j <= bn; j++ {
			v3, v4 := convert(b[:j]), convert(b[j:])
			tv := v1 * (v2 + v3) * v4
			if v == -1 || tv < v {
				s = a[:i] + "(" + a[i:] + "+" + b[:j] + ")" + b[j:]
				v = tv
			}
		}
	}
	return s
}

// https://leetcode-cn.com/contest/cmbchina-2022spring/problems/fWcPGC/
func deleteText(article string, index int) string {
	if article[index] == ' ' {
		return article
	} else {
		n := len(article)
		var l, r string
		for i := index + 1; i < n; i++ {
			if article[i] == ' ' {
				r = article[i+1:]
				break
			}
		}
		for i := index - 1; i >= 0; i-- {
			if article[i] == ' ' {
				l = article[:i]
				break
			}
		}
		if len(l) == 0 {
			return r
		} else if len(r) == 0 {
			return l
		} else {
			return l + " " + r
		}
	}
}

// https://leetcode-cn.com/contest/cmbchina-2022spring/problems/ReWLAw/
func numFlowers(roads [][]int) int {
	dp := make([]int, len(roads)+1)
	ret := 0
	for _, v := range roads {
		dp[v[0]]++
		dp[v[1]]++
		if ret < dp[v[0]] {
			ret = dp[v[0]]
		}
		if ret < dp[v[1]] {
			ret = dp[v[1]]
		}
	}
	return ret + 1
}

//https://leetcode-cn.com/contest/cmbchina-2022spring/problems/Dk2Ytp/
func lightSticks(height int, width int, indices []int) []int {
	return nil
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool { return false }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return nil }

func deserialize(s string) *NestedInteger {
	ns := &NestedInteger{}
	if s[0] == '[' {
		n := len(s)
		var ls *NestedInteger
		for i, si := 1, 1; i < n; i++ {
			if s[i] == ',' || i == n-1 {
				if ls == nil || (*ls).IsInteger() {
					v, _ := strconv.Atoi(s[si:i])
					ls = &NestedInteger{}
					ls.SetInteger(v)
					ns.Add(*ls)
				}
				si = i + 1
			} else if s[i] == '[' {
				j, cnt := i+1, 1
				for {
					if s[j] == '[' {
						cnt++
					} else if s[j] == ']' {
						cnt--
						if cnt == 0 {
							break
						}
					}
					j++
				}
				ls = deserialize(s[i : j+1])
				ns.Add(*ls)
				i = j
				si = j + 1
			}
		}
	} else {
		v, _ := strconv.Atoi(s)
		ns.SetInteger(v)
	}
	return ns
}

//https://leetcode-cn.com/problems/4sum/
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	var tmp []int
	var ret [][]int
	var dfs func(i, t int)
	dfs = func(i, t int) {
		t -= nums[i]
		tl := len(tmp)
		if tl == 2 {
			for j := i + 1; j < n; j++ {
				if t == nums[j] {
					ret = append(ret, []int{tmp[0], tmp[1], nums[i], t})
				}
			}
		} else {
			tmp = append(tmp, nums[i])
			for j := i + 1; j < n; j++ {
				dfs(j, t)
			}
			tmp = tmp[:tl]
		}
	}
	for i := 0; i <= n-4; i++ {
		dfs(i, target)
	}
	return ret
}

func fourSumV2(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

//https://leetcode-cn.com/problems/most-common-word/
func mostCommonWord(paragraph string, banned []string) string {
	bm := map[string]int{}
	for _, s := range banned {
		bm[s] = math.MinInt
	}
	symble := map[byte]bool{}
	for _, c := range " !?',;." {
		symble[byte(c)] = true
	}
	var ret string
	cnt, n := 0, len(paragraph)
	var bs []byte
	for i := 0; i <= n; i++ {
		if i == n || symble[paragraph[i]] {
			if len(bs) > 0 {
				s := string(bs)
				bm[s]++
				if cnt < bm[s] {
					ret, cnt = s, bm[s]
				}
				bs = bs[:0]
			}
		} else {
			c := paragraph[i]
			if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
			}
			bs = append(bs, c)
		}
	}
	return ret
}

//https://leetcode-cn.com/problems/lexicographical-numbers/
func lexicalOrder(n int) []int {
	var ret []int
	var bfs func(pv int) bool
	bfs = func(pv int) bool {
		for l, r := pv*10, min(pv*10+9, n); l <= r; l++ {
			ret = append(ret, l)
			if len(ret) == n || bfs(l) {
				return true
			}
		}
		return false
	}
	for i := 1; i <= 9; i++ {
		ret = append(ret, i)
		if len(ret) == n || bfs(i) {
			break
		}
	}
	return ret
}

//https://leetcode-cn.com/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	n := len(nums)
	cache := make([]bool, n)
	var bfs func(tmp []int)
	bfs = func(tmp []int) {
		tn := len(tmp)
		tmp = append(tmp, 0)
		for i := 0; i < n; i++ {
			if cache[i] || (i > 0 && !cache[i-1] && nums[i-1] == nums[i]) {
				continue
			}
			tmp[tn] = nums[i]
			if tn+1 == n {
				t := make([]int, n)
				copy(t, tmp)
				ret = append(ret, t)
			} else {
				cache[i] = true
				bfs(tmp)
				cache[i] = false
			}
		}
	}
	bfs(nil)
	return ret
}

//https://leetcode-cn.com/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ret [][]int
	n := len(candidates)
	cache := make([]bool, n)
	var bfs func(li, t int, tmp []int)
	bfs = func(li, t int, tmp []int) {
		tn := len(tmp)
		tmp = append(tmp, 0)
		for i := li + 1; i < n && candidates[i] <= t; i++ {
			if cache[i] || (i > 0 && !cache[i-1] && candidates[i-1] == candidates[i]) {
				continue
			}
			tmp[tn] = candidates[i]
			if t == candidates[i] {
				t := make([]int, len(tmp))
				copy(t, tmp)
				ret = append(ret, t)
			} else {
				cache[i] = true
				bfs(i, t-candidates[i], tmp)
				cache[i] = false
			}
		}
	}
	bfs(-1, target, nil)
	return ret
}

//https://leetcode-cn.com/problems/shortest-distance-to-a-character/
func shortestToChar(s string, c byte) []int {
	n, lastMatch := len(s), -1
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == c {
			ret[i] = 0
			for j := i - 1; j > lastMatch; j-- {
				if ret[j] == 0 {
					ret[j] = i - j
				} else if i-j < ret[j] {
					ret[j] = i - j
				} else {
					break
				}
			}
			lastMatch = i
		} else if lastMatch >= 0 {
			ret[i] = i - lastMatch
		}
	}
	return ret
}

//https://leetcode-cn.com/problems/jump-game-ii/
func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	cache := make([]int, n)
	var dfs func(i int) int
	dfs = func(i int) int {
		maxVal := nums[i] + i
		if maxVal >= n-1 {
			cache[i] = 1
			return 1
		}
		v := n
		for j := 1; j <= nums[i]; j++ {
			if nums[i+j] != 0 {
				if cache[i+j] > 0 {
					v = min(cache[i+j], v)
				} else {
					v = min(dfs(i+j), v)
				}
			}
		}
		cache[i] = v + 1
		return cache[i]
	}
	return dfs(0)
}

//https://leetcode-cn.com/problems/combinations/
func combine(n int, k int) [][]int {
	var bfs func(nums []int, n, k int) [][]int
	bfs = func(nums []int, n, k int) [][]int {
		if k == 1 {
			ret := make([][]int, n)
			for i := 0; i < n; i++ {
				ret[i] = []int{nums[i]}
			}
			return ret
		} else if n == k {
			ret := make([]int, n)
			copy(ret, nums)
			return [][]int{ret}
		} else {
			ar := make([]int, n)
			var ret [][]int
			copy(ar, nums)
			for i := n - 1; i >= k-1; i-- {
				subArr := bfs(ar, i, k-1)
				for _, sr := range subArr {
					sr = append(sr, ar[i])
					ret = append(ret, sr)
				}
			}
			return ret
		}
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return bfs(nums, n, k)
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	s, n := []*Node{root}, 1
	for n > 0 {
		for i := 0; i < n; i++ {
			if s[i].Left != nil {
				s = append(s, s[i].Left)
			}
			if s[i].Right != nil {
				s = append(s, s[i].Right)
			}
		}
		s = s[n:]
		n = len(s)
		for i := 1; i < n; i++ {
			s[i-1].Next = s[i]
		}
	}
	return root
}

// https://leetcode-cn.com/problems/pascals-triangle-ii/
func getRow(rowIndex int) []int {
	n := rowIndex + 1
	ret := make([]int, n)
	ret[0] = 1
	half := n >> 1
	if n > 1 && (n&1) == 1 {
		half++
	}
	for i := 1; i <= half; i++ {
		ret[i] = ret[i-1] * (n - i) / i
	}
	for i := rowIndex; i > half; i-- {
		ret[i] = ret[rowIndex-i]
	}
	return ret
}

//https://leetcode-cn.com/problems/triangle/
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	p, dp := make([]int, n-1), make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		pn := i
		for j := 0; j < pn; j++ {
			p[j] = dp[j]
		}
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[j] = p[j]
			} else if j == i {
				dp[j] = p[i-1]
			} else {
				dp[j] = min(p[j], p[j-1])
			}
			dp[j] += triangle[i][j]
		}
	}
	v := dp[0]
	for i := 1; i < n; i++ {
		if v > dp[i] {
			v = dp[i]
		}
	}
	return v
}

// https://leetcode-cn.com/problems/reorder-data-in-log-files/
func reorderLogFiles(logs []string) []string {
	parse := func(s string) (bool, string) {
		i := strings.IndexByte(s, ' ')
		if s[i+1] >= '0' && s[i+1] <= '9' {
			return true, ""
		} else {
			return false, s[i+1:]
		}
	}
	sort.SliceStable(logs, func(i, j int) bool {
		numTag1, s1 := parse(logs[i])
		numTag2, s2 := parse(logs[j])
		if numTag1 && numTag2 {
			return false
		} else if numTag1 || numTag2 {
			return !numTag1
		} else {
			return s1 < s2 || s1 == s2 && logs[i] < logs[j]
		}
	})
	return logs
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *TreeNode) int {
	ret := 0
	var dfs func(node *TreeNode, v int)
	dfs = func(node *TreeNode, v int) {
		v = node.Val + 10*v
		if node.Left == nil && node.Right == nil {
			ret += v
			return
		}
		if node.Left != nil {
			dfs(node.Left, v)
		}
		if node.Right != nil {
			dfs(node.Right, v)
		}
	}
	if root != nil {
		dfs(root, 0)
	}
	return ret
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode-cn.com/problems/reorder-list/
func reorderList(head *ListNode) {
	var link []*ListNode
	for c := head; c != nil; c = c.Next {
		link = append(link, c)
	}
	n := len(link)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		link[i].Next, link[j].Next = link[j], link[i].Next
	}
	link[n>>1].Next = nil
}

//https://leetcode-cn.com/problems/repeated-dna-sequences/
func findRepeatedDnaSequences(s string) []string {
	const N, TAG = 10, 1<<20 - 1
	n := len(s)
	if n <= N {
		return nil
	}
	dict := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	m := map[int]int{}
	v := dict[s[0]]
	for i := 1; i < N; i++ {
		v = v<<2 | dict[s[i]]
	}
	m[v]++
	var ret []string
	for i := N; i < n; i++ {
		v = (v<<2 | dict[s[i]]) & TAG
		m[v]++
		if m[v] > 1 {
			ret = append(ret, s[i-9:i+1])
		}
	}
	return ret
}

//https://leetcode-cn.com/problems/compare-version-numbers/
func compareVersion(version1 string, version2 string) int {
	parse := func(s string) ([]int, int) {
		arr := strings.Split(s, ".")
		n := len(arr)
		vals := make([]int, n)
		for i, v := range arr {
			j, m := 0, len(v)
			for j < m && v[j] == '0' {
				j++
			}
			if j < m {
				vals[i], _ = strconv.Atoi(v[j:])
			}
		}
		for n > 0 && vals[n-1] == 0 {
			n--
		}
		return vals, n
	}
	v1, n1 := parse(version1)
	v2, n2 := parse(version2)
	for i := 0; i < n1 && i < n2; i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}
	if n1 == n2 {
		return 0
	} else if n1 > n2 {
		return 1
	} else {
		return -1
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  https://leetcode-cn.com/problems/insertion-sort-list/
func insertionSortList(head *ListNode) *ListNode {
	node, last := head.Next, head
	for node != nil {
		if last.Val > node.Val {
			nn := node.Next
			if head.Val > node.Val {
				node.Next, head = head, node
			} else {
				p := head
				for p.Next != nil && p.Next.Val <= node.Val {
					p = p.Next
				}
				p.Next, node.Next = node, p.Next
			}
			last.Next, node = nil, nn
		} else {
			last.Next = node
			last, node = node, node.Next
		}
	}
	return head
}

//https://leetcode-cn.com/problems/longest-repeating-character-replacement/
func characterReplacement(s string, k int) int {
	dict := [26]int{}
	n, i, j, maxV := len(s), 0, 0, 0
	for j < n {
		dict[s[j]-'A']++
		if maxV < dict[s[j]-'A'] {
			maxV = dict[s[j]-'A']
		}
		if j-i+1-maxV > k {
			dict[s[i]-'A']--
			i++
		}
		j++
	}
	return j - i
}

//https://leetcode-cn.com/problems/find-the-winner-of-the-circular-game/
func findTheWinner(n int, k int) int {
	q := make([]int, n)
	for i := 0; i < n; i++ {
		q[i] = i + 1
	}
	start := 0
	for n > 1 {
		start = (start + k - 1) % n
		for i := start; i < n-1; i++ {
			q[i] = q[i+1]
		}
		n--
	}
	return q[0]
}

//https://leetcode.cn/problems/one-away-lcci/
func oneEditAway(first string, second string) bool {
	n1, n2 := len(first), len(second)
	if abs(n1-n2) > 1 {
		return false
	}
	cnt := 0
	if n1 == n2 {
		for i := 0; i < n1; i++ {
			if first[i] != second[i] {
				cnt++
				if cnt > 1 {
					return false
				}
			}
		}
		return true
	} else if n1 > n2 {
		n1, first, second = n2, second, first
	}
	for i := 0; i < n1; {
		if first[i] == second[i+cnt] {
			i++
			continue
		}
		cnt++
		if cnt > 1 {
			return false
		}
	}
	return true
}
