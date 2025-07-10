import java.util.HashMap;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> map1 = new HashMap<>();
        int[] result = new int[2];
        for(int i = 0; i < nums.length; i++) {
            int difference = target - nums[i];
            if(map1.containsKey(difference)) {
                result[0] = map1.get(difference);
                result[1] = i;
                return result; 
            }
            map1.put(nums[i], i);
        }
        return result;
    }
}
