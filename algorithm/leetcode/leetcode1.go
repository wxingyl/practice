package main

import (
	"bytes"
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/contest/weekly-contest-91/problems/shortest-subarray-with-sum-at-least-k/
func shortestSubarray(A []int, K int) int {
	m := len(A)
	p := make([]int, m+1)
	for i := 1; i < m+1; i++ {
		p[i] += p[i-1] + A[i-1]
	}
	retVal := -1
	queue := list.New()
	//从0开始建立，维护双向队列存储可以作为开始元素的前一位数的下标，每次便利队列，更新期待最小值，同时删除后续无须比较的元素（负值肯定不会作为开始值的）
	for i := 0; i <= m; i++ {
		for queue.Len() > 0 {
			index := queue.Front().Value.(int)
			if p[i]-p[index] < K {
				break
			}
			if retVal == -1 || retVal > i-index {
				retVal = i - index
			}
			queue.Remove(queue.Front())
		}
		for queue.Len() > 0 {
			index := queue.Back().Value.(int)
			if p[index] < p[i] {
				break
			}
			queue.Remove(queue.Back())
		}
		queue.PushBack(i)
	}
	return retVal
}

//https://leetcode-cn.com/contest/weekly-contest-91/problems/score-after-flipping-matrix/
func matrixScore(A [][]int) int {
	m := len(A)
	n := len(A[0])
	for i := 0; i < m; i++ {
		if A[i][0] == 1 {
			continue
		}
		for j := 0; j < n; j++ {
			A[i][j] ^= 1
		}
	}
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < m; j++ {
			if A[j][i] == 1 {
				count += 1
			}
		}
		if count <= m/2 {
			for j := 0; j < m; j++ {
				A[j][i] ^= 1
			}
		}
	}
	val := 0
	for i := 0; i < m; i++ {
		k := 1
		for j := n - 1; j >= 0; j-- {
			if A[i][j] == 1 {
				val += k
			}
			k *= 2
		}
	}
	return val
}

//https://leetcode-cn.com/problems/sum-of-two-integers/description/
func getSum(a int, b int) int {
	var c int
	var d int
	c = a ^ b
	d = (a & b) << 1
	if c == 0 {
		return d
	} else if d == 0 {
		return c
	} else {
		return getSum(c, d)
	}
}

//https://leetcode-cn.com/contest/weekly-contest-91/problems/all-nodes-distance-k-in-binary-tree/

//通过数组构建二叉树
func buildTree(treeArr []int) []TreeNode {
	treeArrLen := len(treeArr)
	tree := make([]TreeNode, len(treeArr))
	for i := 0; i < treeArrLen; i++ {
		if treeArr[i] == -1 {
			continue
		}
		tree[i] = TreeNode{
			Val: treeArr[i],
		}
	}
	preIndex := 0
	for n := 1; preIndex < treeArrLen; n *= 2 {
		nextLevel := preIndex + n
		if nextLevel > treeArrLen {
			nextLevel = treeArrLen
		}
		for i := 0; i < n; i++ {
			index := i + preIndex
			if index >= treeArrLen {
				break
			}
			val := treeArr[index]
			if val == -1 {
				continue
			}
			childIndex := preIndex + n + 2*i
			if childIndex+1 < treeArrLen {
				tree[index].Left = &tree[childIndex]
				tree[index].Right = &tree[childIndex+1]
			} else if childIndex < treeArrLen {
				tree[index].Left = &tree[childIndex]
			}
		}
		preIndex += n
	}
	return tree
}

func getPrefixNode(statMap map[int]int, val int) []int {
	prefix := []int{val}
	for statMap[val] != -1 {
		val = statMap[val]
		prefix = append([]int{val}, prefix...)
	}
	return prefix
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	statMap := make(map[int]int)
	node := root
	nodeList := list.New()
	preVal := -1
	for node != nil || nodeList.Len() > 0 {
		if node != nil {
			statMap[node.Val] = preVal
			nodeList.PushBack(node)
			preVal = node.Val
			node = node.Left
		} else {
			node = nodeList.Remove(nodeList.Back()).(*TreeNode)
			preVal = node.Val
			node = node.Right
		}
	}
	targetPrefix := getPrefixNode(statMap, target.Val)
	tpLen := len(targetPrefix)
	var retArr []int
	for k := range statMap {
		var commonLen int
		localPrefix := getPrefixNode(statMap, k)
		lpLen := len(localPrefix)
		for commonLen = 0; commonLen < tpLen && commonLen < lpLen; commonLen++ {
			if targetPrefix[commonLen] != localPrefix[commonLen] {
				break
			}
		}
		if tpLen+lpLen-commonLen*2 == K {
			retArr = append(retArr, k)
		}
	}
	return retArr
}

