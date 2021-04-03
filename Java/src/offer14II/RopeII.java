package offer14II;

import java.math.BigInteger;

public class RopeII {
    public int cuttingRope(int n) {
        if(n <= 3) return n-1;
        BigInteger[] dp = new BigInteger[n+1];
        dp[0] = new BigInteger("1");
        dp[1] = new BigInteger("1");
        dp[2] = new BigInteger("1");
        for(int i = 3; i < n+1; i++){  //遍历整个dp数组
            //状态转移方程：dp[i]=max(dp[i], dp[i-j]*j，j*(i-j) )
            dp[i] = new BigInteger("1");
            for(int j = 1; j < i; j++){
//                dp[i] = Math.max(dp[i], Math.max(dp[i-j]*j, j*(i-j)));
                dp[i] = dp[i].max(
                        dp[i-j].multiply(new BigInteger(String.valueOf(j))).max(
                                new BigInteger(String.valueOf(j)).multiply(
                                        new BigInteger(String.valueOf(i)).subtract(
                                                new BigInteger(String.valueOf(j))
                                        ))
                                )
                        );
                //这里的dp[i]暂存了在计算过程中的最大值，因此需要保留，否则可能会丢失最大值
            }
        }
        return dp[n].mod(new BigInteger("1000000007")).intValue();
    }
}
