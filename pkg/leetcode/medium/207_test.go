package medium

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				numCourses:    7,
				prerequisites: [][]int{{2, 1}, {4, 1}, {3, 1}, {3, 4}, {4, 2}, {5, 2}, {4, 5}, {7, 5}, {6, 3}, {6, 7}},
			},
			want: true,
		},
	}

	for i := range tests[2].args.prerequisites {
		tests[2].args.prerequisites[i][0]--
		tests[2].args.prerequisites[i][1]--
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
