package offer16;

import org.junit.Test;

import static org.junit.Assert.*;

public class PowTest {
    @Test
    public void main(){
        Pow p = new Pow();
        System.out.println(p.myPow(0.00001, 2147483647 ));
    }
}