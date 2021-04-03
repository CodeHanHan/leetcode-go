package offer11;

import org.junit.Test;

import static org.junit.Assert.*;

public class SolutionTest {
    @Test
    public void main(){
        int[] a = {3,4,5,1,2};
        Solution solution = new Solution();
        assert(solution.minArray(a) == 1);

        int[] b = {2,2,2,0,1};
        assert(solution.minArray(b) == 0);
    }


}