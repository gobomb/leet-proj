package easy

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n            int
		isBadVersion func(int) bool
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n:            5,
				isBadVersion: func(ver int) bool { return ver >= 4 },
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				n:            1,
				isBadVersion: func(ver int) bool { return ver >= 1 },
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isBadVersion = tt.args.isBadVersion
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