//https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
func lengthOfLIS(nums []int) int {
	numLen := len(nums)
	if numLen == 0 {
		return 0
	} else if numLen == 1 {
		return 1
	}
	dp := make([]int, numLen)
	cn := make([]int, numLen)
	dp[0] = 1
	cn[0] = 1
	maxLen := 1
	for i := 1; i < numLen; i++ {
		dp[i] = 1
		cn[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] <= dp[j] {
					dp[i] = dp[j] + 1
					cn[i] = cn[j]
				} else if dp[j]+1 == dp[i] {
					cn[i] += cn[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

//https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/description/
func findNumberOfLIS(nums []int) int {
	numLen := len(nums)
	if numLen == 0 {
		return 0
	} else if numLen == 1 {
		return 1
	}
	dp := make([]int, numLen)
	cn := make([]int, numLen)
	dp[0] = 1
	cn[0] = 1
	maxLen := 1
	for i := 1; i < numLen; i++ {
		dp[i] = 1
		cn[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] <= dp[j] {
					dp[i] = dp[j] + 1
					cn[i] = cn[j]
				} else if dp[j]+1 == dp[i] {
					cn[i] += cn[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	retVal := 0
	for i := 0; i < numLen; i++ {
		if dp[i] == maxLen {
			retVal += cn[i]
		}
	}
	return retVal
}

//issue:  https://leetcode-cn.com/problems/longest-palindromic-substring/description/
func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen <= 1 {
		return s
	}
	tag := make([][]bool, strLen)
	for i := 0; i < strLen; i++ {
		tag[i] = make([]bool, strLen-i)
		tag[i][0] = true
		if i+1 != strLen && s[i] == s[i+1] {
			tag[i][1] = true
		}
	}
	for n := 2; n < strLen; n++ {
		for i := 0; i < strLen-n; i++ {
			tag[i][n] = tag[i+1][n-2] && s[i] == s[i+n]
		}
	}
	start := 0
	end := 0
	for i := 1; i < strLen; i++ {
		for j := 0; j < i; j++ {
			if tag[j][i-j] && (i-j) > (end-start) {
				start = j
				end = i
			}
		}
	}
	return s[start : end+1]
}

//issue: https://leetcode-cn.com/problems/largest-divisible-subset/description/
func largestDivisibleSubset(nums []int) []int {
	var numLen = len(nums)
	if numLen == 0 {
		return []int{}
	} else if numLen == 1 {
		return nums
	}
	sort.Ints(nums)
	count := make([]int, numLen)
	count[0] = 1
	maxNum := 1
	maxVal := 0
	for i := 1; i < numLen; i++ {
		count[i] = 1
		v := nums[i]
		for j := 0; j < i; j++ {
			if v%nums[j] == 0 {
				count[i] = max(count[j]+1, count[i])
				if maxNum < count[i] {
					maxNum = count[i]
					maxVal = v
				}
			}
		}
	}
	retArr := make([]int, maxNum)
	for i := numLen - 1; i >= 0; i-- {
		if count[i] == maxNum && maxVal%nums[i] == 0 {
			maxNum--
			maxVal = nums[i]
			retArr[maxNum] = maxVal
		}
	}
	return retArr
}

//https://leetcode-cn.com/problems/champagne-tower/description/
func champagneTower(poured int, query_row int, query_glass int) float64 {
	tag := [101][101]float64{}
	tag[0][0] = float64(poured)
	for i := 0; i < query_row; i++ {
		for j := 0; j <= i; j++ {
			q := float64(tag[i][j]-1) / 2
			if q > 0 {
				tag[i+1][j] += q
				tag[i+1][j+1] += q
			}
		}
	}
	if tag[query_row][query_glass] > 1 {
		return 1
	} else {
		return tag[query_row][query_glass]
	}
}

//https://leetcode-cn.com/problems/array-nesting/description/
func arrayNesting(nums []int) int {
	for i := 0; i < len(nums); i++ {

	}
	return 1
}

//https://leetcode-cn.com/problems/longest-absolute-file-path/description/
//"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
//dir
//	subdir1
//	file1.ext
//	subsubdir1
//	subdir2
//	subsubdir2
//	file2.ext
func lengthLongestPath(input string) int {
	maxLen := 0
	arr := make([]int, 10)
	for _, s := range strings.Split(input, "\n") {
		level := strings.LastIndexByte(s, '\t') + 1
		if level >= len(arr) {
			arr = append(arr, 0)
		}
		sLen := len(s) - level
		if strings.LastIndexByte(s, '.') >= 0 {
			var v int
			if level > 0 {
				v = sLen + level + arr[level-1]
			} else {
				v = sLen
			}
			if maxLen < v {
				maxLen = v
			}
		} else {
			if level > 0 {
				arr[level] = sLen + arr[level-1]
			} else {
				arr[level] = sLen
			}
		}
	}
	return maxLen
}

//https://leetcode-cn.com/problems/escape-the-ghosts/description/
func escapeGhosts(ghosts [][]int, target []int) bool {
	targetLen := abs(target[0]) + abs(target[1])
	for _, ghost := range ghosts {
		if abs(target[0]-ghost[0])+abs(target[1]-ghost[1]) <= targetLen {
			return false
		}
	}
	return true
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

//https://leetcode-cn.com/problems/number-of-matching-subsequences/description/
func numMatchingSubseq(S string, words []string) int {
	a := 'a'
	fmt.Println(a)
	return 0
}

//https://leetcode-cn.com/problems/excel-sheet-column-number/description/
func titleToNumber(s string) int {
	v := 0
	for _, letter := range s {
		v *= 26
		v += int(letter - 'A' + 1)
	}
	return v
}

//https://leetcode-cn.com/problems/elimination-game/description/
func lastRemaining(n int) int {
	if n <= 2 {
		return n
	}
	half := n / 2
	return 2 * lastRemainingFlag(half, true)
}

func lastRemainingFlag(n int, flag bool) int {
	if n <= 1 {
		return 1
	}
	v := 2 * lastRemainingFlag(n/2, !flag)
	if flag && n&1 == 0 {
		v -= 1
	}
	return v
}

//https://leetcode-cn.com/problems/super-pow/
func superPow(a int, b []int) int {
	v := a % 1337
	if v <= 1 {
		return v
	}
	n := len(b)
	y := v ^ b[n-1]
	for i := n - 2; i >= 0; i -= 1 {
		y *= (v ^ b[i]) * (v ^ 10)
	}
	return y % 1337
}

func powMod(v int, n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret = ret * v % 1337
	}
	return ret
}

//https://leetcode-cn.com/problems/range-sum-query-2d-immutable/

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
type NumMatrix struct {
	Matrix [][]int
}

func NumMatrixConstructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	if row > 0 {
		col := len(matrix[0])
		if col > 0 {
			for i := 0; i < row; i += 1 {
				val := matrix[i][0]
				for j := 1; j < col; j += 1 {
					matrix[i][j] += val
					val = matrix[i][j]
				}
			}
			for j := 0; j < col; j += 1 {
				val := matrix[0][j]
				for i := 1; i < row; i += 1 {
					matrix[i][j] += val
					val = matrix[i][j]
				}
			}
		}
	}
	m := NumMatrix{
		Matrix: matrix,
	}
	return m
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row := len(this.Matrix)
	if row == 0 {
		return 0
	} else if row2 >= row {
		row2 = row - 1
	}
	col := len(this.Matrix[0])
	if col == 0 {
		return 0
	}
	if col2 >= col {
		col2 = col - 1
	}
	row1 -= 1
	col1 -= 1
	val := 0
	if row1 >= 0 {
		val += this.Matrix[row1][col2]
	}
	if col1 >= 0 {
		val += this.Matrix[row2][col1]
	}
	if row1 >= 0 && col1 >= 0 {
		val -= this.Matrix[row1][col1]
	}
	return this.Matrix[row2][col2] - val
}

//https://leetcode-cn.com/problems/linked-list-cycle-ii/
//Floyd 算法: 当然一个跑得快的人和一个跑得慢的人在一个圆形的赛道上赛跑，会发生什么？在某一个时刻，跑得快的人一定会从后面赶上跑得慢的人
type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Value interface{}
	Next  *Node
}

type Stack struct {
	topNode *Node
}

func (s *Stack) push(value interface{}) {
	s.topNode = &Node{
		Value: value,
		Next:  s.topNode,
	}
}

func (s *Stack) top() interface{} {
	if s.topNode == nil {
		return nil
	} else {
		return s.topNode.Value
	}
}

func (s *Stack) pop() interface{} {
	if s.topNode == nil {
		return nil
	} else {
		node := s.topNode
		s.topNode, node.Next = s.topNode.Next, nil
		return node.Value
	}
}

func (s *Stack) empty() bool {
	return s.topNode == nil
}

func linkPrint(head *ListNode) {
	fmt.Printf("%d", head.Val)
	for head = head.Next; head != nil; head = head.Next {
		fmt.Printf(", %d", head.Val)
	}
	fmt.Println()
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for {
		fast = fast.Next
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next
		slow = slow.Next
		if fast != slow {
			break
		}
	}
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

//https://leetcode-cn.com/problems/lru-cache/
//["LRUCache","get","put","get","put","put","get","get"]
//[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
//[null,-1,null,-1,null,null,2,6]
type LinkNode struct {
	Val  int
	Key  int
	Prev *LinkNode
	Next *LinkNode
}

type LRUCache struct {
	Kv       map[int]*LinkNode
	Capacity int
	Tail     *LinkNode
	Head     *LinkNode
}

func LRUCacheConstructor(capacity int) LRUCache {
	kv := make(map[int]*LinkNode, capacity)
	cache := LRUCache{
		Kv:       kv,
		Capacity: capacity,
		Head:     &LinkNode{},
		Tail:     &LinkNode{},
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head
	return cache
}

func (p *LRUCache) appendTail(node *LinkNode) {
	oldLast := p.Tail.Prev
	if oldLast == nil {
		p.Tail.Prev = node
		node.Prev = p.Head
		node.Next = p.Tail
		p.Head.Next = node
	} else if oldLast == node {
		return
	} else {
		if node.Prev != nil {
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
		}
		oldLast.Next = node
		node.Prev = oldLast
		node.Next = p.Tail
		p.Tail.Prev = node
	}
}

func (p *LRUCache) Get(key int) int {
	node := p.Kv[key]
	if node == nil {
		return -1
	}
	p.appendTail(node)
	return node.Val
}

func (p *LRUCache) Put(key int, value int) {
	if p.Capacity <= 0 {
		return
	}
	node := p.Kv[key]
	if node == nil {
		for len(p.Kv) >= p.Capacity {
			delete(p.Kv, p.Head.Next.Key)
			p.Head.Next = p.Head.Next.Next
			p.Head.Next.Prev = p.Head
		}
		node = &LinkNode{
			Val: value,
			Key: key,
		}
		p.Kv[key] = node
	} else {
		node.Val = value
	}
	p.appendTail(node)
}

//https://leetcode-cn.com/problems/sort-list/
//[4,19,14,5,-3,1,8,5,11,15]
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		if head.Val > head.Next.Val {
			head, head.Next, head.Next.Next = head.Next, nil, head
		}
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	fast, slow.Next = slow.Next, nil
	return linkMerge(sortList(head), sortList(fast))
}

func linkMerge(lp *ListNode, rp *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	for l, r := lp, rp; node != nil; node = node.Next {
		if l == nil {
			node.Next = r
			break
		} else if r == nil {
			node.Next = l
			break
		} else if l.Val <= r.Val {
			node.Next, l = l, l.Next
		} else {
			node.Next, r = r, r.Next
		}
	}
	return head.Next
}

//https://leetcode-cn.com/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	res, up, down := nums[0], nums[0], nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		val := nums[i]
		if val > 0 {
			up, down = max(up*val, val), min(down*val, val)
		} else {
			up, down = max(down*val, val), min(up*val, val)
		}
		res = max(res, up)
	}
	return res
}

//https://leetcode-cn.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				numIslandsDFS(grid, i, j, m, n)
				cnt++
			}
		}
	}
	return cnt
}

