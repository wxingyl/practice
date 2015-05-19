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
/**
 * https://leetcode.com/problems/longest-consecutive-sequence/
 * Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
 * For example,
 * Given [100, 4, 200, 1, 3, 2],
 * The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.
 *
 * Your algorithm should run in O(n) complexity.
 */
int longestConsecutive(int* nums, int numsSize);

/**
 * https://leetcode.com/problems/minimum-size-subarray-sum/
 * Given an array of n positive integers and a positive integer s, find the minimal length of a subarray of which the sum ≥ s. If there isn't one, return 0 instead.
 *
 * For example, given the array [2,3,1,2,4,3] and s = 7,
 * the subarray [4,3] has the minimal length under the problem constraint.
 *
 * If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
 */
int minSubArrayLen(int s, int* nums, int numsSize); 
/**
 * https://leetcode.com/problems/two-sum/
 * Given an array of integers, find two numbers such that they add up to a specific target number.
 *
 * The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.
 *
 * You may assume that each input would have exactly one solution.
 *
 * Input: numbers={2, 7, 11, 15}, target=9
 * Output: index1=1, index2=2
 */
int* twoSum(int* nums, int numsSize, int target);
#endif
