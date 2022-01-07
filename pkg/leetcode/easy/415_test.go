package easy

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				num1: "11",
				num2: "123",
			},
			want: "134",
		},
		{
			name: "",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
		{
			name: "",
			args: args{
				num1: "456",
				num2: "77",
			},
			want: "533",
		},
		{
			name: "",
			args: args{
				num1: "1",
				num2: "9",
			},
			want: "10",
		},
		{
			name: "",
			args: args{
				num1: "6",
				num2: "501",
			},
			want: "507",
		},
		{
			name: "",
			args: args{
				num1: "22222222222222222222",
				num2: "11111111111111111111",
			},
			want: "33333333333333333333",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
