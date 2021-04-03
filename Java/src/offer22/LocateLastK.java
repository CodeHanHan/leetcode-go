package offer22;

import offer13.RobotWalk;

import java.util.List;

public class LocateLastK {

    public ListNode head;

    public static class ListNode {
        int val;
        ListNode next;

        public ListNode(int x) {
            val = x;
        }
    }

    public LocateLastK(int val){
        this.head = new ListNode(val);
    }

    public void insertToTail(int val) {
        ListNode cur = head;
        while(cur.next != null){
            cur = cur.next;
        }
        cur.next = new ListNode(val);
    }

    public ListNode getKthFromEnd(ListNode head, int k) {
        ListNode cur, pre;
        cur = head;
        for (int i = 0; i < k; i++) {
            pre = cur;
            if(cur == null){
                return null;
            }
            cur = cur.next;
        }
        pre = head;
        while (cur != null) {
            pre = pre.next;
            cur = cur.next;
        }
        return pre;
    }
}
