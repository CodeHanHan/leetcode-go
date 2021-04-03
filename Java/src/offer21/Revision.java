package offer21;

public class Revision {
    public int[] exchange(int[] nums) {
        if(nums.length <= 1) return nums;
        int i = 0,j = nums.length-1;
        int temp;

        //初始化
        while(i < nums.length && nums[i] % 2 != 0 ){
            i++;
        }
        while(j >= 0 && nums[j] % 2 != 1){
            j--;
        }
        if(i >= nums.length || j <= 0) return nums;  //如果只有偶数或者只有奇数

        while(i<j){
            temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;

            while(i < nums.length && nums[i] % 2 != 0  ){
                i++;
            }
            while(j >= 0 && nums[j] % 2 != 1 ){
                j--;
            }
        }
        return nums;
    }


}
