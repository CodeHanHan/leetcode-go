package offer11;

public class Solution {

    public int minArray(int[] numbers) {
        return search(numbers, 0, numbers.length-1);
    }

    private int search(int[] numbers, int left, int right){
        int mid = left + ((right-left) >> 1);
//        int mid = (left + right) / 2;
        if(left == right){
            return numbers[mid];
        }
        if(numbers[mid] < numbers[right]){
            return search(numbers, left, mid);
        }else if(numbers[mid] > numbers[right]){
            return search(numbers, mid+1, right);
        }else{
            return search(numbers, left, right-1);
        }
    }
}
