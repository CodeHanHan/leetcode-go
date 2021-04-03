package offer19;

import org.junit.Test;

import static org.junit.Assert.*;

public class PatternTest {
    @Test
    public void main(){
        String s = "aa";
        String p = "a*";
        Pattern pattern = new Pattern();
        assert( pattern.isMatch(s, p));

        String s1 = "ab";
        String p1 = ".*";
        Pattern pattern1 = new Pattern();
        assert( pattern1.isMatch(s1, p1));

        String s2 = "aab";
        String p2 = "c*a*b";
        Pattern pattern2 = new Pattern();
        assert( pattern2.isMatch(s2, p2));

        String s3 = "ippi";
        String p3 = "p*.";
        Pattern pattern3 = new Pattern();
        assert(!pattern2.isMatch(s3, p3));

    }

}