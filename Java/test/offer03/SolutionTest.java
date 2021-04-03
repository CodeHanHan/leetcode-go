package offer03;

import org.junit.Test;
import org.junit.runner.RunWith;
import static org.junit.Assert.*;

public class SolutionTest {
    @Test
    public void main() {
        int[] data = {2,3, 1, 0, 2, 5, 3};
        Solution s = new Solution(data);
        System.out.println(s.findRepeatNumber(data));
    }
}
