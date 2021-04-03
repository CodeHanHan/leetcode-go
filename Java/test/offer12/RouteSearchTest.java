package offer12;

import org.junit.Test;

public class RouteSearchTest {
    @Test
    public void main() {
        RouteSearch r = new RouteSearch();
        char[][] board = {{'a', 'b'},
                {'c', 'd'}};
        String str = "abcd";
        assert (!r.exist(board, str));

        RouteSearch r1 = new RouteSearch();
        char[][] board1 = {
                {'a', 'b', 'c', 'e'},
                {'s', 'f', 'c', 's'},
                {'a', 'd', 'e', 'e'},
        };
        String str1 = "abcced";
        assert (r1.exist(board1, str1));

        RouteSearch r2 = new RouteSearch();
        char[][] board2 = {{'a', 'a'}};
        String str2 = "a";
        assert (r2.exist(board2, str2));

        RouteSearch r3 = new RouteSearch();
        char[][] board3 = {{'a'}};
        String str3 = "a";
        assert (r3.exist(board3, str3));

        RouteSearch r4 = new RouteSearch();
        char[][] board4 = {{'a'}};
        String str4 = "ab";
        assert (!r4.exist(board4, str4));

        RouteSearch r5 = new RouteSearch();
        char[][] board5 = {{'a', 'b'}};
        String str5 = "a";
        assert (r5.exist(board5, str5));

        RouteSearch r6 = new RouteSearch();
        char[][] board6 = {{}};
        String str6 = "a";
        assert (!r6.exist(board6, str6));

        RouteSearch r7 = new RouteSearch();
        char[][] board7 = {{'c', 'a', 'a'},
                           {'a', 'a', 'a'},
                           {'b', 'c', 'd'},
        };
        String str7 = "aab";
        assert (r7.exist(board7, str7));

        RouteSearch r8 = new RouteSearch();
        char[][] board8 = {{'a', 'b', 'c', 'e'},
                           {'s', 'f', 'c', 's'},
                           {'a', 'd', 'e', 'e'},
        };
        String str8 = "abcb";
        assert (!r8.exist(board8, str8));

        RouteSearch r9 = new RouteSearch();
        char[][] board9 = {{'a', 'b', 'c', 'e'},
                {'s', 'f', 'e', 's'},
                {'a', 'd', 'e', 'e'},
        };
        String str9 = "abceseeefs";
        assert (r9.exist(board9, str9));
    }

}