type NumIslandsStackNode struct {
	x    int
	y    int
	Next *NumIslandsStackNode
}

func numIslandsDFS(grid [][]byte, i int, j int, m int, n int) {
	top := &NumIslandsStackNode{
		x: i,
		y: j,
	}
	m, n = m-1, n-1
	for top != nil {
		i, j = top.x, top.y
		grid[i][j] = '0'
		top, top.Next = top.Next, nil
		if j > 0 && grid[i][j-1] == '1' {
			top = &NumIslandsStackNode{x: i, y: j - 1, Next: top}
		}
		if i > 0 && grid[i-1][j] == '1' {
			top = &NumIslandsStackNode{
				x:    i - 1,
				y:    j,
				Next: top,
			}
		}
		if i < m && grid[i+1][j] == '1' {
			top = &NumIslandsStackNode{
				x:    i + 1,
				y:    j,
				Next: top,
			}
		}
		if j < n && grid[i][j+1] == '1' {
			top = &NumIslandsStackNode{
				x:    i,
				y:    j + 1,
				Next: top,
			}
		}
	}
}

//https://leetcode-cn.com/problems/decode-string/
type DecodeStringStackNode struct {
	Num    int
	Prefix string
	Next   *DecodeStringStackNode
}
type DecodeString struct {
	Num    int
	Prefix string
}

