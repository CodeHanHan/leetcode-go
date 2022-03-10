/**
* 八皇后问题
* 有一个 8x8 的棋盘，希望往里放 8 个棋子（皇后），每个棋子所在的行、列、对角线都不能有另一个棋子。
* 八皇后问题就是期望找到所有满足这种要求的放棋子方式。
**/

package backtrack

import (
	"bytes"
	"fmt"
	"strings"
)

type CheckerBoard struct {
	matrix [8]int //下标表示行，值表示queen放在哪一列
}

func (c *CheckerBoard) cal8queen(row int) {
	if row == 8 { // 8个棋子全部放好
		fmt.Println(c)
		return
	}

	for column := 0; column < 8; column++ {
		if c.isOK(row, column) {
			c.matrix[row] = column
			c.cal8queen(row + 1)
		}
	}
}

func (c *CheckerBoard) isOK(row, column int) bool {
	leftUp, rightUp := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		if c.matrix[i] == column { // 当前列上有棋子
			return false
		}

		if leftUp >= 0 && c.matrix[i] == leftUp { // 左上角有棋子
			return false
		}

		if rightUp < 8 && c.matrix[i] == rightUp { // 右上角有棋子
			return false
		}
		leftUp--
		rightUp++
	}

	return true
}

func (c *CheckerBoard) String() string {
	buf := bytes.Buffer{}
	rowStr := []string{"+", "+", "+", "+", "+", "+", "+", "+"}
	for i := 0; i < 8; i++ {
		rowStr[c.matrix[i]] = "Q"
		buf.WriteString(strings.Join(rowStr, " "))
		rowStr[c.matrix[i]] = "+"
		buf.WriteString("\n")
	}

	return buf.String()
}
