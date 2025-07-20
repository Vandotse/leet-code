import java.util.*;

class Solution {
    public int longestConsecutive(int[] nums) {
        HashMap<Integer, Integer> map1 = new HashMap<>();
        int longest = 0;
        ArrayList<Integer> starts = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) {
            map1.put(nums[i],map1.getOrDefault(nums[i], 0) + 1);
        }
    
        for (int j = 0; j < nums.length; j++) {
            if(!map1.containsKey((nums[j] - 1))) {
                starts.add(nums[j]);
            }
        }

        for (int s : starts) {
            boolean exists = true;
            int number = s;
            int sequence = 1;
            while(exists) {
                if(map1.containsKey((number + 1))) {
                    sequence++;
                    number++;
                } else {
                    exists = false;
                }
            }
            if (sequence > longest) {
                longest = sequence;
            }
        }

        return longest;

    }
}
