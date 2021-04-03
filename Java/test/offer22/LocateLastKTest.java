package offer22;

import org.junit.Test;

import static org.junit.Assert.*;

public class LocateLastKTest {
    @Test
    public void main(){
        LocateLastK l = new LocateLastK(1);
        l.insertToTail(2);
        l.insertToTail(3);
        l.insertToTail(4);
        l.insertToTail(5);
        System.out.println(l.getKthFromEnd(l.head, 3).val);
        System.out.println(l.getKthFromEnd(l.head, 1).val);
        System.out.println(l.getKthFromEnd(l.head, 0));


    }

}