package dynamicprogram

import "testing"

func Test_integerReplacement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 8,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				n: 7,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				n: 15,
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				n: 100,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.args.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
