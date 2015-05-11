import java.util.HashSet;
import java.util.Set;

/**
 * Created by xing on 15/4/24.
 * https://leetcode.com/problemset/algorithms/
 */
public class Solution {

    public int longestConsecutive(int[] nums) {
        Set<Integer> set = new HashSet<Integer>();
        for (int i : nums) {
            set.add(i);
        }
        int ret = 0;
        for (int i : nums) {
            boolean in = set.contains(i);
            if (!in) continue;
            int len = 1;
            int j = i+1;
            while (set.contains(j)) {
                len++;
                set.remove(j);
                j++;
            }
            j = i - 1;
            while (set.contains(j)) {
                len++;
                set.remove(j);
                j--;
            }
            if (len > ret) ret = len;
            set.remove(i);
        }
        return ret;
    }
}
