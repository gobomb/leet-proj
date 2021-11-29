package medium

import (
	"testing"
)

func TestConstructorWordDictionary(t *testing.T) {
	obj := ConstructorWordDictionary()
	null := false

	tests := []struct {
		opt    []string
		args   []string
		output []bool
	}{
		{
			opt:    []string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
			args:   []string{"", "bad", "dad", "mad", "pad", "bad", ".ad", "b.."},
			output: []bool{null, null, null, null, false, true, true, true},
		},
		{
			opt:    []string{"WordDictionary", "addWord", "addWord", "addWord", "addWord", "search", "search", "addWord", "search", "search", "search", "search", "search", "search"},
			args:   []string{"", "at", "and", "an", "add", "a", ".at", "bat", ".at", "an.", "a.d.", "b.", "a.d", "."},
			output: []bool{null, null, null, null, null, false, false, null, true, true, false, false, true, false},
		},
	}

	rs := false

	for _, tt := range tests {
		opt := tt.opt
		args := tt.args
		output := tt.output

		for i, o := range opt {
			rs = null

			switch o {
			case "WordDictionary":
				obj = ConstructorWordDictionary()
			case "addWord":
				obj.AddWord(args[i])
			case "search":
				rs = obj.Search(args[i])
			}

			if rs != output[i] {
				t.Fatalf("failed at %v opt: %v, args: %v, output: %v\n %v", i, opt[i], args[i], output[i], obj.t.children)
			} else {
				// log.Printf("%v opt: %v, args: %v, output: %v\n", i, opt[i], args[i], output[i])
			}
		}
	}
}
