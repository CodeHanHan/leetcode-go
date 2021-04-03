package offer13;

import java.util.Map;
import java.util.HashMap;
import java.util.Objects;
public class RobotWalk {

//    private static final int[][] next = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}};
//    private int cnt = 0;
//    private int m;
//    private int n;
//    private int k;
//
//    public int movingCount(int m, int n, int k) {
//
//        this.m = m;
//        this.n = n;
//        this.k = k;
//        int cal;
//        boolean[][] flag = new boolean[m][n];
//        dfs(0,0,flag);
//        return cnt;
//    }
//    private void dfs(int i, int j,boolean[][] flag ) {
//
//        if (i < 0 || i >= m || j< 0 || j >= n || flag[i][j]) {
//            return;
//        }
//        if(cal(i) + cal(j) > k) {
//            return;
//        }
//        cnt++;
//        flag[i][j] = true;
//        for(int p = 0; p < 4; p++) {
//            int newi= i + next[p][0];
//            int newj= j + next[p][1];
//            dfs(newi , newj , flag);
//        }
//    }
//
//
//    private int cal(int num) {
//        int ref = 0;
//        while(num > 0) {
//            ref += num % 10;
//            num /= 10;
//        }
//        return ref;
//    }




    public static class Loc {
        private int m;
        private int n;

        public Loc(int m, int n) {
            this.m = m;
            this.n = n;
        }

        @Override
        public boolean equals(Object obj) {
            if (this == obj) return true;
            if (obj == null) return false;
            Loc loc = (Loc) obj;
            if (loc.m == this.m && loc.n == this.n) return true;
            return false;
        }

        @Override
        public int hashCode() {
            return Objects.hash(m, n);
        }
    }

    Map<Loc, Boolean> holder = new HashMap<>();
    public  int count = -1;

    public int movingCount(int m, int n, int k){
//        holder.put(new Loc(0,0), true);
        walk(m, n, 0, 0, k, holder);
        return count;
    }

    public void walk(int m , int n, int cur_m, int cur_n, int k, Map<Loc, Boolean> holder){
        //这个位置不可访问
        if(cur_m < 0 || cur_m >= m || cur_n < 0 || cur_n >= n) return ;
        if(cur_m/10 + cur_m%10 + cur_n/10 + cur_n%10 > k) return;
        if(holder.containsKey(new Loc(cur_m, cur_n))) return;


        holder.put(new Loc(cur_m, cur_n), true);//此位置可访问，访问当前位置
        if(count < holder.size()) count = holder.size();

        walk(m, n, cur_m-1, cur_n, k, holder);//上
        walk(m, n, cur_m+1, cur_n, k, holder);//下
        walk(m, n, cur_m, cur_n-1, k, holder); //左
        walk(m, n, cur_m, cur_n+1, k, holder);//右

//        holder.remove(new Loc(cur_m, cur_n));//后退时将访问过的位置清空
    }


}
