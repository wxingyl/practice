#ifndef _ACM_H
#define _ACM_H

//最长公共子序列
char* lcs(const char* str1, const int len1, const char* str2, const int len2);
//生成指定数目，范围的随机数，值取值为[min, max)
void init_rand(int*, const int, const int, const int);
/**
 * https://leetcode.com/problems/factorial-trailing-zeroes/
 * Given an integer n, return the number of trailing zeroes in n!.
 * Note: Your solution should be in logarithmic time complexity.
 */
int trailingZeroes(int n); 
/**
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/
 * Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
 *
 * For example, given the range [5, 7], you should return 4
 */
int rangeBitwiseAnd(int m, int n);
#endif
