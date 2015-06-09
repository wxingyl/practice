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

    public int reverse(int x) {
        String maxIntStr = Integer.toString(Integer.MAX_VALUE);
        String minIntStr = Integer.toString(Integer.MIN_VALUE);
        String s = new StringBuilder(Integer.toString(x)).reverse().toString();
        if (x < 0) {
            s = '-' + s.substring(0, s.length()-1);
        }
        
        if ((x > 0 && s.length() == maxIntStr.length() && s.compareTo(maxIntStr) > 0)
                || (x < 0 && s.length() == minIntStr.length() && s.compareTo(minIntStr) > 0)) return 0;
        else return Integer.valueOf(s);
    }
}
