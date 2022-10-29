package main

// https://leetcode.cn/problems/count-items-matching-a-rule/
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	i := 0
	if ruleKey == "color" {
		i = 1
	} else if ruleKey == "name" {
		i = 2
	}
	v := 0
	for _, s := range items {
		if s[i] == ruleValue {
			v++
		}
	}
	return v
}
