package main

import "strconv"

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
