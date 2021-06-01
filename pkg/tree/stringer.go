package tree

import (
	"fmt"
	"io"
	"strconv"
)

type Printer struct {
	root *TreeNode
}

func TreeStringer(t *TreeNode) fmt.Stringer {
	return &Printer{
		root: t,
	}
}

func (p *Printer) String() string {
	if p == nil {
		return "Null"
	}
	return printAvlTree(p.root)
}

/**
 * ref: https://zhuanlan.zhihu.com/p/358578675
 * 
 * row为当前打印行，column为当前打印节点列，treeHeight为真正的树高
 * resArray为输出，填充字符数组
 */
func writeArray(root *TreeNode, row, column, treeHeight int, resArray [][]string) {
	if root == nil {
		return
	}
	resArray[row][column] = strconv.Itoa(root.Val)
	currentHeight := (row + 1) / 2 // 当前高度(传入的row包含连接符行，所以当前行真正的高为(row + 1)/2)
	if currentHeight == treeHeight {
		return // 下面没有子节点了
	}
	gap := treeHeight - currentHeight - 1 // 父亲到左/右儿子的距离
	// 填充左儿子
	if root.Left != nil {
		// 先写树结构符号
		resArray[row+1][column-gap] = "/"
		// 再写左儿子（递归）
		writeArray(root.Left, row+2, column-gap*2, treeHeight, resArray)
	}
	// 填充右儿子
	if root.Right != nil {
		resArray[row+1][column+gap] = "\\"
		writeArray(root.Right, row+2, column+gap*2, treeHeight, resArray)
	}
}

// ref: https://zhuanlan.zhihu.com/p/358578675
// 预先算出高度，通过高度算出宽度，再把全部的空间分配出来，然后递归画图
func printAvlTree(root *TreeNode) string {
	height := getTreeHeight(root)
	// fmt.Printf("height: %v\n", height)
	// 总宽度为节点高度 * 2 - 1, 因为还要画树枝符号
	totalHeight := height*2 - 1
	// 最大宽度为3 * 2^(n - 1) + 1，公式如下：
	/**
	   父亲节点占1, 两个孩子空间各占1, 连接线各占1, 每个父子关系共占5, 每个关系之间空1, 最左最右各空1
	  第2行： 5 + 2 （1个父子结构占位+左右两个空格分割）
	  第3行：2 * 5 + (1 + 2) （2个父子结构占位+1个相邻父子结构间空格+左右两个空格分割）
	  第4行：4 * 5 + (3 + 2) （4个父子结构占位+3个相邻父子结构间空格+左右两个空格分割）
	  第5行：8 * 5 + (7 + 2)
	  第n行：5 * 2 ^ (n - 2) + (2 ^ (n - 2) - 1) + 2 = 6 * 2 ^ (n-2) + 1 = 3 * 2 ^ (n - 1) + 1
	*/
	var totalWidth int
	if height == 0 {
		totalWidth = 1
	} else {
		totalWidth = (2<<(height-2))*3 + 1
	}

	// 创建数组
	printArray := make([][]string, totalHeight)
	for i := range printArray {
		printArray[i] = make([]string, totalWidth)
		for j := range printArray[i] {
			printArray[i][j] = " "
		}
	}

	// 计算打印数组
	writeArray(root, 0, totalWidth/2, height, printArray)

	var s string
	// 打印
	for i := range printArray {
		var res string
		for j := range printArray[i] {
			res = res + printArray[i][j]
		}
		s = fmt.Sprintf("%s\n%s", s, res)
	}
	return s
}

// Ref: github.com/pkg/errors
func (l *TreeNode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", TreeStringer(l))
			return
		}
		fallthrough
	case 'h':
		fmt.Fprintf(s, "%v", getTreeHeight(l))
		return
	case 'p':
		io.WriteString(s, fmt.Sprintf("%p", l))
	}
}
