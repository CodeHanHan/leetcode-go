package offer14;

import org.junit.Test;

import static org.junit.Assert.*;

public class RopeTest {
    @Test
    public void main(){
        Rope r = new Rope();
        assert(r.cuttingRope(10)==36);

        Rope r1 = new Rope();
//        r1.cuttingRope(10);
        System.out.println(r1.cuttingRope(3));
        assert(r1.cuttingRope(3)==2);
    }
}