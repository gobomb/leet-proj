package binsearch

import "testing"

func Test_binAll(t *testing.T) {
	type args struct {
		arr []int
		k   int
		t   T
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "找重复数字的左侧边界",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   3,
				t:   T1,
			},
			want: 1,
		},
		{
			name: "找重复数字的左侧边界:target<all",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   0,
				t:   T1,
			},
			want: -1,
		},
		{
			name: "找重复数字的左侧边界:targe>all",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   8,
				t:   T1,
			},
			want: -1,
		},
		{
			name: "找重复数字的左侧边界:无数字",
			args: args{
				arr: []int{},
				k:   1,
				t:   T1,
			},
			want: -1,
		},

		{
			name: "找重复数字的右侧边界",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   3,
				t:   T2,
			},
			want: 2,
		},
		{
			name: "找重复数字的右侧边界:target<all",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   0,
				t:   T2,
			},
			want: -1,
		},
		{
			name: "找重复数字的右侧边界:target>all",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   8,
				t:   T2,
			},
			want: -1,
		},
		{
			name: "找重复数字的右侧边界:无数字",
			args: args{
				arr: []int{},
				k:   1,
				t:   T2,
			},
			want: -1,
		},


		{
			name: "找小于target的最大下标",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   3,
				t:   T3,
			},
			want: 0,
		},
		{
			name: "找小于target的最大下标:target<all",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   0,
				t:   T3,
			},
			want: -1,
		},
		{
			name: "找小于target的最大下标:target>all",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   8,
				t:   T3,
			},
			want: 3,
		},
		{
			name: "找小于target的最大下标:无数字",
			args: args{
				arr: []int{},
				k:   1,
				t:   T3,
			},
			want: -1,
		},

		{
			name: "找大于target的最小下标",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   3,
				t:   T4,
			},
			want: 3,
		},
		{
			name: "找大于target的最小下标:target<all",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   0,
				t:   T4,
			},
			want: 0,
		},
		{
			name: "找大于target的最小下标:target>all",
			args: args{
				arr: []int{1, 3, 3, 7},
				k:   8,
				t:   T4,
			},
			want: -1,
		},
		{
			name: "找大于target的最小下标:无数字",
			args: args{
				arr: []int{},
				k:   1,
				t:   T4,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binAll(tt.args.arr, tt.args.k, tt.args.t); got != tt.want {
				t.Errorf("binAll(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
