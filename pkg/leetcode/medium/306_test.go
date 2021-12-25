package medium

import "testing"

func Test_isAdditiveNumber(t *testing.T) {
	type args struct {
		num string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				num: "112358",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				num: "199100199",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				num: "11235812359",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				num: "199999100000199999",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				num: "199999",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				num: "1023",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				num: "1203",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				num: "101",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				num: "000",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				num: "011",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				num: "0011",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdditiveNumber(tt.args.num); got != tt.want {
				t.Errorf("isAdditiveNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