func decodeString(s string) string {
	if len(s) == 0 {
		return s
	}
	startNum := -1
	num := 0
	top := &DecodeStringStackNode{}
	for i, c := range s {
		if c >= '0' && c <= '9' {
			if startNum == -1 {
				startNum = i
			}
		} else if c == '[' {
			num, _ = strconv.Atoi(s[startNum:i])
			startNum = -1
			top = &DecodeStringStackNode{Num: num, Next: top}
		} else if c == ']' {
			str := strings.Repeat(top.Prefix, top.Num)
			top = top.Next
			top.Prefix += str
		} else {
			top.Prefix += string(c)
		}
	}
	return top.Prefix
}

//https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	retArray := make([][]int, 0, 10)
	retMap := make(map[string]bool, 10)
	for i, j := 0, len(nums)-1; i+1 < j && nums[i] <= 0 && nums[j] >= 0; {
		val := nums[i] + nums[i+1] + nums[j]
		if val == 0 {
			//OK了
			strKey := fmt.Sprintf("%d-%d", nums[i], nums[i+1])
			_, ok := retMap[strKey]
			if !ok {
				retMap[strKey] = true
				retArray = append(retArray, []int{nums[i], nums[i+1], nums[j]})
			}
		} else if val > 0 {
			j -= 1
			continue
		} else if nums[i]+nums[j-1]+nums[j] < 0 {
			i += 1
			continue
		}
		r := j
		for k := i + 2; k < r && nums[r] >= 0; {
			val = nums[i] + nums[k] + nums[r]
			if val == 0 {
				strKey := fmt.Sprintf("%d-%d", nums[i], nums[k])
				_, ok := retMap[strKey]
				if !ok {
					retMap[strKey] = true
					retArray = append(retArray, []int{nums[i], nums[k], nums[r]})
				}
				k += 1
			} else if val > 0 {
				r -= 1
			} else {
				k += 1
			}
		}
		i += 1
	}
	return retArray
}

