package medium

import (
	"justest/pkg/ds/listnode"
	"justest/pkg/ds/tree"
	"justest/pkg/utils"
	"log"
	"reflect"
	"runtime"
	"strings"
)

var (
	max          = utils.Max
	min          = utils.Min
	Null         = tree.Null
	null         = tree.Null
	MakeListNode = listnode.MakeListNode
	LNDeepCopy   = listnode.LNDeepCopy
	isNull       = tree.IsNull
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

func MakeTree(arg ...int) *TreeNode {
	return tree.MakeTree(arg)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func funcName(f interface{}) string {
	sp := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	return sp[len(sp)-1]
}
