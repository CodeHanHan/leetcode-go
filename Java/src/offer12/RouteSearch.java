package offer12;

public class RouteSearch {
    boolean found = false;
    public boolean exist(char[][] board, String word) {
        if (board.length == 0) return false;
        boolean[][] visited = new boolean[board.length][board[0].length];
        if(board.length*board[0].length < word.length()) return false;  //剪枝，总长度小于模式串长度
        for(int row = 0; row < board.length; row++){
            for(int column = 0; column < board[0].length; column++){
                if(found) return found;  //在前面的搜索中已经找到一个，直接返回
                found = false;
                visited = new boolean[board.length][board[0].length];  // 更新visited矩阵
                if (word.charAt(0) == board[row][column]) { //当前字符成功匹配
                    if(word.length() == 1){  //只有一个字符
                        found = true;   //找到了
                        return found;
                    }
                    //当前字符与第一个字符相同，从此处开始搜索
                    //内部需要回溯，否则可能漏掉正确解
                    visited[row][column] = true;
                    int j = 1;
                    //朝上下左右四个方向搜索
                    if ((row - 1) >= 0 && board[row - 1][column] == word.charAt(j) && !visited[row - 1][column]) {  //上
                        route(board, word, visited, row - 1, column, j + 1);
                    }
                    if ((row + 1) < board.length && board[row + 1][column] == word.charAt(j) && !visited[row + 1][column]) { //下
                        route(board, word, visited, row + 1, column, j + 1);
                    }
                    if ((column - 1) >= 0 && board[row][column - 1] == word.charAt(j) && !visited[row][column - 1]) {  //左
                        route(board, word, visited, row, column - 1, j + 1);
                    }
                    if ((column + 1) < board[0].length && board[row][column + 1] == word.charAt(j) && !visited[row][column + 1]) {  //右
                        route(board, word, visited, row, column + 1, j + 1);
                    }
                }
            }
        }
        return found;
    }

    public void route(char[][] board, String word, boolean[][] visited, int r, int c, int j) {
        /*
         * @Des: 回溯计算路径
         * @Para:
         *      board: 给定的矩阵
         *      word: 要求搜索的字符串
         *      visisted：访问矩阵
         *      r: 当前行
         *      c: 当前列
         *      j：当前字符位置
         * @Return:
         *      是否可行
         **/
        visited[r][c] = true;
        if (j == word.length()) {
            found = true;
            return;
        }
        if ((r - 1) >= 0 && board[r - 1][c] == word.charAt(j) && !visited[r - 1][c]) {  //上
             route(board, word, visited, r - 1, c, j + 1);
        }
        if ((r + 1) < board.length && board[r + 1][c] == word.charAt(j) && !visited[r + 1][c]) { //下
             route(board, word, visited, r + 1, c, j + 1);
        }
        if ((c - 1) >= 0 && board[r][c - 1] == word.charAt(j) && !visited[r][c - 1]) {  //左
             route(board, word, visited, r, c - 1, j + 1);
        }
        if ((c + 1) < board[0].length && board[r][c + 1] == word.charAt(j) && !visited[r][c + 1]) {  //右
             route(board, word, visited, r, c + 1, j + 1);
        }
        visited[r][c] = false;   //四个方向都不能走，此路不通，清除标记
    }
}

