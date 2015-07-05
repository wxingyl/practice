#ifndef _ACM_H
#define _ACM_H

#include <stdbool.h>

void printChars(char** s, int len);
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

struct ListNode {
	int val;
	struct ListNode *next;
};

void free_list_node(struct ListNode* l);

void print_list_node(struct ListNode* l);

/**
 *	https://leetcode.com/problems/add-two-numbers/
 *	You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 *
 *	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 *	Output: 7 -> 0 -> 8
 */
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2);
/**
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/
 * Given a string, find the length of the longest substring without repeating characters. For example, the longest substring without repeating letters for "abcabcbb" is "abc", which the length is 3. For "bbbbb" the longest substring is "b", with the length of 1.
 */
int lengthOfLongestSubstring(char* s);
/**
 * https://leetcode.com/problems/median-of-two-sorted-arrays/
 * There are two sorted arrays nums1 and nums2 of size m and n respectively. Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
 */
double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size);
/**
 * https://leetcode.com/problems/longest-palindromic-substring/
 * Given a string S, find the longest palindromic substring in S. You may assume that the maximum length of S is 1000, and there exists one unique longest palindromic substring.
 */
char* longestPalindrome(char* s);
/**
 * https://leetcode.com/problems/zigzag-conversion/
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 * And then read line by line: "PAHNAPLSIIGYIR"
 * Write the code that will take a string and make this conversion given a number of rows:
 *
 * string convert(string text, int nRows);
 * convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
 */
char* convert(char* s, int numRows); 
/**
 * https://leetcode.com/problems/reverse-integer/
 * Reverse digits of an integer.
 *
 * Example1: x = 123, return 321
 * Example2: x = -123, return -321
 *
 * Have you thought about this?
 * Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!
 *
 * If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.
 *
 * Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer, then the reverse of 1000000003 overflows. How should you handle such cases?
 *
 * For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
 */
int reverse(int x);

/**
 * https://leetcode.com/problems/palindrome-number/
 * Determine whether an integer is a palindrome. Do this without extra space.
 *
 * click to show spoilers.
 *
 * Some hints:
 * Could negative integers be palindromes? (ie, -1)
 *
 * If you are thinking of converting the integer to string, note the restriction of using extra space.
 *
 * You could also try reversing an integer. However, if you have solved the problem "Reverse Integer", you know that the reversed integer might overflow. How would you handle such case?
 *
 * There is a more generic way of solving this problem.
 */
bool isPalindrome(int x);

/**
 * https://leetcode.com/problems/container-with-most-water/
 * Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
 *
 * Note: You may not slant the container.
 */
int maxArea(int* height, int heightSize);
/**
 * https://leetcode.com/problems/longest-common-prefix/
 */
char* longestCommonPrefix(char** strs, int strsSize);
/**
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/
 */
struct ListNode* removeNthFromEnd(struct ListNode* head, int n);
/**
 * https://leetcode.com/problems/generate-parentheses/
 * Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
 *
 * For example, given n = 3, a solution set is:
 *
 * "((()))", "(()())", "(())()", "()(())", "()()()"
 */
char** generateParenthesis(int n, int* returnSize);
#endif
