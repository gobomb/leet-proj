package easy

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				nums: []int{313, 64, 612, 515, 412, 661, 244, 164, 492, 744,
					233, 579, 62, 786, 342, 817, 838, 396, 230, 79,
					570, 702, 124, 727, 586, 542, 919, 372, 187, 626,
					869, 923, 103, 932, 368, 891, 411, 125, 724, 287,
					575, 326, 88, 125, 746, 117, 363, 8, 881, 441,
					542, 653, 211, 180, 620, 175, 747, 276, 832, 772,
					165, 733, 574, 186, 778, 586, 601, 165, 422, 956,
					357, 134, 857, 114, 450, 64, 494, 679, 495, 205,
					859, 136, 477, 879, 940, 139, 903, 689, 954, 335,
					859, 78, 20},
				k: 22,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
