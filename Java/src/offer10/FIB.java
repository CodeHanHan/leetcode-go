package offer10;

public class FIB {
    //此方法超出时间限制
//    public int fib(int n) {
//        if(n == 0) return 0;
//        if(n == 1) return 1;
//        return fib(n-1) + fib(n-2);
//    }

    public int fib(int n){
        if(n == 0) return 0;
        if(n == 1) return 1;
        int[] fibnum = new int[101];
        fibnum[0] = 0;
        fibnum[1] = 1;
        for(int i  = 2; i <= n; i++){
            fibnum[i] = (fibnum[i-1] + fibnum[i-2])%1000000007;
        }
        return fibnum[n];
    }
}
