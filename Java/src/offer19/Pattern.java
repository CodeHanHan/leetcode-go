package offer19;

public class Pattern{
    public boolean isMatch(String s, String p) {
        if (p.length() <= 0)
            return s.length() <= 0;

        //第一个字符匹配成功
        boolean match = (s.length() > 0 && (s.charAt(0) == p.charAt(0) || p.charAt(0) == '.'));

        //下一个字符为*
        if (p.length() > 1 && p.charAt(1) == '*'){  // p.length()放在前面是为了当p长度为0时，短路后面的，

            // p.substring(2) 表示*匹配0个，模式串后移2位
            // s.substring(1) 表示*匹配多个
            return isMatch(s, p.substring(2)) || (match && isMatch(s.substring(1), p));
        } else {  //下一个字符不为*

            // 匹配成功，再匹配下一个
            return match && isMatch(s.substring(1), p.substring(1));
        }
    }
}