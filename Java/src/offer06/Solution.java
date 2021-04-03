package offer06;

import java.util.ArrayList;

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Solution {
    public static class ListNode{
        int val;
        ListNode next;
        public ListNode(int x){
            val = x;
        }

    }
    ArrayList<Integer> holder = new ArrayList<>();

    public int[] reversePrint(ListNode head) {
        recur(head);
        int[] temp = new int[holder.size()];
        for(int i = 0; i < temp.length; i++){
            temp[i] = holder.get(i);
        }
        return temp;
    }

    private void recur(ListNode head){
        if(head == null) return;
        if(head.next!=null){
            recur(head.next);
        }
        holder.add(head.val);
    }

}
