package leetcode

import (
	"container/list"
	"strings"
)

func simplifyPath(path string) string {
	ls := list.New()
	sp := strings.Split(path, "/")
	// log.Printf("%#v\n", sp)
	for i := 0; i < len(sp); i++ {
		switch sp[i] {
		case "..":
			b := ls.Back()
			if b != nil {
				ls.Remove(b)
			}
		case ".":
		case "":
		default:
			ls.PushBack(sp[i])
		}
	}
	rs := ""
	for {
		f := ls.Front()
		if f == nil {
			break
		}
		ls.Remove(f)

		s := f.Value.(string)
		if s == "" {
			continue
		}
		rs = rs + "/" + s
	}
	if rs == "" {
		rs = "/"
	}
	return rs
}

func simplifyPath2(path string) string {
	stack := []string{}
	sp := strings.Split(path, "/")
	// log.Printf("%#v\n", sp)
	for i := 0; i < len(sp); i++ {
		switch sp[i] {
		case "..":
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		case ".":
		case "":
		default:
			stack = append(stack, sp[i])
		}
	}
	return "/" + strings.Join(stack, "/")
}
