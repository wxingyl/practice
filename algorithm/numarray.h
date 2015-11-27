/*
 * Created by xing on 15/11/27.
 * https://leetcode.com/problems/range-sum-query-immutable/
 * Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

    Example:
        Given nums = [-2, 0, 3, -5, 2, -1]

        sumRange(0, 2) -> 1
        sumRange(2, 5) -> -1
        sumRange(0, 5) -> -3

    Note:
        You may assume that the array does not change.
        There are many calls to sumRange function.
 */

#ifndef CLION_TEST_NUMARRAY_H
#define CLION_TEST_NUMARRAY_H

#include <vector>

using namespace std;

class NumArray {

public:
    NumArray(vector<int> &nums);

    void update(int i, int val);

    int sumRange(int i, int j);

private:
    vector<int> nums;

    vector<int> subArray;

    int sum(int i);
};

#endif //CLION_TEST_NUMARRAY_H
