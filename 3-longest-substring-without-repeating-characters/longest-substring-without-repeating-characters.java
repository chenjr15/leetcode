import java.util.HashMap;
// import java.lang.Math.*;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        // HashMap
        int longest = 0;
        HashMap<Character, Integer> got = new HashMap<Character, Integer>();
        for (int i = 0; i < s.length(); i++) {
            if (got.containsKey(s.charAt(i))) {
                longest = Math.max(longest, got.size());
                i = ogt.get(s.charAt(i));
                got.clear();
            } else {
                got.put(s.charAt(i), i);
            }
        }

        return Math.max(longest, got.size());
    }
}t-repeating-characterslongest-substring-without-repeating-characters