//https://leetcode-cn.com/problems/01-matrix/
func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	dis := make([][]int, m)
	for i := 0; i < m; i += 1 {
		dis[i] = make([]int, n)
		for j := 0; j < n; j += 1 {
			if matrix[i][j] == 0 {
				dis[i][j] = 0
			} else {

				dis[i][j] = math.MaxInt32 - 10000
			}
		}
	}
	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			if matrix[i][j] == 0 {
				dis[i][j] = 0
			} else {
				if i > 0 {
					dis[i][j] = min(dis[i-1][j]+1, dis[i][j])
				}
				if j > 0 {
					dis[i][j] = min(dis[i][j-1]+1, dis[i][j])
				}
			}
		}
	}
	m, n = m-1, n-1
	for i := m; i >= 0; i -= 1 {
		for j := n; j >= 0; j -= 1 {
			if i < m {
				dis[i][j] = min(dis[i+1][j]+1, dis[i][j])
			}
			if j < n {
				dis[i][j] = min(dis[i][j+1]+1, dis[i][j])
			}
		}
	}
	return dis
}

type RichNode struct {
	i     int
	quite int
	child *list.List
}

//https://leetcode-cn.com/problems/loud-and-rich/
func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	answer := make([]int, n)
	tree := make([]*RichNode, n)
	for _, r := range richer {
		if tree[r[1]] == nil {
			tree[r[1]] = &RichNode{i: r[1], quite: -1, child: list.New()}
		}
		if tree[r[0]] == nil {
			tree[r[0]] = &RichNode{i: r[0], quite: -1, child: list.New()}
		}
		tree[r[1]].child.PushFront(tree[r[0]])
	}
	for i := 0; i < n; i += 1 {
		r := tree[i]
		if r == nil {
			answer[i] = i
			continue
		} else if r.quite == -1 {
			if r.child.Len() == 0 {
				r.quite = i
			} else {
				r.quite = loudAndRichDnf(r, quiet)
			}
		}
		answer[i] = r.quite
	}
	return answer
}

