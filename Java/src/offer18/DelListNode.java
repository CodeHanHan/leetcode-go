package offer18;

public class DelListNode {
    public static class ListNode{
        public int val;
        public ListNode next;

        public ListNode(int val){
            this.val = val;
            this.next = null;
        }
    }

    public ListNode build(int[] data){
        int l = data.length;
        if(l == 0) return null;
        ListNode[] nodes = new ListNode[l];
        for(int i = 0; i < l; i++){
            nodes[i] = new ListNode(data[i]);
            if(i+1 < l) nodes[i].next = nodes[i+1];
        }
        return nodes[0];
    }

    public ListNode deleteNode(ListNode head, int val) {
        if(head == null ) return head;
        ListNode cur = head;
        ListNode pre = null;
        while(cur.val != val && cur != null){
            pre = cur;
            cur = cur.next;
        }
        if(cur == null) return head;  //没找到节点
        if(pre == null){  //第一个节点就是目标节点
            return null;
        }
        pre.next = cur.next;
        return head;
    }
}
