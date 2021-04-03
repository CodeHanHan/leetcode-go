package offer05;

public class Solution {
    public String replaceSpace(String s) {
        StringBuilder builder = new StringBuilder();
        for(int i = 0; i < s.length(); i++){
            if(s.charAt(i) == ' '){
                builder.append("%20");
                continue;
            }
            builder.append(s.charAt(i));
        }
        return builder.toString();
    }

}