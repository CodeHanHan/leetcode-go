package offer10II;

public class Forg {
    public int numWays(int n) {
        if(n == 0) return 1;
        if(n == 1) return 1;
        if(n == 2) return 2;
        int[] m = new int[101];
        m[0] = 1;
        m[1] = 1;
        m[2] = 2;
        for(int i = 3; i <= n; i++){
            m[i] = (m[i-1] + m[i-2])/1000000007;
        }
        return m[n];
    }
}
