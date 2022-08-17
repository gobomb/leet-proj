package easy

import (
	"log"

	"github.com/google/go-cmp/cmp"

	"justest/pkg/ds/listnode"
	"justest/pkg/ds/tree"
	"justest/pkg/utils"
)

var (
	max    = utils.Max
	min    = utils.Min
	Null   = tree.Null
	null   = tree.Null
	isNull = tree.IsNull

	MakeListNode = listnode.MakeListNode
	LNDeepCopy   = listnode.LNDeepCopy
	funcName     = utils.FuncName
)

type (
	ListNode     = listnode.ListNode
	LNDeepCopyer = listnode.LNDeepCopyer
	TreeNode     = tree.TreeNode
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func ShouldEqualDiff(actual interface{}, expected ...interface{}) string {
	if len(expected) == 0 {
		return "expected should not be empty"
	}

	return cmp.Diff(expected[0], actual)
}

func MakeTree(arg ...int) *TreeNode {
	return tree.MakeTree(arg)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
