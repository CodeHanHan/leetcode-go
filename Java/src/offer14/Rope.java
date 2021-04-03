package offer14;

public class Rope {

    //回溯法
//    private int n;  //绳子长度为n
//
//    private static int max;
//    private int[] k;
//
//    public int cuttingRope(int n) {
//        for(int m = 2; m <= n; m++){
//            k = new int[m];
//            cut(m, n, 0, k);
//        }
//        return max;
//    }
//
//    public void cut(int m, int n, int ki, int[] k){
//        if(ki == m-1){
//            k[ki] = n;
//            int c = cal(m);
//            if(max < c) max = c;
//            return;
//        }
//        for(int i = 1; i <= n-(m-ki)+1; i++){  //至少要分m段，因此要留一定的长度
//            k[ki] = i;
//            cut(m, n-i, ki+1, k);
//        }
//    }
//
//    public int cal(int m){
//        int a = 1;
//        for(int i = 0; i < m; i++){
//            a = a * k[i];
//        }
//        return a;
//    }

    //动态规划法
    public int cuttingRope(int n) {
        int[] dp = new int[n+1];
        dp[0] = 1;
        dp[1] = 1;
        dp[2] = 1;
        for(int i = 3; i < n+1; i++){  //遍历整个dp数组
            //状态转移方程：dp[i]=max(dp[i], dp[i-j]*j，j*(i-j) )
            for(int j = 1; j < i; j++){
                dp[i] = Math.max(dp[i], Math.max(dp[i-j]*j, j*(i-j)));
                //这里的dp[i]暂存了在计算过程中的最大值，因此需要保留，否则可能会丢失最大值
            }
        }
        return dp[n];
    }


}
