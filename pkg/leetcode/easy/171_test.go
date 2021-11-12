package easy

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				columnTitle: "A",
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				columnTitle: "AB",
			},
			want: 28,
		},
		{
			name: "3",
			args: args{
				columnTitle: "ZY",
			},
			want: 701,
		},
		{
			name: "4",
			args: args{
				columnTitle: "FXSHRXW",
			},
			want: 2147483647,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
