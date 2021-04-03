package offer16;

public class Pow {
    // 解法一  当输入为：1.00000， 2147483647 超出时间限制
//    public double myPow(double x, int n) {
//        // n可以为负数
//        if(n == 0) return 1;
//        double result = 1;
//        if(n < 0){
//            for(int i = 0; i < -n; i++){
//                result *= x;
//            }
//            return 1/result;
//        }
//        for(int i = 0; i < n; i++){
//            result *= x;
//        }
//        return result;
//    }

    //解法二 快速幂
    // 当输入为0.00001, 2147483647时超出时间限制
    public double myPow(double x, int n){
        if(n == 0) return 1;
        if(x == 1) return 1;
        if(n == 1) return x;
        if(n < 0) {
            x = 1 / x;
            n = -n;
        }
        if(n % 2 == 0 ) return myPow(x, n/2) * myPow(x, n/2);
        return x * myPow(x, n/2) * myPow(x, n/2);
    }
}
