package offer18;

import org.junit.Test;

import static org.junit.Assert.*;

public class DelListNodeTest {
    @Test
    public void main(){
        DelListNode d = new DelListNode();
        int[] data = {};
        DelListNode.ListNode head = d.build(data);
        d.deleteNode(head, -3);
    }

}