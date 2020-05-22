// 题目描述
// 请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一个格子开始，每一步可以在矩阵中向左，向右，向上，向下移动一个格子。如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。 例如 [abcesfcsadee]\begin{bmatrix} a & b & c &e \\ s & f & c & s \\ a & d & e& e\\ \end{bmatrix}\quad⎣⎡​asa​bfd​cce​ese​⎦⎤​  矩阵中包含一条字符串"bcced"的路径，但是矩阵中不包含"abcb"路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入该格子。

// 题解：回溯法的典型应用
#include <cstring>
#include <iostream>

class Solution {
public:
    bool hasPath(char* matrix, int rows, int cols, char* str)
    {
        if (matrix == nullptr || rows < 1 || cols < 1 || str == nullptr) return false;

        bool* visited = new bool[rows*cols];
        memset(visited,0,rows*cols);

        int pathLen = 0;
        for (int row = 0; row < rows; row++) {
            for (int col = 0; col < cols; col++) {
                if (hasPathCore(matrix,rows,cols,row,col,str,pathLen,visited)) {
                    return true;
                }
            }
        }
        delete[] visited;
        return false;
    }

    bool hasPathCore(const char* matrix, int rows, int cols, int row, int col, const char* str, int& pathLen, bool* visited) {
        if (str[pathLen] == '\0') return true; // r如果路径已经匹配到结尾，则返回

        bool hasPath = false; // 初始化为false是为了预防边界
        // 如果矩阵 row*cols + col 的值和str[pathLen]是一样的，那么我们继续便利该节点上下左右的路径
        if (row >=0 && row < rows && col >=0 && col < cols  && matrix[row*cols+col] == str[pathLen] && !visited[row*cols + col]) {
            ++pathLen;
            visited[row*cols+col] = true;
            hasPath = hasPathCore(matrix,rows,cols,row,col-1,str,pathLen,visited)
                    || hasPathCore(matrix,rows,cols,row-1,col,str,pathLen,visited)
                    || hasPathCore(matrix,rows,cols,row+1,col,str,pathLen,visited)
                    || hasPathCore(matrix,rows,cols,row,col+1,str,pathLen,visited);
            
            if (!hasPath) { // 如果没有路径，那么清理这一步遍历的痕迹，同时打上访问标签
                --pathLen;
                visited[row*cols+col] = false;
            }
        }
        return hasPath;
    }


};

/* 当矩阵坐标为(row,col)的格子和路径字符串中下表为pathLen的字符一样时，从4个相邻的格子(row,col-1)、(row-1,col)、(row+1,col)和(row, col+1)中去定位路径字符串中下标为pathLen+1的字符。
 * 如果4个相邻的格子都没有匹配字符串中下标为pathLen+1的字符，则表明当前路径字符串中下标为pathLen的字符在矩阵中的定位不正确，我们需要回到钱一个字符pathLen-1，然后重新定位。
 * 一直重复这个过程，直到路径字符串上的所有字符都在矩阵中找到合适的位置（此时，str[pathLen] == ‘\0’)
 */
