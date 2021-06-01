package leetcode

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// func Test_minWindow(t *testing.T) {
// 	type args struct {
// 		s string
// 		t string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{
// 			name: "1",
// 			args: args{
// 				s: "ADOBECODEBANC",
// 				t: "ABC",
// 			},
// 			want: "BANC",
// 		},
// 		{
// 			name: "2",
// 			args: args{
// 				s: "a",
// 				t: "a",
// 			},
// 			want: "a",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
// 				t.Errorf("minWindow(%v, %v) = %v, want %v", tt.args.s, tt.args.t, got, tt.want)
// 			}
// 		})
// 	}
// }
