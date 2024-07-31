package medium

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_minSetSize(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				arr: []int{7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "len(arr)==0",
			args: args{
				arr: nil,
			},
			want: 0,
		},
		{
			name: "len(arr)==1",
			args: args{
				arr: []int{2},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				arr: []int{9, 77, 63, 22, 92, 9, 14, 54, 8, 38, 18, 19, 38, 68, 58, 19},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		convey.Convey(tt.name, t, convey.StackFail, func() {
			got := minSetSize(tt.args.arr)
			convey.So(got, ShouldEqualDiff, tt.want)
		})
	}
}
