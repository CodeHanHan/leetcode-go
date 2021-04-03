package offer17;

public class Solution17 {
    public int[] printNumbers(int n) {
        if( n == 0) return new int[0];
        int num = (int)Math.pow(10, n);
        int[] result = new int[num-1];
        for(int i = 1; i < num; i++){
            result[i-1] = i;
        }
        return result;
    }
}
