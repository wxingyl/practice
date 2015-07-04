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

    /**
     * https://leetcode.com/problems/letter-combinations-of-a-phone-number/
     */
    public List<String> letterCombinations(String digits) {
        List<String> ret = new LinkedList<String>();
        if (digits == null || digits.isEmpty()) return ret;
        Map<Character, String> map = new HashMap<Character, String>();
        map.put('2', "abc");
        map.put('3', "def");
        map.put('4', "ghi");
        map.put('5', "jkl");
        map.put('6', "mno");
        map.put('7', "pqrs");
        map.put('8', "tuv");
        map.put('9', "wxyz");
        fun(null, digits, map, ret);
        return ret;
    }

    public void fun(String s, String digits, Map<Character, String> map, List<String> ret) {
        String subStr = digits.substring(1);
        boolean add = subStr.isEmpty();
        String charStr = map.get(digits.charAt(0));
        if (charStr == null || charStr.isEmpty()) {
            fun(s, subStr, map, ret);
        }  else {
            if (s == null) s = "";
            for (char ch : charStr.toCharArray()) {
                if (add) ret.add(s + ch);
                else fun(s + ch, subStr, map, ret);
            }
        }
    }
}
