import java.util.*;

class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> values = new HashMap<>();
        List<Integer>[] freq = new List[nums.length + 1];

        for (int j = 0; j < freq.length; j++) {
            freq[j] = new ArrayList<>();
        }

        for (int i = 0; i < nums.length; i++) {
            values.put(nums[i],values.getOrDefault(nums[i], 0) + 1);
        }

        for (Map.Entry<Integer, Integer> entry : values.entrySet()) {
            freq[entry.getValue()].add(entry.getKey());
        }

        int [] res = new int[k];
        int index = 0;
        for (int i  = freq.length - 1; i > 0 && index < k; i--) {
            for(int n : freq[i]) {
                res[index++] = n;
                if (index == k) {
                    return res;
                }
            }
        }
        return res;


    }
}
