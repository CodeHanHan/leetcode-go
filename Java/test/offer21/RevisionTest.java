package offer21;

import org.junit.Test;

import static org.junit.Assert.*;

public class RevisionTest {
    @Test
    public void main(){
        int[] arr1 = {1,2,3,4};
        int[] arr2 = {1,7,3,5,4,2,7,7,5,5,7};
        int[] arr3 = {1,2,3,4,5,6,7,8,9};
        int[] arr4 = {1,3,5};
        Revision r = new Revision();
        r.exchange(arr1);
        r.exchange(arr2);
        r.exchange(arr3);
        r.exchange(arr4);
        System.out.println(r.exchange(arr1).toString());
        System.out.println(r.exchange(arr2).toString());
    }
}