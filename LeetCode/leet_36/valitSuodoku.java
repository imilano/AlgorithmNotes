// 对每个数字独立编码，保证这个数字在一行或者一列或者一个块内重复出现时能够快速被检测出来。
// 对于每个1到9内的数字来说，其在每行每列和每个小区间内都是唯一的，将数字放在一个括号中，每行上的数字就将行号放在括号左边，
// 每列上的数字就将列数放在括号右边，每个小区间内的数字就将在小区间内的行列数分别放在括号的左右两边，这样每个数字的状态都是独一无二的存在，
// 就可以在 HashSet 中愉快地查找是否有重复存在啦

class Solution {
    public boolean isValidSudoku(char[][] board) {
            Set seen = new HashSet();
            for (int i=0; i<9; ++i) {
                for (int j=0; j<9; ++j) {
                    if (board[i][j] != '.') {
                        String b = "(" + board[i][j] + ")";
                        if (!seen.add(b + i) || !seen.add(j + b) || !seen.add(i/3 + b + j/3))
                            return false;
                    }
                }
            }
            return true;
    }
}