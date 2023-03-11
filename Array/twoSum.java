// nums = [2,7,11,15], target = 9]
// output = [0,1]
// nums[0] + nums[1] == 9, we return [0, 1].


class Solution {
    public int[] twoSum(int[] nums, int target) {
        for(int i = 0; i < nums.length; i++){
            for(int j = i + 1; j < nums.length; j++){
                if(nums[i] + nums[j] == target){
                 return new int[] {i,j};
                }
            }
            
        }
        return new int[] {-1,-1};
    }
}

// time complexity : O(n^2)
// space complexity : O(1)
