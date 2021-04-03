package offer04;
import org.junit.Test;
import org.junit.runner.RunWith;

import static org.junit.Assert.*;

public class SolutionTest {
    @Test
    public  void main() {
                int[][] matrix1 = {  {1,   4,  7, 11, 15},
                            {2,   5,  8, 12, 19},
                            {3,   6,  9, 16, 22},
                            {10, 13, 14, 17, 24},
                            {18, 21, 23, 26, 30}
                            };
        int[][] matrix2 = {};
        Solution s = new Solution();
        s.findNumberIn2DArray(matrix1, 5);
        s.findNumberIn2DArray(matrix2, 5);
    }

}
