package offer06;

import org.junit.Test;

public class SolutionTest {
    @Test
    public void main(){
        Solution.ListNode head = new Solution.ListNode(1);
        head.next = new Solution.ListNode(2);
        head.next.next = new Solution.ListNode(3);
        Solution s = new Solution();

        int[] a = s.reversePrint(head);
        for (int value : a) {
            System.out.println(value);
        }

        Solution.ListNode head_null = null ;
        s.reversePrint(head_null);
    }
}
