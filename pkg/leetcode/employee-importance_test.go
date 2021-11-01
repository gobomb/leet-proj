package leetcode

import "testing"

func Test_getImportance(t *testing.T) {
	type args struct {
		employees []*Employee
		id        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//[[1,5,[2,3]],[2,3,[]],[3,3,[]]]
		{
			name: "1",
			args: args{
				employees: []*Employee{
					{
						ID:           1,
						Importance:   5,
						Subordinates: []int{2, 3},
					},
					{
						ID:           2,
						Importance:   3,
						Subordinates: []int{},
					},
					{
						ID:           3,
						Importance:   3,
						Subordinates: []int{},
					},
				},
				id: 1,
			},
			want: 11,
		},
		{
			name: "2",
			args: args{
				employees: []*Employee{
					{
						ID:           1,
						Importance:   2,
						Subordinates: []int{5},
					},
					{
						ID:           5,
						Importance:   -3,
						Subordinates: []int{},
					},
				},
				id: 5,
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getImportance(tt.args.employees, tt.args.id); got != tt.want {
				t.Errorf("getImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}
