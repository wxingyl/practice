//
// Created by xing on 15/11/27.
//
#include "numarray.h"

NumArray::NumArray(vector<int> &nums) {
    this->nums = nums;
    vector<int>::size_type len = nums.size();
    for (unsigned int i = 0; i < len; i++) {
        if ((i & 1) == 0) {
            this->subArray.push_back(nums.at(i));
        } else {
            int j = int(i + 1);
            this->subArray.push_back(nums.at(i) + this->sumRange(i + 1 - (j & -j), i - 1));
        }
    }
}

void NumArray::update(int i, int val) {
    const int diff = val - this->nums[i];
    this->nums[i] = val;
    i += 1;
    const vector<int>::size_type size = this->subArray.size();
    while (i <= size) {
        subArray[i - 1] += diff;
        i += i & -i;
    }
}

int NumArray::sumRange(int i, int j) {
    return this->sum(j) - this->sum(i) + this->nums[i];
}

int NumArray::sum(int i) {
    i += 1;
    int ret = 0;
    while (i > 0) {
        ret += subArray[i - 1];
        i -= i & -i;
    }
    return ret;
}