func loudAndRichDnf(node *RichNode, quiet []int) int {
	if node.quite >= 0 {
		return node.quite
	} else if node.child.Len() == 0 {
		return node.i
	} else {
		index := -1
		for r := node.child.Front(); r != nil; r = r.Next() {
			c := r.Value.(*RichNode)
			c.quite = loudAndRichDnf(c, quiet)
			if index == -1 || quiet[index] > quiet[c.quite] {
				index = c.quite
			}
		}
		if quiet[index] > quiet[node.i] {
			return node.i
		} else {
			return index
		}
	}
}

//https://leetcode-cn.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {
	const maxInt = 2147483647
	startIndex, negative := -1, -1
	for i := 0; i < len(str); i += 1 {
		c := str[i]
		if c >= '0' && c <= '9' {
			startIndex = i
			if c == '0' {
				continue
			}
			break
		} else if c == '-' || c == '+' {
			if negative != -1 || startIndex != -1 {
				return 0
			}
			if c == '-' {
				negative = 1
			} else {
				negative = 0
			}
		} else if negative != -1 || c != ' ' || startIndex != -1 {
			return 0
		}
	}
	if startIndex == -1 {
		return 0
	}
	if negative == -1 {
		negative = 0
	}
	//-2147483648
	val, modeVal, maxVal, maxModeVal := 0, (int)(str[startIndex]-'0'), maxInt/10, maxInt%10+negative
	if startIndex+1 < len(str) {
		for _, c := range str[startIndex+1:] {
			if c < '0' || c > '9' {
				break
			}
			val = val*10 + modeVal
			modeVal = int(uint8(c - '0'))
			//越值检查
			if val > maxVal || (val == maxVal && modeVal >= maxModeVal) {
				if negative == 1 {
					return -maxInt - 1
				} else {
					return maxInt
				}
			}
		}
	}
	if negative == 1 {
		return 0 - val*10 - modeVal
	} else {
		return val*10 + modeVal
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/
 */
func swapPairs(head *ListNode) *ListNode {
	var pre *ListNode
	for cur := head; cur != nil && cur.Next != nil; pre, cur = cur, cur.Next {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = cur
		if pre == nil {
			head = tmp
		} else {
			pre.Next = tmp
		}
	}
	return head
}

/**
 *	https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
 */
func searchRange(nums []int, target int) []int {
	from, to, t := 0, len(nums)-1, -1
	for from <= to {
		mid := (to + from) >> 1
		if target == nums[mid] {
			t = mid
			break
		} else if target > nums[mid] {
			from = mid + 1
		} else {
			to = mid - 1
		}
	}
	if t == -1 {
		return []int{-1, -1}
	}
	start, end := t, t
	for from < start {
		mid := (start + from) >> 1
		if target == nums[mid] {
			start = mid
		} else {
			from = mid + 1
		}
	}
	for end < to {
		mid := (end + to + 1) >> 1
		if target == nums[mid] {
			end = mid
		} else {
			to = mid - 1
		}
	}
	return []int{start, end}
}

//https://leetcode-cn.com/problems/integer-to-roman/
func intToRoman(num int) string {
	var buffer bytes.Buffer
	sy := []rune{'I', 'V', 'X', 'L', 'C', 'D', 'M', ' ', ' '}
	for i := 0; num > 0; num, i = num/10, i+1 {
		v := num % 10
		m, n := v%5, v/5
		f, s, t := sy[2*i], sy[2*i+1], sy[2*i+2]
		if m == 4 {
			if n == 1 {
				buffer.WriteRune(t)
				buffer.WriteRune(f)
			} else {
				buffer.WriteRune(s)
				buffer.WriteRune(f)
			}
		} else {
			for j := 0; j < m; j += 1 {
				buffer.WriteRune(f)
			}
			if n == 1 {
				buffer.WriteRune(s)
			}
		}
	}
	byteArray := buffer.Bytes()
	l := len(byteArray)
	ret := make([]byte, l)
	for i := l; i > 0; i-- {
		ret[l-i] = byteArray[i-1]
	}
	return string(ret)
}

//https://leetcode-cn.com/problems/partition-array-into-disjoint-intervals/
func partitionDisjoint(A []int) int {
	i, j := 0, len(A)-1
	l, r := A[0], A[j]
	for ; i < j; i, j = i+1, j-1 {
		if l < A[i] {
			l = A[i]
		}
		if r > A[j] {
			r = A[j]
		}
		if l > r {
			break
		}
	}
	return i + 1
}
