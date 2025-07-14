import java.util.stream.IntStream;

class Solution {
    public int[] productExceptSelf(int[] nums) {
        int product = 1;
        int zeroCount = 0;
        for(int i : nums) {
            if (i != 0){
                product = product * i;
            } else {
            zeroCount++;
            }
        }

        if (zeroCount > 1) {
            return new int[nums.length];
        }

        int [] res = new int[nums.length];
        boolean found = IntStream.of(nums).anyMatch(n -> n == 0);
        for(int j = 0; j < nums.length; j++) {
            if(nums[j] == 0) {
                res[j] = product;
                continue;
            } else if(found){
                res[j] = 0;
            } else {
                res[j] = product/nums[j];
            }
        }
        return res;
    }
}  
