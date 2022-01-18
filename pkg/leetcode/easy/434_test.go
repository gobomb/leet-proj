package easy

import "testing"

func Test_countSegments(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "Hello, my name is John",
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				s: "Hello",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments(tt.args.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}
