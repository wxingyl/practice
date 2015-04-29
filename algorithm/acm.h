#ifndef _ACM_H
#define _ACM_H

#include <stdbool.h>

//最长公共子序列
char* lcs(const char* str1, const int len1, const char* str2, const int len2);
//生成指定数目，范围的随机数，值取值为[min, max)
void init_rand(int* a, const int size, const int min, const int max);
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
/**
 * https://leetcode.com/problems/rotate-array/
 * Rotate an array of n elements to the right by k steps.
 * For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
 * Note: Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
 */
void rotate(int nums[], int n, int k);
/**
 * https://leetcode.com/problems/happy-number/
 * Write an algorithm to determine if a number is "happy".
 *
 * A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
 *
 * Example: 19 is a happy number
 */
bool isHappy(int n);
/**
 * https://leetcode.com/problems/single-number/
 * Given an array of integers, every element appears twice except for one. Find that single one.
 * Note:
 * Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 */
int singleNumber(int* nums, int numsSize);
/**
 * https://leetcode.com/problems/single-number-ii
 * Given an array of integers, every element appears three times except for one. Find that single one.
 * Note:
 Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 */
int singleNumberll(int* nums, int numsSize); 
#endif
